package binarize

import (
	"bytes"
	"text/template"
)

type interfaceBuilder struct {
	SerializeFuncName   string
	DeserializeFuncName string
}

func (t *interfaceBuilder) Build() string {
	var templateText = `
type Binarizer interface {
	{{.SerializeFuncName}}() []byte
	{{.DeserializeFuncName}}(buf []byte) error
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
