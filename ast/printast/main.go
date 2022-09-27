package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)


func main() {
	fset := token.NewFileSet()

	astFile, err := parser.ParseFile(fset, "test.go", nil, 0)
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("tmp.txt")

	ast.Fprint(f, fset, astFile, ast.NotNilFilter)
}