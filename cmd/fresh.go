package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var freshCmd = &cobra.Command{
	Use:   "fresh",
	Short: "Create a blank pair of config and credentials files",
	RunE: func(cmd *cobra.Command, args []string) error {
		configFile, err := os.Create(awsConfigFile)
		if err != nil {
			return err
		}
		defer configFile.Close()
		credentialsFile, err := os.Create(awsCredentialsFile)
		if err != nil {
			return err
		}
		defer credentialsFile.Close()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(freshCmd)
}
