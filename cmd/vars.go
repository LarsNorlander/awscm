package cmd

import (
	"os"
	"path"
)

var (
	awsConfigLocation  string
	awsConfigFile      string
	awsCredentialsFile string
)

func init() {
	var err error
	awsConfigLocation, err = os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	awsConfigLocation = path.Join(awsConfigLocation, ".aws")

	awsConfigFile = path.Join(awsConfigLocation, "config")
	awsCredentialsFile = path.Join(awsConfigLocation, "credentials")
}
