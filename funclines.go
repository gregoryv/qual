package qual

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"
)

// FuncHeight fails if any func body exceeds maxLines.
// All lines are considered, source and comments.
func FuncHeight(maxLines int, includeVendor bool, t T) {
	t.Helper()
	files := findGoFiles(includeVendor)
	result := []string{}
	for _, file := range files {
		fset := token.NewFileSet()
		f := parseFile(t, fset, file)

		ast.Inspect(f, func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.FuncDecl:
				got := height(fset, x.Body)
				if got > maxLines {
					r := fmt.Sprintf(
						"%s:%v: %s",
						file, fset.Position(x.Name.Pos()).Line, x.Name,
					)
					result = append(result, r)
				}
			}
			return true
		})
	}
	if len(result) > 0 {
		t.Errorf("Func height %v exceeded in\n%s", maxLines,
			strings.Join(result, "\n"),
		)
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
