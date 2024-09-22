package integration

//
// import (
// 	"encoding/json"
// 	"testing"
//
// 	"github.com/Hidden-Pixel/api-diff/client"
// 	"github.com/Hidden-Pixel/api-diff/src/database"
// 	"github.com/stretchr/testify/assert"
// )
//
// func TestHTTPServerIntegrationWithRealDB(t *testing.T) {
// 	client := client.NewClient("http://localhost:8081")
//
// 	// Create version 1
// 	v1 := database.APIVersion{VersionName: "v1.0.0"}
// 	createdV1, err := client.PostVersion(v1)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	versions, err := client.GetVersions()
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, versions, createdV1) {
// 		t.Fatal()
// 	}
//
// 	// Create version 2
// 	v2 := database.APIVersion{VersionName: "v2.0.0"}
// 	createdV2, err := client.PostVersion(v2)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	versions, err = client.GetVersions()
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, versions, createdV2) {
// 		t.Fatal()
// 	}
//
// 	// Create test response
// 	type TestResponse struct {
// 		ID   int    `json:"id"`
// 		Name string `json:"name"`
// 	}
// 	tr := TestResponse{
// 		ID:   1,
// 		Name: "test-name",
// 	}
// 	responseBody, err := json.Marshal(tr)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	// Create requests with different request bodies
// 	v1Request := database.APIRequest{
// 		VersionID:    createdV1.ID,
// 		Endpoint:     "/test",
// 		Method:       "GET",
// 		RequestBody:  json.RawMessage(`{"field1": "value1", "field2": "value2"}`),
// 		ResponseBody: responseBody,
// 	}
// 	createdV1Request, err := client.PostRequest(v1Request)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	v2Request := database.APIRequest{
// 		VersionID:    createdV2.ID,
// 		Endpoint:     "/test",
// 		Method:       "GET",
// 		RequestBody:  json.RawMessage(`{"field1": "value1", "field3": "value3"}`), // Note the difference here
// 		ResponseBody: responseBody,
// 	}
// 	createdV2Request, err := client.PostRequest(v2Request)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	// Create another version
// 	v3 := database.APIVersion{VersionName: "v3.0.0"}
// 	createdV3, err := client.PostVersion(v3)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	// Create another request with a new difference
// 	anotherRequest := database.APIRequest{
// 		VersionID:    createdV3.ID,
// 		Endpoint:     "/test",
// 		Method:       "GET",
// 		RequestBody:  json.RawMessage(`{"field1": "value1", "field4": "value4"}`), // New field introduced
// 		ResponseBody: responseBody,
// 	}
// 	createdAnotherRequest, err := client.PostRequest(anotherRequest)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
//
// 	// Create a diff request between the first and the third request
// 	diffReq1 := database.APIDiff{
// 		SourceRequestID: createdAnotherRequest.ID,
// 		TargetRequestID: createdV1Request.ID,
// 	}
// 	diff1, err := client.PostDiff(&diffReq1)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
// 	if !assert.NotNil(t, diff1) {
// 		t.Fatal()
// 	}
//
// 	// Create a diff request between the second and the third request
// 	diffReq2 := database.APIDiff{
// 		SourceRequestID: createdAnotherRequest.ID,
// 		TargetRequestID: createdV2Request.ID,
// 	}
// 	diff2, err := client.PostDiff(&diffReq2)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
// 	if !assert.NotNil(t, diff2) {
// 		t.Fatal()
// 	}
//
// 	// Create a diff request between the second and the third request
// 	diffReq3 := database.APIDiff{
// 		SourceRequestID: createdV2Request.ID,
// 		TargetRequestID: createdV2Request.ID,
// 	}
// 	diff3, err := client.PostDiff(&diffReq3)
// 	if !assert.NoError(t, err) {
// 		t.Fatal()
// 	}
// 	if !assert.NotNil(t, diff2) {
// 		t.Fatal()
// 	}
//
// 	// Additional assertions to verify the diff content
// 	// For diff1, we expect differences between the first request and the additional request
// 	// TODO(nick): add checks for each diff field
// 	if !assert.Contains(t, diff1.DiffMetric, "added") {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, diff1.DiffMetric, "removed") {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, diff1.DiffMetric, "changed") {
// 		t.Fatal()
// 	}
//
// 	// For diff2, we expect differences between the second request and the additional request
// 	// TODO(nick): add checks for each diff field
// 	if !assert.Contains(t, diff2.DiffMetric, "added") {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, diff2.DiffMetric, "removed") {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, diff2.DiffMetric, "changed") {
// 		t.Fatal()
// 	}
//
// 	// For diff3, expect not diffs should all be null
// 	if !assert.Contains(t, diff3.DiffMetric, "added") {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, diff3.DiffMetric, "removed") {
// 		t.Fatal()
// 	}
// 	if !assert.Contains(t, diff3.DiffMetric, "changed") {
// 		t.Fatal()
// 	}
// }
