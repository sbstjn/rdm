package cmd

import (
	"fmt"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/AlecAivazis/survey.v1"
)

var cfgFile string
var cfgOutput string
var cfgForce bool
var version string

var rootCmd = &cobra.Command{
	Use: "rdm",
}

// Run initializes the rdm command
func Run(v string) {
	version = v

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(config)

	rootCmd.PersistentFlags().BoolVarP(&cfgForce, "force", "f", false, "Overwrite files")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rdm.yaml)")
	rootCmd.PersistentFlags().StringVarP(&cfgOutput, "out", "o", "", "output path (default is current directory)")
}

func resolveOutputDirectory() {
	// Prepand current directory if path is relative or none is set
	if cfgOutput == "" || !strings.HasPrefix(cfgOutput, "/") {
		dir, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		cfgOutput = dir + "/" + cfgOutput
	}

	// Add trailing slash to path
	if !strings.HasSuffix(cfgOutput, "/") {
		cfgOutput = cfgOutput + "/"
	}
}

func config() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".rdm" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rdm")
	}

	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	resolveOutputDirectory()
}

// Wrapper to prompt for user input using AlecAivazis/survey
func askForData(fields []*survey.Question) map[string]interface{} {
	data := map[string]interface{}{}
	err := survey.Ask(fields, &data)

	if err != nil {
		panic(err.Error())
	}

	return data
}
