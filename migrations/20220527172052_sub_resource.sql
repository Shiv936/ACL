-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sub_resource (
    id VARCHAR(64) NOT NULL,
    resource_id VARCHAR(64) NOT NULL,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(id),
    CONSTRAINT idx_resource_id FOREIGN KEY(resource_id) REFERENCES resource(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sub_resource DROP CONSTRAINT idx_resource_id;
DROP TABLE IF EXISTS sub_resource;
-- +goose StatementEnd
