package cmd

import (
	"github.com/spf13/cobra"
	"payment-gw/pkg/conf"
	"payment-gw/pkg/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := conf.Load(); err != nil {
			return err
		}
		return database.RunMigrations()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}