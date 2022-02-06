package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [config name]",
	Short: "Delete a named config and credentials pair",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configName := args[0]
		err := os.Remove(awsConfigFile + "." + configName)
		if err != nil {
			return err
		}
		err = os.Remove(awsCredentialsFile + "." + configName)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
