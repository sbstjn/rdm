package cmd

import (
	"time"

	"github.com/sbstjn/rdm/file"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

var licenseFile = file.File{
	File: "LICENSE.md",
}

var licenseMit = file.Text{
	Name: "MIT",
	File: "license/mit.md",
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

var licenseUnlicense = file.Text{
	Name: "Unlicense",
	File: "license/unlicense.md",
}

var licenseList = []file.Text{
	licenseMit,
	licenseUnlicense,
}

func licenseByName(name string) *file.Text {
	for _, element := range licenseList {
		if name == element.Name {
			return &element
		}
	}

	return nil
}

func chooseLicense() *file.Text {
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
			data := map[string]interface{}{}

			license := chooseLicense()
			content := license.Render(data)

			err := licenseFile.Save(content, cfgOutput, cfgForce)
			if err != nil {
				panic(err)
			}
		},
	})
}
