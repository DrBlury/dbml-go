package gen

import (
	"fmt"
	"time"
)

func (g *generator) headerComments() string {
	return fmt.Sprintf(`Generated by dbml-go
    version: %s
    timestamp: %s
`, version, time.Now().String())
}
