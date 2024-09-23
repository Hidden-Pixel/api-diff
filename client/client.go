 package client
//
// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
//
// 	"github.com/Hidden-Pixel/api-diff/src/database"
// )
//
// type Client struct {
// 	BaseURL    string
// 	HTTPClient *http.Client
// }
//
// func NewClient(baseURL string) *Client {
// 	return &Client{
// 		BaseURL:    baseURL,
// 		HTTPClient: &http.Client{},
// 	}
// }
//
// func (c *Client) GetVersions() ([]*database.APIVersion, error) {
// 	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/v1/version", c.BaseURL))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("failed to get versions: %s", string(body))
// 	}
//
// 	var versions []*database.APIVersion
// 	if err := json.NewDecoder(resp.Body).Decode(&versions); err != nil {
// 		return nil, err
// 	}
// 	return versions, nil
// }
//
// func (c *Client) PostVersion(version database.APIVersion) (*database.APIVersion, error) {
// 	body, err := json.Marshal(version)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	resp, err := c.HTTPClient.Post(fmt.Sprintf("%s/v1/version", c.BaseURL), "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("failed to post version: %s", string(body))
// 	}
//
// 	var createdVersion database.APIVersion
// 	if err := json.NewDecoder(resp.Body).Decode(&createdVersion); err != nil {
// 		return nil, err
// 	}
// 	return &createdVersion, nil
// }
//
// func (c *Client) GetRequests() ([]database.APIRequest, error) {
// 	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/v1/request", c.BaseURL))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("failed to get requests: %s", string(body))
// 	}
//
// 	var requests []database.APIRequest
// 	if err := json.NewDecoder(resp.Body).Decode(&requests); err != nil {
// 		return nil, err
// 	}
// 	return requests, nil
// }
//
// func (c *Client) PostRequest(request database.APIRequest) (*database.APIRequest, error) {
// 	body, err := json.Marshal(request)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	resp, err := c.HTTPClient.Post(fmt.Sprintf("%s/v1/request", c.BaseURL), "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("failed to post request: %s", string(body))
// 	}
//
// 	var createdRequest database.APIRequest
// 	if err := json.NewDecoder(resp.Body).Decode(&createdRequest); err != nil {
// 		return nil, err
// 	}
// 	return &createdRequest, nil
// }
//
// func (c *Client) GetDiffs() ([]database.APIDiff, error) {
// 	resp, err := c.HTTPClient.Get(fmt.Sprintf("%s/v1/diff", c.BaseURL))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("failed to get diffs: %s", string(body))
// 	}
//
// 	var diffs []database.APIDiff
// 	if err := json.NewDecoder(resp.Body).Decode(&diffs); err != nil {
// 		return nil, err
// 	}
// 	return diffs, nil
// }
//
// func (c *Client) PostDiff(diff *database.APIDiff) (*database.APIDiff, error) {
// 	body, err := json.Marshal(diff)
// 	if err != nil {
// 		return nil, err
// 	}
// 	resp, err := c.HTTPClient.Post(fmt.Sprintf("%s/v1/diff", c.BaseURL), "application/json", bytes.NewBuffer(body))
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := io.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("failed to post diff: %s", string(body))
// 	}
// 	var createdDiff database.APIDiff
// 	if err := json.NewDecoder(resp.Body).Decode(&createdDiff); err != nil {
// 		return nil, err
// 	}
// 	return &createdDiff, nil
// }
