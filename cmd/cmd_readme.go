package cmd

import (
	"github.com/sbstjn/rdm/file"
	"github.com/spf13/cobra"

	"gopkg.in/AlecAivazis/survey.v1"
)

var readme = file.File{
	File: "README.md",
}

var text = file.Text{
	Name:     "Dummy",
	Template: "# {{ .Project }}{{ range $child := .Content }}{{ if $child }}\n\n{{ $child }}{{ end }}{{ end }}",
	Fields: []*survey.Question{
		{
			Name:     "Project",
			Prompt:   &survey.Input{Message: "Name:", Default: "rdm - ReadMe Scaffolding"},
			Validate: survey.Required,
		},
		{
			Name:     "GitHub",
			Prompt:   &survey.Input{Message: "GitHub:", Default: "sbstjn/rdm"},
			Validate: survey.Required,
		},
	},
	Children: []file.Text{
		file.Text{
			Name: "Shields",
			Fields: []*survey.Question{
				{
					Name: "Shields",
					Prompt: &survey.MultiSelect{
						Message: "Shields:",
						Options: []string{"CircleCI", "GitHub Release", "License"},
						Default: []string{"GitHub Release", "License"},
					},
				},
			},
			File: "readme/shields.md",
		},
		file.Text{
			Template: "This is a description for the project {{ .Project }}.",
		},
		file.Text{
			Name: "License",
			Fields: []*survey.Question{
				{
					Name: "License",
					Prompt: &survey.Select{
						Message: "License:",
						Options: []string{"MIT", "Unlicense"},
					},
				},
			},
			File: "readme/license.md",
		},
		file.Text{
			Name: "Contribution",
			File: "readme/contribution.md",
		},
	},
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "readme",
		Short: "Generate README.md file",
		Run: func(cmd *cobra.Command, args []string) {
			data := map[string]interface{}{}
			content := text.Render(data)

			err := readme.Save(content, cfgOutput, cfgForce)
			if err != nil {
				panic(err)
			}
		},
	})
}
