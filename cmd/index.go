package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/luiscape/luiscape/pkg/db"
	"github.com/spf13/cobra"
)

func IndexCommand() *cobra.Command {
	baseCmd := &cobra.Command{
		Use:   "index",
		Short: "generates database index from local path",
		Run: func(cmd *cobra.Command, args []string) {
			path, _ := cmd.Flags().GetString("path")

			// read database from path
			entries, err := db.Walk(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
				os.Exit(1)
			}

			// write json of database as serialized entries
			databasePath := filepath.Join(path, db.DEFAULT_DATABASE_FILENAME)
			db.Write(databasePath, entries)
		},
	}

	baseCmd.Flags().String("path", "./db", "path containing blog post articles in markdown")
	return baseCmd
}
