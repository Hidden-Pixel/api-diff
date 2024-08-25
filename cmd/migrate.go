package cmd

import (
	"fmt"
	"log"
	"net/url"

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
	dsn := PostgresConnectionString()
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

func PostgresConnectionString() string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("POSTGRES_USER"),
		url.QueryEscape(viper.GetString("POSTGRES_PASSWORD")),
		viper.GetString("POSTGRES_HOST"),
		viper.GetUint("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB"))
	return dsn
}
