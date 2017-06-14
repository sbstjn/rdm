package cmd

import (
	"github.com/sbstjn/rdm/file"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var conduct = file.File{
	Name:     "Code of Conduct",
	File:     "CODE_OF_CONDUCT.md",
	Template: "conduct.md",
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
			data := askForData(conduct.Fields)

			err := conduct.Save(data, cfgOutput, cfgForce)
			if err != nil {
				panic(err)
			}
		},
	})
}
