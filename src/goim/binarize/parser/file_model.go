package parser

import "go/ast"

// Describes a whole file containing struct models
type FileModel struct {
	PackageName string
	Structs     map[int32]*StructModel
}

func (m *FileModel) initFromAST(packageName string, structs map[string]*ast.StructType, docs map[string]*ast.CommentGroup) {
	m.PackageName = packageName
	m.Structs = make(map[int32]*StructModel)

	for name, structType := range structs {
		structModel := &StructModel{}
		structModel.initFromAST(name, structType, docs[name])
		if _, ok := m.Structs[structModel.Id]; ok {
			panic("two structs have the same id")
		}
		m.Structs[structModel.Id] = structModel
	}
}
