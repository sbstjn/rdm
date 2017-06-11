package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show command version",
		Run: func(cmd *cobra.Command, args []string) {
			if version == "" {
				fmt.Println("dev")
			} else {
				fmt.Println(version)
			}
		},
	})
}
