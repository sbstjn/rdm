package file

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"

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

func (f File) exists(path string) bool {
	if _, err := os.Stat(path + f.File); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

// Save the file
func (f File) Save(data map[string]interface{}, path string, force bool) error {
	content := []byte(f.Render(data))

	if !force && f.exists(path) {
		return errors.New("File already exists! Use -f to force overwrite …")
	}

	return ioutil.WriteFile(path+f.File, content, 0644)
}
