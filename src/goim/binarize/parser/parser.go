package parser

import (
	"go/token"
	"go/parser"
	"go/ast"
)

// This is an entry point to parser package which expects fileName of go file to be passed
// As an output FileModel is returned
func Parse(fileName string) *FileModel {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, fileName, nil, parser.ParseComments)
	if err != nil {
		panic(err.Error())
	}

	visitor := newStructVisitor()
	ast.Walk(visitor, node)

	fileModel := &FileModel{}
	fileModel.initFromAST(node.Name.Name, visitor.structs, visitor.docs)

	return fileModel
}
