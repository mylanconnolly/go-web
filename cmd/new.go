package cmd

import (
	"fmt"
	"os"

	"github.com/mylanconnolly/go-web/lib/generators"
	"github.com/spf13/cobra"
)

var (
	pkgFlag string
)

// newCmd represents the new command when called without any subcommands
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generates a new project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := generators.New(pkgFlag); err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	newCmd.Flags().StringVarP(&pkgFlag, "package", "p", "", "The package name")
}
