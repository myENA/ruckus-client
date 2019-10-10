package scguser

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateScgUser struct {
	// AccountLockout
	// The number of successive failures before a lockout will occur. (System default admin ONLY)
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// Id
	// User id
	Id *string `json:"id,omitempty"`

	// LockoutDuration
	// The duration for which the account is automatically locked without administrative intervention. (System
	// default admin ONLY)
	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// MinimumPasswordLength
	// The minimum length of the password for the account. (System default admin ONLY)
	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// NewPassphrase
	// User login passphrase
	NewPassphrase *string `json:"newPassphrase" validate:"required"`

	// PasswordExpiration
	// A simple timer that forces the administrator to change their password regularly. (System default admin
	// ONLY)
	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	// PasswordReuse
	// A validation the prevents reuse of the same password(s). (System default admin ONLY)
	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	// SessionIdle
	// A period of idle used to invalid that session. (System default admin ONLY)
	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName" validate:"required"`
}

type GetScgUser struct {
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// Enabled
	// User enabled or not
	Enabled *int `json:"enabled,omitempty" validate:"omitempty,oneof=0 1"`

	// Id
	// User id
	Id *string `json:"id,omitempty"`

	// Locked
	// User locked or not (0:unlocked/1:locked)
	Locked *int `json:"locked,omitempty" validate:"omitempty,oneof=0 1 2"`

	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"omitempty,gte=1,lte=1440"`

	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// TenantUUID
	// Tenant id
	TenantUUID *string `json:"tenantUUID,omitempty"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

type ModifyScgUser struct {
	AccountLockout *int `json:"accountLockout,omitempty" validate:"omitempty,gte=1,lte=100"`

	// DomainId
	// Domain id
	DomainId *string `json:"domainId,omitempty"`

	// Email
	// User email
	Email *string `json:"email,omitempty"`

	// Id
	// User id
	Id *string `json:"id" validate:"required"`

	LockoutDuration *int `json:"lockoutDuration,omitempty" validate:"omitempty,gte=1,lte=1440"`

	MinimumPasswordLength *int `json:"minimumPasswordLength,omitempty" validate:"omitempty,gte=8,lte=64"`

	// NewPassphrase
	// User new login passphrase
	NewPassphrase *string `json:"newPassphrase,omitempty"`

	// Passphrase
	// User login passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	PasswordExpiration *int `json:"passwordExpiration,omitempty" validate:"omitempty,gte=1,lte=365"`

	PasswordReuse *int `json:"passwordReuse,omitempty" validate:"omitempty,gte=1,lte=6"`

	// Phone
	// User phone
	Phone *string `json:"phone,omitempty"`

	// RealName
	// User real name
	RealName *string `json:"realName,omitempty"`

	SessionIdle *int `json:"sessionIdle,omitempty" validate:"omitempty,gte=1,lte=1440"`

	// Title
	// User title
	Title *string `json:"title,omitempty"`

	// UserName
	// User name
	UserName *string `json:"userName,omitempty"`
}

type PatchScgUserGroup struct {
	AccountSecurityProfileId *string `json:"accountSecurityProfileId,omitempty"`

	CreateDateTime *int `json:"createDateTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	Id *string `json:"id,omitempty"`

	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	ModifierId *string `json:"modifierId,omitempty"`

	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *string `json:"name,omitempty"`

	Permissions []*ScgUserGroupPermissionWithoutDetailItems `json:"permissions,omitempty"`

	ResourceGroups []*ScgUserGroupResourceGroup `json:"resourceGroups,omitempty"`

	Role *string `json:"role,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	Users []*GetScgUser `json:"users,omitempty"`
}

type QueryCriteria struct{}

type ScgUserAuditId struct {
	// Id
	// the identifier of the SCG user
	Id *string `json:"id,omitempty"`
}

type ScgUserGroup struct {
	AccountSecurityProfileId *string `json:"accountSecurityProfileId" validate:"required"`

	AccountSecurityProfileName *string `json:"accountSecurityProfileName,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// Description
	// User group description
	Description *string `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// User group Id
	Id *string `json:"id,omitempty"`

	IsFactoryDefault *bool `json:"isFactoryDefault,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	// Name
	// User group name
	Name *string `json:"name" validate:"required"`

	// Permissions
	// Permission list
	Permissions []*ScgUserGroupPermissionWithoutDetailItems `json:"permissions" validate:"required"`

	// ResourceGroups
	// Resource group id list
	ResourceGroups []*ScgUserGroupResourceGroup `json:"resourceGroups" validate:"required"`

	// Role
	// User group role
	Role *string `json:"role" validate:"required"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Users
	// Users in this user group
	Users []*GetScgUser `json:"users,omitempty"`
}

type ScgUserGroupAuditId struct {
	// Id
	// the identifier of the SCG user group
	Id *string `json:"id,omitempty"`
}

type ScgUserGroupList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ScgUserGroupPermission struct {
	// Access
	// Access level
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=READ MODIFY FULL_ACCESS"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	Ids []string `json:"ids,omitempty"`

	// Items
	// Resource items
	Items []*ScgUserGroupPermissionItemsType `json:"items,omitempty"`

	// ItemsDescription
	// Descriptions of Resource items
	ItemsDescription []string `json:"itemsDescription,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupPermissionItemsType struct {
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=NA READ MODIFY FULL_ACCESS"`

	Display *string `json:"display,omitempty"`

	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupPermissionList struct {
	// Extra
	// Any additional response data.
	Extra *ScgUserGroupPermissionListExtraType `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroupPermission `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// ScgUserGroupPermissionListExtraType
//
// Any additional response data.
type ScgUserGroupPermissionListExtraType struct {
	// IsSuperAdmin
	// whether or not current user is a 'Super Admin' that possesses all 6 permission categories with
	// 'FULL_ACCESS'
	IsSuperAdmin *bool `json:"isSuperAdmin,omitempty"`

	// IsSuperAdminOfDomain
	// whether or not current user is a 'Super Admin of Partner Domain' that possesses all 6 permission
	// categories with 'FULL_ACCESS'
	IsSuperAdminOfDomain *bool `json:"isSuperAdminOfDomain,omitempty"`
}

type ScgUserGroupPermissionWithoutDetailItems struct {
	// Access
	// Access level
	Access *string `json:"access,omitempty" validate:"omitempty,oneof=READ MODIFY FULL_ACCESS"`

	// Display
	// Resource display name
	Display *string `json:"display,omitempty"`

	// Ids
	// Resource id list
	Ids []string `json:"ids,omitempty"`

	// Resource
	// Resource type
	Resource *string `json:"resource,omitempty"`
}

type ScgUserGroupResourceGroup struct {
	// Id
	// the identifier of the SCG resource group
	Id *string `json:"id,omitempty"`

	// Name
	// the name of the SCG resource group
	Name *string `json:"name,omitempty"`

	// Type
	// the type of SCG resource group
	Type *string `json:"type,omitempty" validate:"omitempty,oneof=DOMAIN ZONE APGROUP"`
}

type ScgUserGroupRoleLabelValue struct {
	// Label
	// Role display name
	Label *string `json:"label,omitempty"`

	// Value
	// Role value
	Value *string `json:"value,omitempty"`
}

type ScgUserGroupRoleLabelValueList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*ScgUserGroupRoleLabelValue `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type ScgUserList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*GetScgUser `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}
