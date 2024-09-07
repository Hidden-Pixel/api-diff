package integration

import (
	"testing"

	"github.com/Hidden-Pixel/api-diff/client"
	"github.com/Hidden-Pixel/api-diff/src/database"
	"github.com/stretchr/testify/assert"
)

func TestHTTPServerIntegrationWithRealDB(t *testing.T) {
	client := client.NewClient("http://localhost:8081")

	// Test POST /version
	newVersion := database.APIVersion{VersionName: "v1.2.0"}
	createdVersion, err := client.PostVersion(newVersion)
	assert.NoError(t, err)
	assert.Equal(t, newVersion, *createdVersion)

	// Test GET /version
	versions, err := client.GetVersions()
	assert.NoError(t, err)
	assert.Contains(t, versions, newVersion)
}
