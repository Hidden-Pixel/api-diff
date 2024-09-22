CREATE TABLE api_request (
    id SERIAL PRIMARY KEY,
    endpoint VARCHAR(255) NOT NULL,
    method VARCHAR(10) NOT NULL,
    -- source data
    source_version_id VARCHAR(10) NOT NULL,
    source_request_body JSONB,
    source_response_body JSONB,
    -- target data
    target_version_id VARCHAR(10) NOT NULL,
    target_request_body JSONB,
    target_response_body JSONB,
    -- creation 
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
