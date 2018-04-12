package binarize

import (
	"text/template"
	"bytes"
)

type structFuncsBuilder struct {
	StructId            int32
	StructName          string
	SerializeFuncName   string
	DeserializeFuncName string
	UtilsReference      string
	Fields              []structFuncItemBuilder
}

func (t *structFuncsBuilder) Build() string {
	var templateText = `
func (s *{{.StructName}}) {{.SerializeFuncName}}() []byte {
	offset := int32(0)
	buf := make([]byte, binarize.IdSize(0){{range .Fields}}+binarize.{{.SizeFuncName}}(s.{{.FieldName}}){{end}})
	offset = binarize.IdPut(buf, offset, {{.StructId}})
	{{range .Fields}}
	offset = binarize.{{.PutFuncName}}(buf, offset, s.{{.FieldName}})
	{{- end}}

	return buf
}

func (s *{{.StructName}}) {{.DeserializeFuncName}}(buf []byte) error {
	var err error
	offset := int32(0)
	offset, err = binarize.IdValidate(buf, offset, {{.StructId}})
	if err != nil {
		return err
	}
	{{range .Fields}}
	s.{{.FieldName}}, offset, err = binarize.{{.ReadFuncName}}(buf, offset)
	if err != nil {
		return err
	}
	{{- end}}

	return nil
}
`

	tmpl := template.New(t.StructName)
	if _, err := tmpl.Parse(templateText); err != nil {
		panic(err.Error())
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, t); err != nil {
		panic(err.Error())
	}

	return string(buf.Bytes())
}

type structFuncItemBuilder struct {
	FieldName string
	Kind      string
}

func (f *structFuncItemBuilder) ReadFuncName() string {
	fname, ok := map[string]string{
		"int32":              "Int32Read",
		"int64":              "Int64Read",
		"*string":            "StrRead",
		"[]*string":          "ArrRead",
		"map[string]*string": "MapRead",
	}[f.Kind]

	if ok != true {
		panic(f.Kind + " in unsupported. check comments for more info")
	}

	return fname
}

func (f *structFuncItemBuilder) PutFuncName() string {
	fname, ok := map[string]string{
		"int32":              "Int32Put",
		"int64":              "Int64Put",
		"*string":            "StrPut",
		"[]*string":          "ArrPut",
		"map[string]*string": "MapPut",
	}[f.Kind]

	if ok != true {
		panic(f.Kind + " in unsupported. check comments for more info")
	}

	return fname
}

func (f *structFuncItemBuilder) SizeFuncName() string {
	fname, ok := map[string]string{
		"int32":              "Int32Size",
		"int64":              "Int64Size",
		"*string":            "StrSize",
		"[]*string":          "ArrSize",
		"map[string]*string": "MapSize",
	}[f.Kind]

	if ok != true {
		panic(f.Kind + " in unsupported. check comments for more info")
	}

	return fname
}
