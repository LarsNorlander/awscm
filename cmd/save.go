package cmd

import (
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [config name]",
	Short: "Save your current config and credentials files into a named pair",
	RunE: func(cmd *cobra.Command, args []string) error {
		configName := args[0]
		err := copy(awsConfigFile, awsConfigFile+"."+configName)
		if err != nil {
			return err
		}
		err = copy(awsCredentialsFile, awsCredentialsFile+"."+configName)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}
