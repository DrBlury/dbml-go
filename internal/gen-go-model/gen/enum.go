package gen

import (
	"dbml-go/core"
	"dbml-go/internal/gen-go-model/genutil"
	"fmt"
	"path/filepath"

	"github.com/dave/jennifer/jen"
)

func (g *generator) genEnum(enum core.Enum) error {
	f := jen.NewFilePathName(g.out, g.gopackage)

	enumOriginName := genutil.NormalizeTypeName(enum.Name)
	enumGoTypeName := genutil.NormalizeGoTypeName(enum.Name)

	f.Commentf("%s is generated type for enum '%s'", enumGoTypeName, enumOriginName)

	f.Type().Id(enumGoTypeName).Int64()

	f.Const().DefsFunc(func(group *jen.Group) {
		group.Id("_").Id(enumGoTypeName).Op("=").Iota()

		for _, value := range enum.Values {
			v := group.Id(genutil.NormalLizeGoName(value.Name))
			if value.Note != "" {
				v.Comment(value.Note)
			}
		}
	})

	f.Commentf("Scan implements Scan() from the sql.Scanner interface for %q", enumOriginName)
	f.Commentf("See: https://pkg.go.dev/database/sql#Scanner")

	f.Func().Params(
		jen.Id("v").Op("*").Id(enumGoTypeName),
	).Id("Scan").Params(
		jen.Id("i").Interface(),
	).Params(jen.Err().Error()).Block(
		jen.List(jen.Id("s"), jen.Id("ok")).Op(":=").Id("i").Assert(jen.Index().String()),
		jen.If(jen.Op("!").Id("ok").Block(
			jen.Return(jen.Qual("fmt", "Errorf").Call(jen.List(jen.Lit("value (%#v) is not of the expected type string"), jen.Id("i")))),
		)),
		jen.Switch().BlockFunc(func(group *jen.Group) {
			for _, value := range enum.Values {
				v := genutil.NormalLizeGoName(value.Name)

				group.Case(jen.Qual("reflect", "DeepEqual").Call(jen.List(jen.Id("s"), jen.Index().String().Params(jen.Lit(value.Name))))).Block(
					jen.Op("*").Id("v").Op("=").Id(v),
				)
			}

			group.Default().Block(
				jen.Err().Op("=").Qual("fmt", "Errorf").Call(jen.List(jen.Lit(fmt.Sprintf("unable to parse %%#v into a known %s type", enumGoTypeName)), jen.Id("i"))),
			)
		}),
		jen.Return(),
	)

	f.Commentf("Value implements Value() from the driver.Valuer interface for %q", enumOriginName)
	f.Commentf("See: https://pkg.go.dev/database/sql/driver#Valuer")

	if g.isPostgres {
		f.Func().Params(
			jen.Id("v").Id(enumGoTypeName),
		).Id("Value").Params().Params(jen.Id("out").Qual("database/sql/driver", "Value"), jen.Id("err").Error()).Block(
			jen.Switch(jen.Id("v")).BlockFunc(func(group *jen.Group) {
				for idx, value := range enum.Values {
					group.Case(jen.Lit(idx + 1)).Block(
						jen.Id("out").Op("=").Lit(value.Name),
					)
				}

				group.Default().Block(
					jen.Err().Op("=").Qual("fmt", "Errorf").Call(jen.List(jen.Lit(fmt.Sprintf("unable to parse %%#v into a known %s string", enumGoTypeName)), jen.Id("v"))),
				)
			}),
			jen.Return(),
		)
	} else {
		f.Func().Params(
			jen.Id("v").Id(enumGoTypeName),
		).Id("Value").Params().Params(jen.Qual("database/sql/driver", "Value"), jen.Error()).Block(
			jen.Return(jen.List(jen.Int64().Params(jen.Id("v")), jen.Nil())),
		)
	}

	f.Comment(g.headerComments())

	g.types[enum.Name] = jen.Id(enumGoTypeName)

	return f.Save(filepath.Join(g.out, fmt.Sprintf("%s.enum.go", genutil.Normalize(enum.Name))))
}
