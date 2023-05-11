package myanalyzers

import (
	"bytes"
	"go/ast"
	"go/printer"
	"strings"

	"golang.org/x/tools/go/analysis"
)

// Анализатор для проверки os.Exit в main функции
var ExitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for os exit in main",
	Run:  run,
}

// функция для анализатора
func run(pass *analysis.Pass) (interface{}, error) {
	checkExit := func(node ast.Node) bool {
		ast.Inspect(node, func(n ast.Node) bool {
			m, ok := n.(*ast.FuncDecl)
			if ok {
				if m.Name.Name == "main" {
					ast.Inspect(m, func(node ast.Node) bool {
						c, ok := node.(*ast.CallExpr)
						if ok {
							var fName bytes.Buffer
							printer.Fprint(&fName, pass.Fset, c)
							if strings.Contains(fName.String(), "os.Exit") {
								pass.Reportf(c.Pos(), "os.Exit must not call in main")
								return true
							}
						}
						return true
					})
				}
			}
			return true
		})
		return true
	}

	for _, f := range pass.Files {
		for _, decl := range f.Decls {
			if fd, ok := decl.(*ast.FuncDecl); ok {
				if fd.Name.Name == "main" {
					checkExit(f)
				}
			}
		}
	}
	return nil, nil
}
