package gen

import (
	"dbml-go/core"

	"github.com/dave/jennifer/jen"
)

const (
	version = "v1.0.0"
)

type generator struct {
	dbml             *core.DBML
	out              string
	gopackage        string
	fieldtags        []string
	types            map[string]jen.Code
	shouldGenTblName bool
	isPostgres       bool
}

func newgen() *generator {
	return &generator{
		types: make(map[string]jen.Code),
	}
}

func (g *generator) reset(rememberAlias bool) {
	g.dbml = nil
	if !rememberAlias {
		g.types = make(map[string]jen.Code)
	}
}

func (g *generator) file() *jen.File {
	return jen.NewFilePathName(g.out, g.gopackage)
}

func (g *generator) generate() (err error) {
	for _, enum := range g.dbml.Enums {
		if err = g.genEnum(enum); err != nil {
			return
		}
	}
	for _, table := range g.dbml.Tables {
		if err = g.genTable(table); err != nil {
			return err
		}
	}

	return nil
}
