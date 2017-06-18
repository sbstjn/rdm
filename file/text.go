package file

import (
	"bytes"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/alecthomas/template"
)

// Wrapper to prompt for user input using AlecAivazis/survey
func askForData(fields []*survey.Question) map[string]interface{} {
	data := map[string]interface{}{}
	err := survey.Ask(fields, &data)

	if err != nil {
		panic(err.Error())
	}

	return data
}

// Text represents a text block in a file
type Text struct {
	Name     string
	File     string
	Template string
	Fields   []*survey.Question
	Children []Text
}

func (t Text) RenderChildren(data map[string]interface{}) []string {
	var list []string

	for _, child := range t.Children {
		list = append(list, child.Render(data))
	}

	return list
}

func has(list []string, item string) bool {
	for _, str := range list {
		if str == item {
			return true
		}
	}

	return false
}

var funcMap = template.FuncMap{
	"has": has,
}

// Render merges the template with data
func (t Text) Render(data map[string]interface{}) string {
	var templateStr string

	if t.File != "" {
		raw, err := Asset("templates/" + t.File)

		if err != nil {
			panic(err)
		}

		templateStr = string(raw)
	} else {
		templateStr = t.Template
	}

	tmpl, err := template.New(t.Name).Funcs(funcMap).Parse(templateStr)

	if err != nil {
		panic(err)
	}

	for k, v := range askForData(t.Fields) {
		data[k] = v
	}

	if len(t.Children) > 0 {
		data["Content"] = t.RenderChildren(data)
	}

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)

	if err != nil {
		panic(err)
	}

	return strings.Trim(doc.String(), "\n")
}
