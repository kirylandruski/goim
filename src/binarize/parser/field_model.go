package parser

import (
	"go/ast"
	"reflect"
)

// Describes a field in a structure
type FieldModel struct {
	Name string
	Kind string
}

func (m *FieldModel) initFromAST(name string, node ast.Expr) {
	m.Name = name
	m.Kind = parseKind(node)
}

func parseKind(node ast.Expr) string {
	switch expr := node.(type) {
	case *ast.Ident:
		obj := expr.Obj
		if obj != nil {
			spec, ok := obj.Decl.(*ast.TypeSpec)
			if ok {
				return parseKind(spec.Type)
			}
			return obj.Name
		}
		return expr.Name
	case *ast.StarExpr:
		return "*" + parseKind(expr.X.(*ast.Ident))
	case *ast.MapType:
		return "map" + "[" + parseKind(expr.Key) + "]" + parseKind(expr.Value)
	case *ast.ArrayType:
		return "[]" + parseKind(expr.Elt)
	default:
		panic(reflect.ValueOf(expr).String() + " is unsupported. only basic types, maps and slices are supported. see docs for more details")
	}
}
