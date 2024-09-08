package database

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/viper"
)

type APIVersion struct {
	ID          int       `json:"id"`
	VersionName string    `json:"version_name"`
	ReleaseDate time.Time `json:"release_date"`
	Description string    `json:"description"`
}

type APIRequest struct {
	ID           int       `json:"id"`
	VersionID    int       `json:"version_id"`
	Endpoint     string    `json:"endpoint"`
	Method       string    `json:"method"`
	RequestBody  []byte    `json:"request_body"`
	ResponseBody []byte    `json:"response_body"`
	Timestamp    time.Time `json:"timestamp"`
}

type APIDiff struct {
	ID              int       `json:"id"`
	SourceRequestID int       `json:"source_request_id"`
	TargetRequestID int       `json:"target_request_id"`
	DiffMetric      string    `json:"diff_metric"`
	DivergenceScore float64   `json:"divergence_score"`
	CreatedAt       time.Time `json:"created_at"`
}

type DB struct {
	*pgxpool.Pool
}

type DiffDatabase interface {
	InsertAPIVersion(version *APIVersion) error
	GetAllAPIVersions() ([]APIVersion, error)
	InsertAPIRequest(request *APIRequest) error
	GetAllAPIRequests() ([]APIRequest, error)
	InsertAPIDiff(diff *APIDiff) error
	GetAllAPIDiffs() ([]APIDiff, error)
	CreateAPIDiff(sourceRequestID int, targetRequestID int) (*APIDiff, error)
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

func (db *DB) InsertAPIVersion(version *APIVersion) error {
	query := `INSERT INTO api_version (version_name, release_date, description)
	          VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(context.Background(),
		query,
		version.VersionName,
		version.ReleaseDate,
		version.Description).Scan(&version.ID)
	return err
}

func (db *DB) GetAllAPIVersions() ([]APIVersion, error) {
	rows, err := db.Query(context.Background(), "SELECT id, version_name, release_date, description FROM api_version")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var versions []APIVersion
	for rows.Next() {
		var version APIVersion
		err := rows.Scan(&version.ID,
			&version.VersionName,
			&version.ReleaseDate,
			&version.Description)
		if err != nil {
			return nil, err
		}
		versions = append(versions, version)
	}
	return versions, nil
}

func (db *DB) InsertAPIRequest(request *APIRequest) error {
	query := `INSERT INTO api_request (version_id, endpoint, method, request_body, response_body)
	          VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.QueryRow(context.Background(),
		query,
		request.VersionID,
		request.Endpoint,
		request.Method,
		request.RequestBody,
		request.ResponseBody).Scan(&request.ID)
	return err
}

func (db *DB) GetAllAPIRequests() ([]APIRequest, error) {
	rows, err := db.Query(context.Background(),
		`SELECT id, version_id, endpoint, method, request_body, response_body, timestamp FROM api_request`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var requests []APIRequest
	for rows.Next() {
		var request APIRequest
		err := rows.Scan(&request.ID,
			&request.VersionID,
			&request.Endpoint,
			&request.Method,
			&request.RequestBody,
			&request.ResponseBody,
			&request.Timestamp)
		if err != nil {
			return nil, err
		}
		requests = append(requests, request)
	}
	return requests, nil
}

func (db *DB) InsertAPIDiff(diff *APIDiff) error {
	query := `INSERT INTO api_diff (source_request_id, target_requet_id, diff_metric, divergence_score)
	          VALUES ($1, $2, $3) RETURNING id`
	err := db.QueryRow(context.Background(),
		query,
		diff.SourceRequestID,
		diff.TargetRequestID,
		diff.DiffMetric,
		diff.DivergenceScore).Scan(&diff.ID)
	return err
}

func (db *DB) GetAllAPIDiffs() ([]APIDiff, error) {
	rows, err := db.Query(context.Background(),
		`SELECT id, source_request_id, target_request_id, diff_metric, divergence_score, created_at FROM api_diff`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var diffs []APIDiff
	for rows.Next() {
		var diff APIDiff
		err := rows.Scan(&diff.ID,
			&diff.SourceRequestID,
			&diff.TargetRequestID,
			&diff.DiffMetric,
			&diff.DivergenceScore,
			&diff.CreatedAt)
		if err != nil {
			return nil, err
		}
		diffs = append(diffs, diff)
	}
	return diffs, nil
}

func (db *DB) GetAPIDiffByID(id int) (*APIDiff, error) {
	row := db.QueryRow(context.Background(),
		`SELECT id, request_source_id, request_target_id, diff_metric, divergence_score, created_at FROM api_diff WHERE id = $1`, id)
	diff := APIDiff{}
	err := row.Scan(&diff.ID,
		&diff.SourceRequestID,
		&diff.TargetRequestID,
		&diff.DiffMetric,
		&diff.DivergenceScore,
		&diff.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &diff, nil
}

func (db *DB) CreateAPIDiff(sourceRequestID int, targetRequestID int) (*APIDiff, error) {
	var id int
	err := db.QueryRow(context.Background(),
		diffQuery,
		sourceRequestID,
		targetRequestID).Scan(&id)
	if err != nil {
		fmt.Printf("failed to insert diff record, error: %+v\n", err)
		return nil, err
	}
	fmt.Printf("successfully inserted diff record, id: %d\n", id)
	newDiff, err := db.GetAPIDiffByID(id)
	if err != nil {
		fmt.Printf("failed to fetch diff record, error: %+v\n", err)
		return nil, err
	}
	return newDiff, nil
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
