package file

import (
	"bytes"

	"gopkg.in/AlecAivazis/survey.v1"

	"text/template"
)

// File handles basic files
type File struct {
	Name     string
	File     string
	Template string
	Fields   []*survey.Question
}

// Render returns the content for the file
func (f File) Render(data map[string]interface{}) string {
	raw, err := Asset("templates/" + f.Template)

	if err != nil {
		panic(err)
	}

	tmpl, err := template.New(f.File).Parse(string(raw))

	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)

	if err != nil {
		panic(err)
	}

	return doc.String()
}
