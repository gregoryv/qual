package qual

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
)

// LineLength fails if any go file contains lines exceeding maxChars.
// All lines are considered, source and comments.
func FuncHeight(maxLines int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	for _, file := range files {
		fset := token.NewFileSet()
		f := parseFile(t, fset, file)

		// Inspect the AST and print all identifiers and literals.
		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				got := height(fset, x.Body)
				if got > maxLines {
					t.Errorf("Got %v", got)
				}
			}
			return true
		})
	}
}

func parseFile(t T, fset *token.FileSet, file string) *ast.File {
	t.Helper()
	var src interface{}
	if path.Ext(file) != ".go" {
		src = file
		file = "source.go"
	}
	f, err := parser.ParseFile(fset, file, src, 0)
	if err != nil {
		t.Error(err)
	}
	return f
}

func height(fset *token.FileSet, bl *ast.BlockStmt) int {
	right := fset.Position(bl.Rbrace).Line - 1
	left := fset.Position(bl.Lbrace).Line
	n := right - left
	if n <= 1 {
		n = 1
	}
	return n
}
