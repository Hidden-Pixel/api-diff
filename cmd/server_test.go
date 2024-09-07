package cmd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hidden-Pixel/api-diff/src/database"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	APIVersions []database.APIVersion
	InsertError error
	GetError    error
}

func (m *MockDB) GetAllAPIVersions() ([]database.APIVersion, error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	return m.APIVersions, nil
}

func (m *MockDB) InsertAPIVersion(version *database.APIVersion) error {
	if m.InsertError != nil {
		return m.InsertError
	}
	m.APIVersions = append(m.APIVersions, *version)
	return nil
}

func (m *MockDB) InsertAPIRequest(request *database.APIRequest) error {
	return nil
}

func (m *MockDB) GetAllAPIRequests() ([]database.APIRequest, error) {
	return nil, nil
}

func (m *MockDB) InsertAPIDiff(diff *database.APIDiff) error {
	return nil
}

func (m *MockDB) GetAllAPIDiffs() ([]database.APIDiff, error) {
	return nil, nil
}

func (m *MockDB) CreateAPIDiff(sourceRequestID int, targetRequestID int) (*database.APIDiff, error) {
	return nil, nil
}

func TestGETVersionHandler(t *testing.T) {
	// Initialize mock database and handler
	mockDB := &MockDB{
		APIVersions: []database.APIVersion{
			{VersionName: "v1.0.0"},
			{VersionName: "v1.1.0"},
		},
	}
	handler := &HTTPServer{DB: mockDB, Router: http.NewServeMux()}
	handler.AttachRoutes()

	// Create a request
	req, err := http.NewRequest("GET", "/v1/version", nil)
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.Router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var actualVersions []database.APIVersion
	err = json.NewDecoder(rr.Body).Decode(&actualVersions)
	assert.NoError(t, err)
	assert.Equal(t, mockDB.APIVersions, actualVersions)
}

func TestPOSTVersionHandler(t *testing.T) {
	// Initialize mock database and handler
	mockDB := &MockDB{}
	handler := &HTTPServer{DB: mockDB, Router: http.NewServeMux()}
	handler.AttachRoutes()

	// Prepare payload
	newVersion := database.APIVersion{VersionName: "v1.2.0"}

	// Create a request body
	body, err := json.Marshal(newVersion)
	assert.NoError(t, err)

	// Create a request
	req, err := http.NewRequest("POST", "/v1/version", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler
	handler.Router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var actualVersion database.APIVersion
	err = json.NewDecoder(rr.Body).Decode(&actualVersion)
	assert.NoError(t, err)
	assert.Equal(t, newVersion, actualVersion)

	// Verify that the version was added to the mock database
	assert.Contains(t, mockDB.APIVersions, newVersion)
}
