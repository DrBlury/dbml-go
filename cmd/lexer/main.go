package main

import (
	"fmt"

	"os"

	token "dbml-go/token"

	scanner "dbml-go/scanner"
)

func main() {
	f2, _ := os.Open("test.dbml")
	s := scanner.NewScanner(f2)
	for {
		l, c := s.LineInfo()
		tok, lit := s.Read()
		if tok == token.EOF {
			break
		}
		fmt.Printf("[%d:%d]\tToken: %s\tlit: %s\n", l, c, tok, lit)
	}
}
