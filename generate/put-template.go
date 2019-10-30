package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"reflect"
)

var putFunction = `
package funcSet
func (srv *Service) fillPutFields(object *model.Object, req *PutRequest) (fields []string) {
	if len(req.A) != 0 {
		object.A = req.A
		field = append(field, "a")
	}
	if req.A != 0 {
		object.A = req.A
		field = append(field, "a")
	}
	if req.A != nil {
		object.A = req.A
		field = append(field, "a")
	}
	return nil
}
`

var namespace struct {
	fset *token.FileSet
	PutTemplate struct {
		template              *ast.FuncDecl
		stringAssignTemplate  *ast.IfStmt
		numberAssignTemplate  *ast.IfStmt
		pointerAssignTemplate *ast.IfStmt
	}
}

type PutTemplate struct {
	template *ast.FuncDecl
}

func (template *PutTemplate) FPrint(w io.Writer) error {
	return printer.Fprint(w, namespace.fset, template.template)
}

func generateNumberAssignStmt(ident *ast.Ident) ast.Stmt {
	var template = &ast.IfStmt{
		Init: nil,
		Cond: &ast.BinaryExpr{
			X:     &ast.CallExpr{
				Fun:      &ast.Ident{
					NamePos: 0,
					Name:    "len",
					Obj:     nil,
				},
				Args:     []ast.Expr{
					&ast.SelectorExpr{
						X:   &ast.Ident{
							Name:    "req",
						},
						Sel: &ast.Ident{
							Name:    ident.Name,
						},
					},
				},
			},
			Op:    token.NEQ,
			Y:     &ast.BasicLit{
				Kind:     token.INT,
				Value:    "0",
			},
		},
		Body: &ast.BlockStmt{
			List:   []ast.Stmt{
				&ast.AssignStmt{
					Lhs:    []ast.Expr{
						&ast.SelectorExpr{
							X:   &ast.Ident{
								Name:    "object",
							},
							Sel: &ast.Ident{
								Name:    ident.Name,
							},
						},
					},
					Tok:    token.EQL,
					Rhs:    []ast.Expr{
						&ast.SelectorExpr{
							X:   ast.NewIdent("req"),
							Sel: ast.NewIdent("A"),
						},
					},
				},
				&ast.AssignStmt{
					Lhs:    []ast.Expr{
						ast.NewIdent("field"),
					},
					Tok:    token.EQL,
					Rhs:    []ast.Expr{
						&ast.CallExpr{
							Fun:      ast.NewIdent("append"),
							Lparen:   0,
							Args:     []ast.Expr{
								ast.NewIdent("field"),
								ast.BasicLit{
									Kind:     0,
									Value:    "",
								}
							},
							Ellipsis: 0,
							Rparen:   0,
						}
					},
				}
			},
			Rbrace: 0,
		},
		Else: nil,
	}
	//*template = *
	fmt.Println(template.Cond)

	ast.Inspect(template, func(node ast.Node) bool {
		if aIdent, ok := node.(*ast.Ident); ok && aIdent.Name == "A" {
			*aIdent = *ident
		}
		return true
	})
	_ = ast.Print(namespace.fset, template.Cond)
	return template
}

func (template *PutTemplate) AppendNumberField(idents ...*ast.Ident) {
	for i := range idents {
		returnStmt := template.template.Body.List[
			len(template.template.Body.List) - 1]
		template.template.Body.List[
			len(template.template.Body.List) - 1] =
				generateNumberAssignStmt(idents[i])
		template.template.Body.List = append(template.template.Body.List,
			returnStmt)
	}
}


func NewPutTemplate() *PutTemplate {
	template := &ast.FuncDecl{}
	*template = *namespace.PutTemplate.template
	template.Body = &ast.BlockStmt{}
	*template.Body = *namespace.PutTemplate.template.Body
	template.Body.List = make([]ast.Stmt, len(template.Body.List))
	copy(template.Body.List, namespace.PutTemplate.template.Body.List)

	return &PutTemplate{
		template: template,
	}
}


func init() {
	fset := token.NewFileSet()
	pfile, err := parser.ParseFile(fset, "", putFunction, 0)
	if err != nil {
		panic(err)
	}

	err = ast.Fprint(os.Stdout, fset, pfile, nil)
	if err != nil {
		panic(err)
	}

	template := pfile.Decls[0].(*ast.FuncDecl)
	stringAssignTemplate := template.Body.List[0].(*ast.IfStmt)
	template.Body.List = template.Body.List[1:]
	numberAssignTemplate := template.Body.List[0].(*ast.IfStmt)
	template.Body.List = template.Body.List[1:]
	pointerAssignTemplate := template.Body.List[0].(*ast.IfStmt)
	template.Body.List = template.Body.List[1:]

	namespace.PutTemplate.template = template
	namespace.PutTemplate.stringAssignTemplate = stringAssignTemplate
	namespace.PutTemplate.numberAssignTemplate = numberAssignTemplate
	namespace.PutTemplate.pointerAssignTemplate = pointerAssignTemplate
	namespace.fset = fset
}
