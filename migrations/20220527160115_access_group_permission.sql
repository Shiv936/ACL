-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_group_permission
(
    access_group_id        VARCHAR(64) NOT NULL REFRENCES access_group(id),
	resource_id            VARCHAR(64) NOT NULL REFRENCES resource(id),
	sub_resource_id        VARCHAR(64) NOT NULL REFRENCES sub_resource(id),
	sub_resource_action_id VARCHAR(64) NOT NULL REFRENCES sub_resource_action(id),
	created_at             TIMESTAMP NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC'),
	modified_at            TIMESTAMP NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC'),
    PRIMARY KEY (access_group_id ,resource_id ,sub_resource_id ,sub_resource_action_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE access_group DROP FOREIGN KEY;
DROP TABLE IF EXISTS access_group_permission;
-- +goose StatementEnd
