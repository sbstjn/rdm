package file

import (
	"errors"
	"io/ioutil"

	"github.com/spf13/afero"
)

var appFS = afero.NewOsFs()

// File handles basic files
type File struct {
	Name string
	File string
}

func (f File) absolute(path string) string {
	return path + f.File
}

func (f File) exists(path string) bool {
	if _, err := appFS.Stat(f.absolute(path)); err != nil {
		return false
	}

	return true
}

// Save the file
func (f File) Save(content string, path string, force bool) error {
	if !force && f.exists(path) {
		return errors.New("File already exists! Use -f to force overwrite …")
	}

	return ioutil.WriteFile(f.absolute(path), []byte(content), 0644)
}
