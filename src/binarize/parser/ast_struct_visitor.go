package parser

import (
	"go/ast"
	"go/token"
)

type structVisitor struct {
	structs map[string]*ast.StructType
	docs    map[string]*ast.CommentGroup

	parentDoc     *ast.CommentGroup
	parentVisitor *structVisitor
}

func newStructVisitor() *structVisitor {
	v := structVisitor{}
	v.structs = make(map[string]*ast.StructType)
	v.docs = make(map[string]*ast.CommentGroup)

	return &v
}

func (v *structVisitor) newChild() *structVisitor {
	child := newStructVisitor()
	child.parentVisitor = v

	return child
}

func (v *structVisitor) pushStruct(name string, structType *ast.StructType, docs *ast.CommentGroup) {
	root := v
	for root.parentVisitor != nil {
		root = root.parentVisitor
	}

	root.structs[name] = structType
	root.docs[name] = docs
}

func (v *structVisitor) Visit(node ast.Node) ast.Visitor {
	switch node := node.(type) {
	case *ast.File:
		return v
	case *ast.GenDecl:
		if node.Tok == token.TYPE {
			child := v.newChild()
			child.parentDoc = node.Doc

			return child
		}
	case *ast.TypeSpec:
		structType, ok := node.Type.(*ast.StructType)
		if ok {
			v.pushStruct(node.Name.Name, structType, v.parentDoc)
		}
	}

	return nil
}
