package binarize

import (
	"text/template"
	"bytes"
)

type dynamicDeserializeBuilder struct {
	DeserializeFuncName string
	StructBuilders      []*structFuncsBuilder
}

func (t *dynamicDeserializeBuilder) Build() string {
	templateText := `

func Dynamic{{.DeserializeFuncName}}(buf []byte) (Binarizer, error) {
	var err error
	offset := int32(0)
	id, offset, err := binarize.IdRead(buf, offset)
	if err != nil {
		return nil, err
	}

	switch id {
 	{{range .StructBuilders}}
	case {{.StructId}}:
		res := &{{.StructName}}{}
		err := res.{{.DeserializeFuncName}}(buf)
		if err != nil {
			return nil, err
		}
		return res, nil
	{{end}}
	default:
		return nil, errors.New("could not parse struct - unexpected struct id")
	}
}

`

	tmpl := template.New("")
	if _, err := tmpl.Parse(templateText); err != nil {
		panic(err.Error())
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, t); err != nil {
		panic(err.Error())
	}

	return string(buf.Bytes())
}
