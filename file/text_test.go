package file

import (
	"strings"
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

func TestTextHelperHas(t *testing.T) {
	list := []string{
		"Test",
		"Lorem",
		"Ipsum",
	}

	assert.True(t, has(list, "Lorem"))
	assert.False(t, has(list, "Dolor"))
}

func TestTextTemplate(t *testing.T) {
	text := Text{
		Name: "Test Name",
		File: "conduct.md",
	}

	assert.Equal(t, "Test Name", text.Name)
	assert.True(t, strings.HasPrefix(text.Render(nil), "# Contributor Covenant Code of Conduct"))
}

func TestChildren(t *testing.T) {
	text := Text{
		Name:     "Test Name",
		Template: "{{ range $child := .Content }}{{ $child }}{{ end }}",
	}

	assert.Equal(t, "Test Name", text.Name)
	assert.Equal(t, "", text.Render(nil))
}

func TestChildren2(t *testing.T) {
	text := Text{
		Name:     "Test Name",
		Template: "Data {{ .Name }} and children: {{ range $child := .Content }}{{ if $child }}{{ $child }}{{ end }}{{ end }}",
		Children: []Text{
			Text{
				Template: "Nested {{ .Name }}",
			},
		},
	}

	data := map[string]interface{}{
		"Name": "Example",
	}

	assert.Equal(t, "Test Name", text.Name)
	assert.Equal(t, "Data Example and children: Nested Example", text.Render(data))
}
