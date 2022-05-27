-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS access_group_permission
(
    access_group_id        VARCHAR(64) NOT NULL ,
	resource_id            VARCHAR(64) NOT NULL ,
	sub_resource_id        VARCHAR(64) NOT NULL ,
	sub_resource_action_id VARCHAR(64) NOT NULL ,
	created_at             TIMESTAMP NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC'),
	modified_at            TIMESTAMP NOT NULL DEFAULT(current_timestamp AT TIME ZONE 'UTC'),
    CONSTRAINT fx_access_group_id FOREIGN KEY(access_group_id) REFERENCES access_group(id),
    CONSTRAINT fx_resource_id FOREIGN KEY(resource_id) REFERENCES resource(id),
    CONSTRAINT fx_sub_resource_id FOREIGN KEY(sub_resource_id) REFERENCES sub_resource(id),
    CONSTRAINT fx_sub_resource_action_id FOREIGN KEY(sub_resource_action_id) REFERENCES sub_resource_action(id)
);
CREATE UNIQUE INDEX idx_access_group_permission
ON access_group_permission(access_group_id ,resource_id ,sub_resource_id ,sub_resource_action_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE access_group_permission DROP CONSTRAINT fx_access_group_id;
ALTER TABLE access_group_permission DROP CONSTRAINT fx_resource_id;
ALTER TABLE access_group_permission DROP CONSTRAINT fx_sub_resource_id;
ALTER TABLE access_group_permission DROP CONSTRAINT fx_sub_resource_action_id;
DROP INDEX idx_access_group_permission;
DROP TABLE IF EXISTS access_group_permission;
-- +goose StatementEnd
