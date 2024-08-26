package cmd

import (
	"log"

	db "github.com/Hidden-Pixel/api-diff/src/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "",
	Long:  ``,
	Run:   RunMigrate,
}

func RunMigrate(cmd *cobra.Command, args []string) {
	dsn := db.PostgresConnectionString()
	migrationPath := viper.GetString("MIGRATION_PATH")
	logger := log.Logger{}
	m, err := migrate.New(
		migrationPath,
		dsn)
	logger.Printf("Running Database Migrations ...")
	if err != nil {
		logger.Fatalf("can't create migration helper: %s", err.Error())
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.Fatalf("can't migrate the database: %s", err.Error())
	}
	logger.Print("Finished Database Migrations ...")
}
