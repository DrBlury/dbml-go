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

	f.Type().Id(enumGoTypeName).String()

	f.Const().DefsFunc(func(group *jen.Group) {
		for _, value := range enum.Values {
			capsName := genutil.NormalLizeGoName(value.Name)
			stringEncapsuledEnumValue := fmt.Sprintf("%q", value.Name)
			v := group.Id(capsName).Id(enumGoTypeName).Op("=").Id(stringEncapsuledEnumValue)
			if value.Note != "" {
				v.Comment(value.Note)
			}
		}
	})

	f.Comment(g.headerComments())

	g.types[enum.Name] = jen.Id(enumGoTypeName)

	return f.Save(filepath.Join(g.out, fmt.Sprintf("%s.enum.go", genutil.Normalize(enum.Name))))
}
