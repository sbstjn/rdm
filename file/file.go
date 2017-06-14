package file

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"gopkg.in/AlecAivazis/survey.v1"

	"text/template"
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

	data["Content"] = t.RenderChildren(data)

	var doc bytes.Buffer
	err = tmpl.Execute(&doc, data)

	if err != nil {
		panic(err)
	}

	return strings.Trim(doc.String(), "\n")
}

// File handles basic files
type File struct {
	Name     string
	File     string
	Template string
	Fields   []*survey.Question
}

func (f File) exists(path string) bool {
	if _, err := os.Stat(path + f.File); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// Save the file
func (f File) Save(content string, path string, force bool) error {
	if !force && f.exists(path) {
		return errors.New("File already exists! Use -f to force overwrite …")
	}

	return ioutil.WriteFile(path+f.File, []byte(content), 0644)
}
