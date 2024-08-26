package cmd

import (
	"log"

	db "github.com/Hidden-Pixel/api-diff/src/database"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
	connectionString := db.PGConnectionString()
	migrationPath := viper.GetString("MIGRATION_PATH")
	log.Printf("connection string: %s", connectionString)
	log.Printf("migration path: %s", migrationPath)
	m, err := migrate.New(
		migrationPath,
		connectionString)
	log.Printf("Running Database Migrations ...")
	if err != nil {
		log.Fatalf("can't create migration helper: %s", err.Error())
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("can't migrate the database: %s", err.Error())
	}
	log.Print("Finished Database Migrations ...")
}
