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

	v1 := database.APIVersion{VersionName: "v1.0.0"}
	createdV1, err := client.PostVersion(v1)
	assert.NoError(t, err)

	versions, err := client.GetVersions()
	assert.NoError(t, err)
	assert.Contains(t, versions, createdV1)

	v2 := database.APIVersion{VersionName: "v2.0.0"}
	createdV2, err := client.PostVersion(v2)
	assert.NoError(t, err)

	versions, err = client.GetVersions()
	assert.NoError(t, err)
	assert.Contains(t, versions, createdV2)

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

	v1Request := database.APIRequest{
		VersionID:    createdV1.ID,
		Endpoint:     "/test",
		Method:       "GET",
		RequestBody:  nil,
		ResponseBody: responseBody,
	}
	createdV1Request, err := client.PostRequest(v1Request)
	assert.NoError(t, err)

	v2Request := database.APIRequest{
		VersionID:    createdV2.ID,
		Endpoint:     "/test",
		Method:       "GET",
		RequestBody:  nil,
		ResponseBody: responseBody,
	}
	createdV2Request, err := client.PostRequest(v2Request)
	assert.NoError(t, err)

	diffReq := database.APIDiff{
		SourceRequestID: createdV2Request.ID,
		TargetRequestID: createdV1Request.ID,
	}
	diff, err := client.PostDiff(&diffReq)
	assert.NoError(t, err)
	assert.NotNil(t, diff)
}
