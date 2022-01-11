package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "capelo",
	Short: "capelo.io site management tools",
}

func Execute() {

	// adds all available commands
	rootCommand.AddCommand(IndexCommand())

	// raise if there are any errors
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
