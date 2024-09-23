package client

//
// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
//
// 	"github.com/Hidden-Pixel/api-diff/cmd"
// 	"github.com/Hidden-Pixel/api-diff/src/database"
// 	"github.com/Hidden-Pixel/api-diff/src/database/mock"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestHTTPClientIntegration(t *testing.T) {
// 	// Setup the mock database and server
// 	mockDB := &mock.MockDB{}
// 	server := &cmd.HTTPServer{DB: mockDB, Router: http.NewServeMux()}
// 	server.AttachRoutes()
//
// 	// Create an in-memory test server
// 	testServer := httptest.NewServer(server.Router)
// 	defer testServer.Close()
//
// 	// Create an HTTP client
// 	client := NewClient(testServer.URL)
//
// 	// Test POST /version
// 	newVersion := database.APIVersion{VersionName: "v1.2.0"}
// 	createdVersion, err := client.PostVersion(newVersion)
// 	assert.NoError(t, err)
// 	assert.Equal(t, newVersion, *createdVersion)
//
// 	// Test GET /version
// 	versions, err := client.GetVersions()
// 	assert.NoError(t, err)
// 	assert.Contains(t, versions, newVersion)
//
// 	// Test POST /request
// 	newRequest := database.APIRequest{
// 		// Fill out with test request fields
// 	}
// 	createdRequest, err := client.PostRequest(newRequest)
// 	assert.NoError(t, err)
// 	assert.Equal(t, newRequest, *createdRequest)
//
// 	// Test GET /request
// 	requests, err := client.GetRequests()
// 	assert.NoError(t, err)
// 	assert.Contains(t, requests, newRequest)
//
// 	// Test POST /diff
// 	diffRequest := database.APIDiff{
// 		SourceRequestID: 1,
// 		TargetRequestID: 2,
// 	}
// 	createdDiff, err := client.PostDiff(&diffRequest)
// 	assert.NoError(t, err)
// 	assert.NotNil(t, createdDiff)
//
// 	// Test GET /diff
// 	diffs, err := client.GetDiffs()
// 	assert.NoError(t, err)
// 	assert.Contains(t, diffs, *createdDiff)
// }
