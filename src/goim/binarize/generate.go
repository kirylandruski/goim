package binarize

import (
	"io/ioutil"
	"os"
	"goim/binarize/parser"
	"sort"
)

func mapFieldModelToBuilder(model *parser.FieldModel) *structFuncItemBuilder {
	builder := structFuncItemBuilder{
		FieldName: model.Name,
		Kind:      model.Kind,
	}

	return &builder
}

func mapStructModelToBuilder(model *parser.StructModel, serializeFuncName string, deserializeFuncName string) structFuncsBuilder {
	sb := make([]structFuncItemBuilder, 0)
	for _, fieldAppModel := range model.Fields {
		sb = append(sb, *mapFieldModelToBuilder(fieldAppModel))
	}

	return structFuncsBuilder{
		StructId:            model.Id,
		StructName:          model.Name,
		SerializeFuncName:   serializeFuncName,
		DeserializeFuncName: deserializeFuncName,
		Fields:              sb,
	}
}

func mapFleModelToBuilder(importPath string, model *parser.FileModel, serializeFuncName string, deserializeFuncName string) fileBuilder {
	fb := fileBuilder{}
	fb.PackageName = model.PackageName
	fb.Imports = []string{importPath, "errors"}
	fb.ChildBuilders = make([]builder, 0)

	structIds := make([]int, 0, len(model.Structs))
	for k, _ := range model.Structs {
		structIds = append(structIds, int(k))
	}
	sort.Ints(structIds)

	for _, structId := range structIds {
		structModel := model.Structs[int32(structId)]
		sb := mapStructModelToBuilder(structModel, serializeFuncName, deserializeFuncName)
		fb.ChildBuilders = append(fb.ChildBuilders, &sb)
	}

	db := dynamicDeserializeBuilder{DeserializeFuncName: deserializeFuncName}
	db.StructBuilders = make([]*structFuncsBuilder, 0)
	for _, structId := range structIds {
		structModel := model.Structs[int32(structId)]
		sb := mapStructModelToBuilder(structModel, serializeFuncName, deserializeFuncName)
		db.StructBuilders = append(db.StructBuilders, &sb)
	}

	fb.ChildBuilders = append(fb.ChildBuilders, &db)

	ib := interfaceBuilder{SerializeFuncName: serializeFuncName, DeserializeFuncName: deserializeFuncName}
	fb.ChildBuilders = append(fb.ChildBuilders, &ib)

	return fb
}

func Generate(sourceFilePath string, destinationFilePath string, importPath string, serializeFuncName string, deserializeFuncName string) {

	inputStat, err := os.Stat(sourceFilePath)
	if err != nil {
		panic(err.Error())
	}

	m := parser.Parse(sourceFilePath)

	b := mapFleModelToBuilder(importPath, m, serializeFuncName, deserializeFuncName)
	err = ioutil.WriteFile(destinationFilePath, []byte(b.Build()), inputStat.Mode())
	if err != nil {
		panic(err.Error())
	}
}
