package database

import (
	"fmt"
	"net/url"

	"github.com/spf13/viper"
)

func PGConnectionString() string {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("POSTGRES_USER"),
		url.QueryEscape(viper.GetString("POSTGRES_PASSWORD")),
		viper.GetString("POSTGRES_HOST"),
		viper.GetUint("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB"))
	return connectionString
}
