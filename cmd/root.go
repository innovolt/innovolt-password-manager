package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var rootCmd = &cobra.Command{
	Use:   "innovolt-pm",
	Short: "Innovolt Password Manager.",
	Long:  "Innovolt Password Manager which uses Fortanix SDKMS to store Secret.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	if err := initConfigFile(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func initConfigFile() error {
	home, err := homedir.Dir()
	if err != nil {
		return err
	}

	cfgDir := home + "/.innovolt-pm"
	cfgFile = cfgDir + "/authConfig.json"
	viper.SetConfigFile(cfgFile)

	// Create cfgFile if it does not exist
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		// First create cfgDir
		os.MkdirAll(cfgDir, 0700)
		if _, err = os.Create(cfgFile); err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}
