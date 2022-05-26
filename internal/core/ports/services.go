package ports

import "acl/internal/core/domain/services"

type ResourceService interface {
	Create(
		name *string,
		description string,
	) (services.Resource, error)
	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]services.Resource, bool, error)
	Get(id *string) (services.Resource, error)
	Modify(
		id *string,
		name *string,
		description string,
	) (services.Resource, error)
	Remove(id *string) error
}

type SubResourceService interface {
	Create(
		name *string,
		resourceId *string,
		description string,
	) (services.Resource, error)
	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]services.Resource, bool, error)
	Get(id *string) (services.Resource, error)
	Modify(
		id *string,
		name *string,
		description string,
	) (services.Resource, error)
	Remove(id *string) error
}

type SubResourceActionService interface {
	Create(
		name *string,
		subResourceId *string,
		description string,
	) (services.SubResource, error)
	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]services.SubResource, bool, error)
	Get(id *string) (services.SubResource, error)
	Modify(
		id *string,
		name *string,
		description string,
	) (services.SubResource, error)
	Remove(id *string) error
}

type AccessGroupService interface {
	Create(
		name *string,
		description string,
	) (services.AccessGroup, error)
	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]services.AccessGroup, bool, error)
	Get(id *string) (services.AccessGroup, error)
	Modify(
		id *string,
		name *string,
		description string,
	) (services.AccessGroup, error)
	Remove(id *string) error
}

type AccessGroupPermissionService interface {
	Create(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (bool, error)
	GetMany(
		pageNumber *uint,
		itemsPerPage uint,
	) ([]services.AccessGroupPremission, bool, error)
	GetByAccessGroup(accessGroupId *string) ([]services.AccessGroupPremission, bool, error)
	GetByResource(resourceId *string) ([]services.AccessGroupPremission, bool, error)
	GetBySubResource(subResourceId *string) ([]services.AccessGroupPremission, bool, error)
	GetBySubResourceAction(subResourceActionId *string) ([]services.AccessGroupPremission, bool, error)
	ModifyByAccessGroup(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (string, error)
	ModifyBySubResource(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (string, error)
	ModifyByResource(
		accessGroupId *string,
		resourceId *string,
		subResourceId *string,
		subResourceActionId *string,
	) (string, error)
	Remove(services.AccessGroupPremission) error
}
