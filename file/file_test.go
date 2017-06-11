package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile(t *testing.T) {
	readme := File{"ReadMe", "README.md", "", nil}

	assert.Equal(t, "ReadMe", readme.Name)
	assert.Equal(t, "README.md", readme.File)
}
