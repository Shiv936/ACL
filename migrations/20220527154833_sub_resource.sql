-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sub_resource (
    id VARCHAR(64) NOT NULL,
    resource_id VARCHAR(64) NOT NULL REFERENCES resource(id) ,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sub_resource DROP FOREIGN KEY;
DROP TABLE IF EXISTS sub_resource;
-- +goose StatementEnd
