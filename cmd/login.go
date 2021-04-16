package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"innovolt-pm/auth"
)

func init() {
	loginCmd := loginCmd()
	// rootCmd is defined in root.go
	rootCmd.AddCommand(&loginCmd)
}

func loginCmd() cobra.Command {
	var username string
	var password string
	var apikey string

	var cmdLogin = cobra.Command{
		Use:   "login user|app",
		Short: "Log into SDKMS (https://sdkms.fortanix.com using user credentials)",
		Long:  `Authenticate an User into SDKMS using its username and password.`,
		Args:  cobra.ExactArgs(1),
	}

	var cmdUser = &cobra.Command{
		Use:   "user",
		Short: "Log into SDKMS (https://sdkms.fortanix.com) using user credentials",
		Long:  `Authenticate an User into SDKMS using its username and password.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth.Authenticate(&auth.Credential{
				User: auth.UserCredential{
					Username: viper.GetString("username"),
					Password: viper.GetString("password"),
				},
				App: nil,
			})
		},
	}
	cmdUser.Flags().StringVarP(&username, "username", "u", "", "Enter SDKMS Username")
	cmdUser.Flags().StringVarP(&password, "password", "p", "", "Enter SDKMS Password")
	cmdUser.MarkFlagRequired("username")
	cmdUser.MarkFlagRequired("password")
	viper.BindPFlag("username", cmdUser.Flags().Lookup("username"))
	viper.BindPFlag("password", cmdUser.Flags().Lookup("password"))

	var cmdApp = &cobra.Command{
		Use:   "app",
		Short: "Log into SDKMS (https://sdkms.fortanix.com) using App API Key",
		Long:  `Authenticate an App into SDKMS using its API Key.`,
		Run: func(cmd *cobra.Command, args []string) {
			auth.Authenticate(&auth.Credential{
				User: nil,
				App: auth.AppCredential{
					ApiKey: viper.GetString("apiKey"),
				},
			})
		},
	}
	cmdApp.Flags().StringVarP(&apikey, "apikey", "a", "", "Enter SDKMS App API Key")
	cmdApp.MarkFlagRequired("apikey")
	viper.BindPFlag("apikey", cmdApp.Flags().Lookup("apikey"))

	cmdLogin.AddCommand(cmdUser)
	cmdLogin.AddCommand(cmdApp)

	return cmdLogin
}
