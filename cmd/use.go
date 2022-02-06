package cmd

import (
	"github.com/spf13/cobra"
)

var useCmd = &cobra.Command{
	Use:   "use [config name]",
	Short: "Replace your current config and credentials files with a named pair",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		configName := args[0]
		err := copy(awsConfigFile+"."+configName, awsConfigFile)
		if err != nil {
			return err
		}
		err = copy(awsCredentialsFile+"."+configName, awsCredentialsFile)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(useCmd)
}
