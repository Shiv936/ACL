-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sub_resource_action (
    id VARCHAR(64) NOT NULL,
    sub_resource_id VARCHAR(64) NOT NULL REFERENCES sub_resource(id) ,
    name VARCHAR(128) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    modified_at TIMESTAMP NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY(id),
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sub_resource_action DROP FOREIGN KEY;
DROP TABLE IF EXISTS sub_resource_action;
-- +goose StatementEnd
