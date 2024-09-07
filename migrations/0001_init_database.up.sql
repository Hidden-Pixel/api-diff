CREATE TABLE api_version (
    id SERIAL PRIMARY KEY,
    version_name VARCHAR(50) NOT NULL,
    release_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description TEXT
);

CREATE TABLE api_request (
    id SERIAL PRIMARY KEY,
    version_id INT REFERENCES api_version(id),
    endpoint VARCHAR(255) NOT NULL,
    method VARCHAR(10) NOT NULL, -- GET, POST, etc.
    request_body JSONB,
    response_body JSONB,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE api_diff (
    id SERIAL PRIMARY KEY,
    request_source_id INT REFERENCES api_request(id),
    target_source_id INT REFERENCES api_request(id),
    diff_metric JSONB, -- Storing the diff results as JSON
    divergence_score NUMERIC, -- Score representing the divergence level
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
