package file

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	initTestFS()
}

func initTestFS() {
	appFS = afero.NewMemMapFs()
	appFS.MkdirAll("src/", 0755)
}

func TestFile(t *testing.T) {
	initTestFS()

	readme := File{"ReadMe", "README.md"}

	assert.Equal(t, "ReadMe", readme.Name)
	assert.Equal(t, "README.md", readme.File)
	assert.Equal(t, "/tmp/README.md", readme.absolute("/tmp/"))
}

func TestFileExists(t *testing.T) {
	initTestFS()

	afero.WriteFile(appFS, "/src/LICENSE.md", []byte("Content of LICENSE.md"), 0644)

	readme := File{"ReadMe", "README.md"}
	license := File{"ReadMe", "LICENSE.md"}

	assert.False(t, readme.exists("/src/"))
	assert.True(t, license.exists("/src/"))

	assert.False(t, license.exists("/otherpath/"))
}
