package mock

import "github.com/Hidden-Pixel/api-diff/src/database"

type MockDB struct {
	APIVersions []database.APIVersion
	APIRequests []database.APIRequest
	APIDiffs    []database.APIDiff
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

func (m *MockDB) GetAllAPIRequests() ([]database.APIRequest, error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	return m.APIRequests, nil
}

func (m *MockDB) InsertAPIRequest(request *database.APIRequest) error {
	if m.InsertError != nil {
		return m.InsertError
	}
	m.APIRequests = append(m.APIRequests, *request)
	return nil
}

func (m *MockDB) GetAllAPIDiffs() ([]database.APIDiff, error) {
	if m.GetError != nil {
		return nil, m.GetError
	}
	return m.APIDiffs, nil
}

func (m *MockDB) InsertAPIDiff(diff *database.APIDiff) error {
	if m.InsertError != nil {
		return m.InsertError
	}
	m.APIDiffs = append(m.APIDiffs, *diff)
	return nil
}

func (m *MockDB) CreateAPIDiff(sourceRequestID int, targetRequestID int) (*database.APIDiff, error) {
	if m.InsertError != nil {
		return nil, m.InsertError
	}

	// Create a new APIDiff based on the provided IDs (mocked behavior)
	diff := &database.APIDiff{
		// Add fields accordingly (ID, SourceRequestID, TargetRequestID, etc.)
		SourceRequestID: sourceRequestID,
		TargetRequestID: targetRequestID,
	}

	m.APIDiffs = append(m.APIDiffs, *diff)
	return diff, nil
}
