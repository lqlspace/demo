package addline

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"testing"
)






func TestAddLine(t *testing.T) {
    tests := []testing.InternalTest{{"TestAst", TestAst}}

	testing.Main(matchAll, tests, nil, nil)
}

func matchAll(t string, pat string) (bool, error) {
	return true, nil 
}

func TestAst(t *testing.T) {

    source := `package a

// B comment
type B struct {
    // C comment
    C string
}`

    buffer := make([]byte, 1024, 1024)
    for idx := range buffer {
        buffer[idx] = 0x20
    }
    copy(buffer[:], source)
	
    fset := token.NewFileSet()
    file, err := parser.ParseFile(fset, "", buffer, parser.ParseComments)
    if err != nil {
        t.Error(err)
    }

    v := &visitor{
        file: file,
        fset: fset,
    }
    ast.Walk(v, file)

    var output []byte
    buf := bytes.NewBuffer(output)
    if err := printer.Fprint(buf, fset, file); err != nil {
        t.Error(err)
    }

    expected := `package a

// B comment
type B struct {
    // C comment
    C   string
    // D comment
    D   int
    // E comment
    E   float64
}
`
    if buf.String() != expected {
        t.Errorf(fmt.Sprintf("Test failed. Expected:\n%s\nGot:\n%s", expected, buf.String()))
    }

}
