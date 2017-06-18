package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestText(t *testing.T) {
	text := Text{
		Name:     "Test Name",
		Template: "This is a template",
	}

	assert.Equal(t, "Test Name", text.Name)
	assert.Equal(t, "This is a template", text.Render(nil))
}
