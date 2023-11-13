package root

import (
	"fmt"
	"os"

	"github.com/lewis-catley/hapi/cmd/project"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hapi",
	Short: "hapi is the perfect api dev accompany tool",
	Long: `hapi becuase it makes api developers happy.
		Call APIs
		Import API specifications, save responses.
		Maybe one day write integration tests.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	rootCmd.AddCommand(project.NewCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
