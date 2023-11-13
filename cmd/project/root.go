package project

import "github.com/spf13/cobra"

func Root() *cobra.Command {

	rootCmd := &cobra.Command{
		Use:   "project",
		Short: "manage all projects",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
	}

	rootCmd.AddCommand(newCmd)

	return rootCmd
}
