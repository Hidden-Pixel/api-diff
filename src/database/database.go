package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type APIRequest struct {
	ID                 int       `json:"id"`
	Endpoint           string    `json:"endpoint"`
	Method             string    `json:"method"`
	SourceVersionID    string    `json:"source_version_id"`
	SourceRequestBody  []byte    `json:"source_request_body"`
	SourceResponseBody []byte    `json:"source_response_body"`
	TargetVersionID    string    `json:"target_version_id"`
	TargetRequestBody  []byte    `json:"target_request_body"`
	TargetResponseBody []byte    `json:"target_response_body"`
	CreatedAt          time.Time `json:"created_at"`
}

type DB struct {
	*pgxpool.Pool
}

type DiffDatabase interface {
	InsertAPIRequest(request *APIRequest) error
	GetAllAPIRequests() ([]APIRequest, error)
}

func NewDB() (*DB, error) {
	connString := PGConnectionString()
	dbpool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &DB{
		Pool: dbpool,
	}, nil
}

func (db *DB) InsertAPIRequest(request *APIRequest) error {
	query := `INSERT INTO api_request (
				endpoint, method,
				source_version_id, source_request_body, source_response_body,
				target_version_id, target_request_body, target_response_body
			  )
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := db.QueryRow(context.Background(),
		query,
		request.Endpoint,
		request.Method,
		request.SourceVersionID,
		request.SourceRequestBody,
		request.SourceResponseBody,
		request.TargetVersionID,
		request.TargetRequestBody,
		request.TargetResponseBody,
	).Scan(&request.ID)
	return err
}

func (db *DB) GetAllAPIRequests() ([]APIRequest, error) {
	rows, err := db.Query(context.Background(),
		`SELECT id, endpoint, method,
				source_version_id, source_request_body, source_response_body,
				target_version_id, target_request_body, target_response_body,
				created_at FROM api_request`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var requests []APIRequest
	for rows.Next() {
		var request APIRequest
		err := rows.Scan(
			&request.ID,
			&request.Endpoint,
			&request.Method,
			&request.SourceVersionID,
			&request.SourceRequestBody,
			&request.SourceResponseBody,
			&request.TargetVersionID,
			&request.TargetRequestBody,
			&request.TargetResponseBody,
			&request.CreatedAt)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func PGConnectionString() string {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		viper.GetString("POSTGRES_USER"),
		url.QueryEscape(viper.GetString("POSTGRES_PASSWORD")),
		viper.GetString("POSTGRES_HOST"),
		viper.GetUint("POSTGRES_PORT"),
		viper.GetString("POSTGRES_DB"))
	return connectionString
}
