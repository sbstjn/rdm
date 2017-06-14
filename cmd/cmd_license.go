package cmd

import (
	"time"

	"github.com/sbstjn/rdm/file"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var licenseMit = file.File{
	Name:     "MIT",
	File:     "LICENSE.md",
	Template: "license/mit.md",
	Fields: []*survey.Question{
		{
			Name:     "Author",
			Prompt:   &survey.Input{Message: "Author:", Default: "Sebastian Müller <mail@sbstjn.com>"},
			Validate: survey.Required,
		},
		{
			Name:     "Year",
			Prompt:   &survey.Input{Message: "Year:", Default: time.Now().Local().Format("2006")},
			Validate: survey.Required,
		},
	},
}

var licenseUnlicense = file.File{
	Name:     "Unlicense",
	File:     "LICENSE.md",
	Template: "license/unlicense.md",
}

var licenseList = []file.File{
	licenseMit,
	licenseUnlicense,
}

func licenseByName(name string) *file.File {
	for _, element := range licenseList {
		if name == element.Name {
			return &element
		}
	}

	return nil
}

func chooseLicense() *file.File {
	var list = []string{}
	for _, element := range licenseList {
		list = append(list, element.Name)
	}

	answers := askForData([]*survey.Question{
		{
			Name: "Type",
			Prompt: &survey.Select{
				Message: "License:",
				Options: list,
				Default: licenseList[0].Name,
			},
		},
	})

	return licenseByName(answers["Type"].(string))
}

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "license",
		Short: "Generate LICENSE.md file",
		Run: func(cmd *cobra.Command, args []string) {
			license := chooseLicense()
			data := askForData(license.Fields)

			err := license.Save(data, cfgOutput, cfgForce)
			if err != nil {
				panic(err)
			}
		},
	})
}
