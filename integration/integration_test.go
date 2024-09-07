package integration

import (
	"encoding/json"
	"testing"

	"github.com/Hidden-Pixel/api-diff/client"
	"github.com/Hidden-Pixel/api-diff/src/database"
	"github.com/stretchr/testify/assert"
)

func TestHTTPServerIntegrationWithRealDB(t *testing.T) {
	client := client.NewClient("http://localhost:8081")

	// Create version 1
	v1 := database.APIVersion{VersionName: "v1.0.0"}
	createdV1, err := client.PostVersion(v1)
	assert.NoError(t, err)

	versions, err := client.GetVersions()
	assert.NoError(t, err)
	assert.Contains(t, versions, createdV1)

	// Create version 2
	v2 := database.APIVersion{VersionName: "v2.0.0"}
	createdV2, err := client.PostVersion(v2)
	assert.NoError(t, err)

	versions, err = client.GetVersions()
	assert.NoError(t, err)
	assert.Contains(t, versions, createdV2)

	// Create test response
	type TestResponse struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	tr := TestResponse{
		ID:   1,
		Name: "test-name",
	}
	responseBody, err := json.Marshal(tr)
	assert.NoError(t, err)

	// Create requests with different request bodies
	v1Request := database.APIRequest{
		VersionID:    createdV1.ID,
		Endpoint:     "/test",
		Method:       "GET",
		RequestBody:  json.RawMessage(`{"field1": "value1", "field2": "value2"}`),
		ResponseBody: responseBody,
	}
	createdV1Request, err := client.PostRequest(v1Request)
	assert.NoError(t, err)

	v2Request := database.APIRequest{
		VersionID:    createdV2.ID,
		Endpoint:     "/test",
		Method:       "GET",
		RequestBody:  json.RawMessage(`{"field1": "value1", "field3": "value3"}`), // Note the difference here
		ResponseBody: responseBody,
	}
	createdV2Request, err := client.PostRequest(v2Request)
	assert.NoError(t, err)

	// Create another version
	v3 := database.APIVersion{VersionName: "v3.0.0"}
	createdV3, err := client.PostVersion(v3)
	assert.NoError(t, err)

	// Create another request with a new difference
	anotherRequest := database.APIRequest{
		VersionID:    createdV3.ID,
		Endpoint:     "/test",
		Method:       "GET",
		RequestBody:  json.RawMessage(`{"field1": "value1", "field4": "value4"}`), // New field introduced
		ResponseBody: responseBody,
	}
	createdAnotherRequest, err := client.PostRequest(anotherRequest)
	assert.NoError(t, err)

	// Create a diff request between the first and the third request
	diffReq1 := database.APIDiff{
		SourceRequestID: createdAnotherRequest.ID,
		TargetRequestID: createdV1Request.ID,
	}
	diff1, err := client.PostDiff(&diffReq1)
	assert.NoError(t, err)
	assert.NotNil(t, diff1)

	// Create a diff request between the second and the third request
	diffReq2 := database.APIDiff{
		SourceRequestID: createdAnotherRequest.ID,
		TargetRequestID: createdV2Request.ID,
	}
	diff2, err := client.PostDiff(&diffReq2)
	assert.NoError(t, err)
	assert.NotNil(t, diff2)

	// Additional assertions to verify the diff content
	// For diff1, we expect differences between the first request and the additional request
	assert.Contains(t, diff1.DiffMetric, "added")
	assert.Contains(t, diff1.DiffMetric, "removed")
	assert.Contains(t, diff1.DiffMetric, "changed")

	// For diff2, we expect differences between the second request and the additional request
	assert.Contains(t, diff2.DiffMetric, "added")
	assert.Contains(t, diff2.DiffMetric, "removed")
	assert.Contains(t, diff2.DiffMetric, "changed")
}
