package gen

import (
	"dbml-go/core"
	"dbml-go/internal/gen-go-model/genutil"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dave/jennifer/jen"
)

func (g *generator) genTable(table core.Table) error {
	f := jen.NewFilePathName(g.out, g.gopackage)

	tableOriginName := genutil.Normalize(table.Name)
	tableGoTypeName := genutil.NormalizeGoTypeName(table.Name)

	f.Commentf("%s is generated type for table '%s'", tableGoTypeName, tableOriginName)

	var genColumnErr error

	cols := make([]string, 0)

	f.Type().Id(tableGoTypeName).StructFunc(func(group *jen.Group) {
		for _, column := range table.Columns {
			columnName := genutil.NormalLizeGoName(column.Name)
			columnOriginName := genutil.Normalize(column.Name)
			t, ok := g.getJenType(column.Type)
			if !ok {
				genColumnErr = fmt.Errorf("type '%s' is not support", column.Type)
			}
			if column.Settings.Note != "" {
				group.Comment(column.Settings.Note)
			}

			gotags := make(map[string]string)
			for _, t := range g.fieldtags {
				gotags[strings.TrimSpace(t)] = columnOriginName
			}
			group.Id(columnName).Add(t).Tag(gotags)
			cols = append(cols, columnOriginName)
		}
	})

	// ###############################
	// table metadata
	tableMetadataType := "__tbl_" + tableOriginName
	tableMetadataColumnsType := tableMetadataType + "_columns"

	f.Commentf("// table '%s' columns list struct", tableOriginName)
	f.Type().Id(tableMetadataColumnsType).StructFunc(func(group *jen.Group) {
		for _, column := range table.Columns {
			group.Id(genutil.NormalLizeGoName(column.Name)).String()
		}
	})

	f.Commentf("// table '%s' metadata struct", tableOriginName)
	f.Type().Id("__tbl_"+tableOriginName).Struct(
		jen.Id("Name").String(),
		jen.Id("Columns").Id(tableMetadataColumnsType),
	)

	tableMetadataVar := "_tbl_" + tableOriginName

	f.Commentf("// table '%s' metadata info", tableOriginName)
	f.Var().Id(tableMetadataVar).Op("=").Id(tableMetadataType).Values(jen.DictFunc(func(d jen.Dict) {
		d[jen.Id("Name")] = jen.Lit(tableOriginName)
		d[jen.Id("Columns")] = jen.Id(tableMetadataColumnsType).Values(jen.DictFunc(func(d jen.Dict) {
			for _, column := range table.Columns {
				columnName := genutil.NormalLizeGoName(column.Name)
				columnOriginName := genutil.Normalize(column.Name)
				d[jen.Id(columnName)] = jen.Lit(columnOriginName)
			}
		}))
	}))

	f.Commentf("GetColumns return list columns name for table '%s'", tableOriginName)
	f.Func().Params(
		jen.Op("*").Id(tableMetadataType),
	).Id("GetColumns").Params().Index().String().Block(
		jen.Return(jen.Index().String().ValuesFunc(func(g *jen.Group) {
			for _, col := range table.Columns {
				g.Lit(col.Name)
			}
		})),
	)

	f.Commentf("T return metadata info for table '%s'", tableOriginName)
	f.Func().Params(
		jen.Op("*").Id(tableGoTypeName),
	).Id("T").Params().Op("*").Id(tableMetadataType).Block(
		jen.Return().Op("&").Id(tableMetadataVar),
	)

	if g.shouldGenTblName {
		f.Commentf("TableName return table name")
		f.Func().Params(
			jen.Id(tableGoTypeName),
		).Id("TableName").Params().Id("string").Block(
			jen.Return(jen.Lit(tableOriginName)),
		)
	}

	f.Comment(g.headerComments())

	if genColumnErr != nil {
		return genColumnErr
	}

	return f.Save(filepath.Join(g.out, fmt.Sprintf("%s.table.go", genutil.Normalize(table.Name))))
}
