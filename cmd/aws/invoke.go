package main

import (
	"get.porter.sh/mixin/aws/pkg/aws"
	"github.com/spf13/cobra"
)

func buildInvokeCommand(m *aws.Mixin) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "Execute the invoke functionality of this mixin",
		RunE: func(cmd *cobra.Command, args []string) error {
			return m.Execute()
		},
	}

	// Define a flag for --action so that its presence doesn't cause errors, but ignore it since exec doesn't need it
	var action string
	cmd.Flags().StringVar(&action, "action", "", "Custom action name to invoke.")

	return cmd
}
