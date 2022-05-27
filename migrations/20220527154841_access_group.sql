-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_group
(
    id VARCHAR(64) NOT NULL PRIMARY KEY,
    name VARCHAR(64) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC'), 
    modified_at TIMESTAMP NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS access_group;
-- +goose StatementEnd
