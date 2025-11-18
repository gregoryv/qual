package qual

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func Test_countlines(t *testing.T) {
	ok := func(src string, exp int) {
		t.Helper()
		// Create the AST by parsing src.
		fset := token.NewFileSet() // positions are relative to fset
		f, err := parser.ParseFile(fset, "src.go", src, 0)
		if err != nil {
			panic(err)
		}

		// Inspect the AST and print all identifiers and literals.
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				got := height(fset, x.Body)
				if got != exp {
					t.Log(src)
					t.Errorf("Got %v expected %v", got, exp)
				}
			}
			return true
		})
	}
	ok(`package x
        func one() {
        	fmt.Println("1")
        }`, 1)

	ok(`package x
        func one() { fmt.Println("1") }`, 1)

	ok(`package x
        func two() {
        	fmt.Println("1")
            fmt.Println("2")
        }`, 2)
}

func TestFuncHeight(t *testing.T) {
	FuncHeight(10, false, &noopT{})
}

func Test_parseFile(t *testing.T) {
	fset := token.NewFileSet()
	parseFile(&noopT{}, fset, "something bad")
}
