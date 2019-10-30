package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	//"go/printer"
	"go/token"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

var src = `package model
type (
	Model struct {
		A,B int `+"`put_option:\"\"`"+`
		A,B int `+"`put_option:\"\"`"+`
		A struct{}
	}
	ModelB struct {
		A int
	}
)

type ModelC struct {
	A int
}
`


type PutOption struct {

}

func parsePutOption(tag reflect.StructTag) *PutOption {
	if value, ok := tag.Lookup("put_option"); ok {
		return nil
	} else {
		options := strings.Split(value, ";")
		option := &PutOption{}
		for i := range options {
			_ = options[i]
		}
		return option
	}
}

func Gen(ostream io.Writer, modelName string) (err error) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return err
	}
	var decl *ast.TypeSpec

	for i := range f.Decls {
		if dec, ok := f.Decls[i].(*ast.GenDecl); ok {
			if dec.Tok != token.TYPE {
				continue
			}
			for j := range dec.Specs {
				if spec, ok := dec.Specs[j].(*ast.TypeSpec); ok {
					if spec.Name.Name == modelName {
						if decl != nil {
							return fmt.Errorf("redefinition of model Struct in file")
						}
						decl = spec
					}
				}
			}
		}
	}

	if decl == nil {
		return fmt.Errorf("aiming model not found")
	}

	template := NewPutTemplate()

	if structType, ok := decl.Type.(*ast.StructType); ok {
		fmt.Println(structType)
		for i := range structType.Fields.List {
			var option *PutOption
			if structType.Fields.List[i].Tag != nil {
				option = parsePutOption(reflect.StructTag(structType.Fields.List[i].Tag.Value))
			}
			if option == nil {
				continue
			}

			if fieldType, ok := structType.Fields.List[i].Type.(*ast.Ident); ok {
				switch fieldType.Name {
				case "int", "int8", "int16", "int32", "int64",
					"uint", "uint8", "uint16", "uint32", "uint64",
					"float32", "float64":
					template.AppendNumberField(structType.Fields.List[i].Names...)
				}
			}
		}
	} else {
		return fmt.Errorf("model decl %T is not the struct type", decl.Type)
	}

	return template.FPrint(ostream)
}

func main() {
	if err := Gen(os.Stdout, "Model"); err != nil {
		log.Fatal(err)
	}
}
