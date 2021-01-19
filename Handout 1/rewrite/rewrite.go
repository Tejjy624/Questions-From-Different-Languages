package rewrite

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"hw1/expr"
	"hw1/simplify"
	//"hw1/expr/eval.go"
	"strconv" // This package may be helpful...
)

// rewriteCalls should modify the passed AST
func rewriteCalls(node ast.Node) {
	//TODO Write the rewriteCalls function
	//panic("TODO: implement this!")
	//string_value := node
	//e, err := expr.Parse(node)
	//simplify.simplify(node)
	ast.Inspect(node, func(node ast.Node) bool {
		var s string
		switch n := node.(type) {
		case *ast.CallExpr:
			if len(n.Args) == 2{
				switch ct := n.Fun.(type){
				case ast.Expr:
					TC, check := ct.(*ast.SelectorExpr)
					if (check){
						switch ctname := TC.X.(type){
						case ast.Expr:
							Tok, checkother := ctname.(*ast.Ident)
							if (checkother){
								if Tok.Name == "eval" && TC.Sel.Name == "ParseAndEval" {
									switch argument_type := n.Args[0].(type){
									case ast.Expr:
										Z, checklast := argument_type.(*ast.BasicLit)
										if (checklast){
											arg1,_ := strconv.Unquote(Z.Value)
											expr_arg, err := expr.Parse(arg1)
											if (err == nil){
												result := simplify.Simplify(expr.Env{}, expr.Env)
												str1 := strconv.Quote(expr.Format(result))
												Z.Value = str1
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	return true
	})

	//ast.Inspect(node, func(node ast.Node) bool{
	//	switch n := node.(type)
	//	case *ast.CallExpr {
	//		if len(n.Args) == 2{
	//			i := n.Fun.(type)
	//			//if i := ast.Expr:
	//		}
	//	}
	//})
}

func SimplifyParseAndEval(src string) string {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	rewriteCalls(f)

	var buf bytes.Buffer
	format.Node(&buf, fset, f)
	return buf.String()
}
