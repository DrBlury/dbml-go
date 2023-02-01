package gen

import (
	"regexp"

	"github.com/dave/jennifer/jen"
)

const primeTypePattern = `^(\w+)(\(d+\))?`

var (
	regexType    = regexp.MustCompile(primeTypePattern)
	builtinTypes = map[string]jen.Code{
		"int":       jen.Int(),
		"integer":   jen.Int(),
		"int8":      jen.Int8(),
		"int16":     jen.Int16(),
		"int32":     jen.Int32(),
		"int64":     jen.Int64(),
		"bigint":    jen.Int64(),
		"uint":      jen.Uint(),
		"uint8":     jen.Uint8(),
		"uint16":    jen.Uint16(),
		"uint32":    jen.Uint32(),
		"uint64":    jen.Uint64(),
		"float":     jen.Float64(),
		"float32":   jen.Float32(),
		"float64":   jen.Float64(),
		"bool":      jen.Bool(),
		"boolean":   jen.Bool(),
		"text":      jen.String(),
		"varchar":   jen.String(),
		"char":      jen.String(),
		"byte":      jen.Byte(),
		"rune":      jen.Rune(),
		"timestamp": jen.Qual("time", "Time"),
		"datetime":  jen.Qual("time", "Time"),
	}
)

func (g *generator) getJenType(s string) (jen.Code, bool) {
	m := regexType.FindStringSubmatch(s)
	if len(m) >= 2 {
		// lookup for builtin type
		if t, ok := builtinTypes[m[1]]; ok {
			return t, ok
		}
	}
	t, ok := g.types[s]
	return t, ok
}
