package project

import (
	"fmt"

	"github.com/spf13/cobra"
)

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new hapi project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from Start")
	},
}
