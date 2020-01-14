package artist

import "fmt"

type FunctionTemplateFactory = func(template ObjectTemplate, addition TemplateContext) []FunctionTemplate

type FunctionTemplateMiddleware = func(template FunctionTemplate, addition TemplateContext) FunctionTemplate

// func (:receiver) :FunctionName(:XParams) :ReturnType {
//     :statements
//     return
// }

type FunctionTemplate interface {
	fmt.Stringer
	GetResponsibleObject() ObjectTemplate
}
