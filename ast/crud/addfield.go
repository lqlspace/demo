package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
)

func main() {
	fset := token.NewFileSet()

	astFile, err := parser.ParseFile(fset, "test.go", nil, 0)
	if err != nil {
		panic(err)
	}

	v := new(ExistVisitor)
	for _, item := range []string{"Student", "teacher"} {
		v.obj, v.exist = item, false
		ast.Walk(v, astFile)
		fmt.Printf("Student exist in test.go: %v\n", v.exist)
	}

	printer.Fprint(os.Stdout, fset, astFile)
}

type ExistVisitor struct {
	obj string 
	exist bool 
}

func (l *ExistVisitor) Visit(node ast.Node) ast.Visitor {
	if v, ok := node.(*ast.TypeSpec); ok {
		if v.Name.Name != l.obj {
			return l 
		}

		l.exist = true
		if val, ok := v.Type.(*ast.StructType); ok {
			val.Fields.List = append(val.Fields.List, &ast.Field{
				Names: []*ast.Ident{
					{
						Name: "Age",
					},
				},
				Type: &ast.Ident{
					Name: "int",
				},
				Tag: &ast.BasicLit{
					Kind: token.INT,
					Value: `json:"age"`,
				},
			})
		}
	}
	return l 
}

