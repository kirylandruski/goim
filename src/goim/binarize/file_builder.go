package binarize

import (
	"text/template"
	"bytes"
)

type fileBuilder struct {
	PackageName   string
	Imports       []string
	ChildBuilders []builder
}

func (t *fileBuilder) Build() string {
	templateText := `
// THIS IS FILE IS GENERATE BY BINARIZE CODEGEN TOOL
// DO NOT EDIT!
		
package {{.PackageName}}
		
{{range .Imports -}}
import "{{.}}"
{{end -}}

{{- range .ChildBuilders -}}
{{.Build}}
{{- end -}}

`

	tmpl := template.New(t.PackageName)
	if _, err := tmpl.Parse(templateText); err != nil {
		panic(err.Error())
	}

	buf := bytes.NewBuffer(nil)
	if err := tmpl.Execute(buf, t); err != nil {
		panic(err.Error())
	}

	return string(buf.Bytes())
}
