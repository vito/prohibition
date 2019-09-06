package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"

	"github.com/fatih/astrewrite"
)

var write = flag.Bool(
	"w",
	false,
	"overwrite files",
)

func main() {
	flag.Parse()

	files := flag.Args()
	fset := token.NewFileSet()

	for _, filename := range files {
		file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
		if err != nil {
			log.Fatal(err)
		}

		rewritten := astrewrite.Walk(file, func(n ast.Node) (ast.Node, bool) {
			// is it a function?
			f, ok := n.(*ast.FuncDecl)
			if !ok {
				return n, true
			}

			// are there any parameters?
			params := f.Type.Params.List
			if len(params) == 0 {
				return n, true
			}

			// is the first parameter a logger?
			if !isLogger(params[0]) {
				return n, true
			}

			// remove it.
			f.Type.Params.List = f.Type.Params.List[1:]
			return n, false
		})

		var buf bytes.Buffer
		printer.Fprint(&buf, fset, rewritten)

		if *write {
			if err := ioutil.WriteFile(filename, buf.Bytes(), 0644); err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println(buf.String())
		}
	}
}

func isLogger(param *ast.Field) bool {
	a, ok := param.Type.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	pkg := a.X.(*ast.Ident).Name
	typ := a.Sel.Name

	// BUG: Someone could pathologically import another package as lager with a
	// type called logger in it. For complete correctness we should exract the
	// expect package name from the imports in the file.
	return pkg == "lager" && typ == "Logger"
}
