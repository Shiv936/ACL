package ports

import "acl/internal/core/domain/repositories"

type ResourceRepository interface {
	Insert(
		name *string,
		description string,
	) (string, error)
	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]repositories.Resource, bool, error)
	ReadOne(id *string) (repositories.Resource, bool, error)
	Update(
		id *string,
		name *string,
		description string,
	) (string, error)
	Delete(id string) error
}

type SubResourceRepository interface {
	Insert(
		name *string,
		resourceId *string,
		description string,
	) (string, error)
	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]repositories.SubResource, bool, error)
	ReadOne(id *string) (repositories.SubResource, bool, error)
	Update(
		id *string,
		name *string,
		resourceId *string,
		description string,
	) (string, error)
	Delete(id *string) error
}

type SubResourceActionRepository interface {
	Insert(
		name *string,
		subResourceId *string,
		description string,
	) (string, error)
	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]repositories.SubResourceAction, bool, error)
	ReadOne(id *string) (repositories.SubResourceAction, bool, error)
	Update(
		id *string,
		name *string,
		subResourceId *string,
		description string,
	) (string, error)
	Delete(id *string) error
}

type AccessGroup interface {
	Insert(
		name *string,
		description string,
	) (string, error)
	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]repositories.AccessGroup, bool, error)
	ReadOne(id *string) (repositories.AccessGroup, bool, error)
	Update(
		id *string,
		name *string,
		description string,
	) (string, error)
	Delete(id *string) error
}

type AccessGroupPermission interface {
	Insert(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (bool, error)
	ReadMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]repositories.AccessGroupPermission, bool, error)
	ReadByAccessGroup(accessGroupId *string) ([]repositories.AccessGroupPermission, bool, error)
	ReadByResource(resourceId *string) ([]repositories.AccessGroupPermission, bool, error)
	ReadBySubResource(subResourceId *string) ([]repositories.AccessGroupPermission, bool, error)
	ReadBySubResourceAction(subResourceActionId *string) ([]repositories.AccessGroupPermission, bool, error)
	UpdateByAccessGroup(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (string, error)
	UpdateBySubResource(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (string, error)
	UpdateByResource(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (string, error)
	Delete(repositories.AccessGroupPermission) error
}
