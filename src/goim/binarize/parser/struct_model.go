package parser

import (
	"go/ast"
	"regexp"
	"strconv"
)

// Describes a struct 
type StructModel struct {
	Id     int32
	Name   string
	Fields []*FieldModel
}

func (m *StructModel) initFromAST(name string, structType *ast.StructType, doc *ast.CommentGroup) {
	m.Name = name
	m.Id = parseStructId(doc)

	fields := make([]*FieldModel, len(structType.Fields.List))
	for _, field := range structType.Fields.List {
		order := parseFieldOrder(field.Doc)
		if order >= int32(len(fields)) {
			panic("field order is bigger than Fields count in struct")
		}
		if fields[order] != nil {
			panic("field order is duplicated by another field in that struct")
		}

		fieldModel := &FieldModel{}
		fieldModel.initFromAST(field.Names[0].Name, field.Type)
		fields[order] = fieldModel
	}

	m.Fields = fields
}

func parseStructId(doc *ast.CommentGroup) int32 {
	order, ok := extractInt32(doc, `^//\s*binarize\s*id\s*(\d+)\s*$`) // matches a one line comment like `// binarize Id 65`
	if !ok {
		panic("each struct should have a binarize comment describing struct Id. see docs for more details")
	}

	return order
}

func parseFieldOrder(doc *ast.CommentGroup) int32 {
	order, ok := extractInt32(doc, `^//\s*binarize\s*order\s*(\d+)\s*$`) // matches a one line comment like `// binarize order 65`
	if !ok {
		panic("each field should have a binarize comment describing serialization order. see docs for more details")
	}

	return order
}

func extractInt32(doc *ast.CommentGroup, pattern string) (int32, bool) {
	if doc == nil {
		return 0, false
	}

	reg := regexp.MustCompile(pattern)
	for _, comment := range doc.List {
		if reg.MatchString(comment.Text) {
			match := reg.FindStringSubmatch(comment.Text)[1]
			res, _ := strconv.ParseInt(match, 0, 32)
			return int32(res), true
		}
	}

	return 0, false
}
