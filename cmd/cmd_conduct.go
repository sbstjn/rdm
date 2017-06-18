package cmd

import (
	"github.com/sbstjn/rdm/file"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var conduct = file.File{
	File: "CODE_OF_CONDUCT.md",
}

var conductContent = file.Text{
	Name: "Conduct",
	File: "conduct.md",
	Fields: []*survey.Question{
		{
			Name:     "Mail",
			Prompt:   &survey.Input{Message: "Mail:", Default: "code@sbstjn.com"},
			Validate: survey.Required,
		},
	},
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "conduct",
		Short: "Generate CODE_OF_CONDUCT.md file",
		Run: func(cmd *cobra.Command, args []string) {
			data := map[string]interface{}{}
			content := conductContent.Render(data)

			err := conduct.Save(content, cfgOutput, cfgForce)
			if err != nil {
				panic(err)
			}
		},
	})
}
