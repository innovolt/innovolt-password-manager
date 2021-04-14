package cmd

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"innovolt-pm/secret"
)

func init() {
	secretCmd := secretCmd()
	// rootCmd is defined in root.go
	rootCmd.AddCommand(&secretCmd)
}

func secretCmd() cobra.Command {
	var secretCmd = cobra.Command{
		Use:   "secret create|delete|get",
		Short: "Secret management",
		Long:  "Secret management",
		Args:  cobra.ExactArgs(1),
	}

	var createSecretCmd = &cobra.Command{
		Use:   "create secretName",
		Short: "Secret creation",
		Long:  "Secret creation",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := secret.CreateSecret(args[0])
			if err != nil {
				color.Red(err.Error())
			}
		},
	}

	var getSecretCmd = &cobra.Command{
		Use:   "get secretName",
		Short: "Fetch secret(s).",
		Long:  "Fetch a particular secret if name is provided otherwise all.",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := secret.GetSecret(args[0])
			if err != nil {
				color.Red(err.Error())
			}
		},
	}

	secretCmd.AddCommand(createSecretCmd)
	secretCmd.AddCommand(getSecretCmd)

	return secretCmd
}
