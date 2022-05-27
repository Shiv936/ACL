-- +goose Up
-- +goose StatementBegin
CREATE TABLE resource (
    id VARCHAR(64) NOT NULL,
    name VARCHAR(64) NOT NULL,
    description text,
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS resource;
-- +goose StatementEnd
