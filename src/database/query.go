package database

const diffQuery = `
WITH source AS (
    SELECT id, request_body AS source_request_body, response_body AS source_response_body
    FROM api_request
    WHERE id = $1
),
target AS (
    SELECT id, request_body AS target_request_body, response_body AS target_response_body
    FROM api_request
    WHERE id = $2
),
diffs AS (
    SELECT
        jsonb_build_object(
            'added', (
                SELECT jsonb_object_agg(t.key, t.value)
                FROM jsonb_each(target.target_request_body) AS t(key, value)
                WHERE NOT jsonb_exists(source.source_request_body, t.key)
            ),
            'removed', (
                SELECT jsonb_object_agg(s.key, s.value)
                FROM jsonb_each(source.source_request_body) AS s(key, value)
                WHERE NOT jsonb_exists(target.target_request_body, s.key)
            ),
            'changed', (
                SELECT jsonb_object_agg(key, jsonb_build_object('from', value1, 'to', value2))
                FROM (
                    SELECT
                        COALESCE(old.key, new.key) AS key,
                        old.value AS value1,
                        new.value AS value2
                    FROM (
                        SELECT key, value
                        FROM jsonb_each(source.source_request_body)
                    ) AS old
                    FULL JOIN (
                        SELECT key, value
                        FROM jsonb_each(target.target_request_body)
                    ) AS new
                    ON old.key = new.key
                    WHERE old.value IS DISTINCT FROM new.value
                ) AS diff
            )
        ) AS diff_metric
    FROM source, target
)
INSERT INTO api_diff (request_source_id, target_source_id, diff_metric, divergence_score)
SELECT
    source.id AS request_source_id,
    target.id AS target_source_id,
    diffs.diff_metric,
    -- Example divergence score based on the number of changes
	-- TODO(nick): add other field and accumulate?
	-- NOTE(nick): default behavior, maybe change to custom layer in go instead?
    COALESCE(
        jsonb_array_length(jsonb_path_query_array(
            diffs.diff_metric->'changed', '$.*'
        )) +
        jsonb_array_length(jsonb_path_query_array(
            diffs.diff_metric->'added', '$.*'
        )) +
        jsonb_array_length(jsonb_path_query_array(
            diffs.diff_metric->'removed', '$.*'
        )), 
        0
    ) AS divergence_score
FROM source, target, diffs
`
