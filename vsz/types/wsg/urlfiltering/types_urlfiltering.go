package urlfiltering

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - required
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel" validate:"required,oneof=NO_ADULT CLEAN_AND_SAFE CHILD_AND_STUDENT_FRIENDLY STRICT CUSTOM"`

	// Name
	// Constraints:
	//    - required
	Name *common.NormalName `json:"name" validate:"required"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

func NewCreateUrlFilteringPolicy() *CreateUrlFilteringPolicy {
	createUrlFilteringPolicyType := new(CreateUrlFilteringPolicy)
	return createUrlFilteringPolicyType
}

func NewCreateUrlFilteringPolicyWithDefaults() *CreateUrlFilteringPolicy {
	createUrlFilteringPolicyType := new(CreateUrlFilteringPolicy)
	return createUrlFilteringPolicyType
}

type DeleteBulk struct {
	IdList common.IdList `json:"idList,omitempty"`
}

func NewDeleteBulk() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
}

func NewDeleteBulkWithDefaults() *DeleteBulk {
	deleteBulkType := new(DeleteBulk)
	return deleteBulkType
}

type ModifyUrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - nullable
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty" validate:"omitempty,oneof=NO_ADULT CLEAN_AND_SAFE CHILD_AND_STUDENT_FRIENDLY STRICT CUSTOM"`

	Name *common.NormalName `json:"name,omitempty"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

func NewModifyUrlFilteringPolicy() *ModifyUrlFilteringPolicy {
	modifyUrlFilteringPolicyType := new(ModifyUrlFilteringPolicy)
	return modifyUrlFilteringPolicyType
}

func NewModifyUrlFilteringPolicyWithDefaults() *ModifyUrlFilteringPolicy {
	modifyUrlFilteringPolicyType := new(ModifyUrlFilteringPolicy)
	return modifyUrlFilteringPolicyType
}

type UrlFilteringBlockCategoriesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UrlFilteringBlockCategory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewUrlFilteringBlockCategoriesList() *UrlFilteringBlockCategoriesList {
	urlFilteringBlockCategoriesListType := new(UrlFilteringBlockCategoriesList)
	return urlFilteringBlockCategoriesListType
}

func NewUrlFilteringBlockCategoriesListWithDefaults() *UrlFilteringBlockCategoriesList {
	urlFilteringBlockCategoriesListType := new(UrlFilteringBlockCategoriesList)
	return urlFilteringBlockCategoriesListType
}

type UrlFilteringBlockCategory struct {
	// Id
	// Identifier of the URL Filtering Category
	Id *int `json:"id,omitempty"`

	// Name
	// name of the URL Filtering Category
	Name *string `json:"name,omitempty"`
}

func NewUrlFilteringBlockCategory() *UrlFilteringBlockCategory {
	urlFilteringBlockCategoryType := new(UrlFilteringBlockCategory)
	return urlFilteringBlockCategoryType
}

func NewUrlFilteringBlockCategoryWithDefaults() *UrlFilteringBlockCategory {
	urlFilteringBlockCategoryType := new(UrlFilteringBlockCategory)
	return urlFilteringBlockCategoryType
}

type UrlFilteringPolicy struct {
	// Blacklist
	// The blacklist of the URL Filtering policy
	Blacklist []string `json:"blacklist,omitempty"`

	// BlockCategories
	// The block category IDs of the URL Filtering policy
	BlockCategories []int `json:"blockCategories,omitempty"`

	// CreateDateTime
	// Timestamp of being created
	CreateDateTime *int `json:"createDateTime,omitempty"`

	// CreatorId
	// Creator ID
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// Creator Name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DomainId
	// Identifier of the System (root) domain or partner managed domain to which the URL filtering policy belongs
	DomainId *string `json:"domainId,omitempty"`

	// FilteringLevel
	// The filtering level of the URL Filtering policy
	// Constraints:
	//    - nullable
	//    - oneof:[NO_ADULT,CLEAN_AND_SAFE,CHILD_AND_STUDENT_FRIENDLY,STRICT,CUSTOM]
	FilteringLevel *string `json:"filteringLevel,omitempty" validate:"omitempty,oneof=NO_ADULT CLEAN_AND_SAFE CHILD_AND_STUDENT_FRIENDLY STRICT CUSTOM"`

	// Id
	// Identifier of the URL Filtering policy
	Id *string `json:"id,omitempty"`

	// ModifiedDateTime
	// Timestamp of being modified
	ModifiedDateTime *int `json:"modifiedDateTime,omitempty"`

	// ModifierId
	// Modifier ID
	ModifierId *string `json:"modifierId,omitempty"`

	// ModifierUsername
	// Modifier Name
	ModifierUsername *string `json:"modifierUsername,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// SafeSearchBingDns
	// Bing DNS for safe search of the URL Filtering policy
	SafeSearchBingDns *string `json:"safeSearchBingDns,omitempty"`

	// SafeSearchBingEnabled
	// Enable Bing safe search of the URL Filtering policy
	SafeSearchBingEnabled *bool `json:"safeSearchBingEnabled,omitempty"`

	// SafeSearchGoogleDns
	// Google DNS for safe search of the URL Filtering policy
	SafeSearchGoogleDns *string `json:"safeSearchGoogleDns,omitempty"`

	// SafeSearchGoogleEnabled
	// Enable Google safe search of the URL Filtering policy
	SafeSearchGoogleEnabled *bool `json:"safeSearchGoogleEnabled,omitempty"`

	// SafeSearchYouTubeDns
	// YouTube DNS for safe search of the URL Filtering policy
	SafeSearchYouTubeDns *string `json:"safeSearchYouTubeDns,omitempty"`

	// SafeSearchYouTubeEnabled
	// Enable YouTube safe search of the URL Filtering policy
	SafeSearchYouTubeEnabled *bool `json:"safeSearchYouTubeEnabled,omitempty"`

	// TenantId
	// Tenant Id
	TenantId *string `json:"tenantId,omitempty"`

	// Whitelist
	// The whitelist of the URL Filtering policy
	Whitelist []string `json:"whitelist,omitempty"`
}

func NewUrlFilteringPolicy() *UrlFilteringPolicy {
	urlFilteringPolicyType := new(UrlFilteringPolicy)
	return urlFilteringPolicyType
}

func NewUrlFilteringPolicyWithDefaults() *UrlFilteringPolicy {
	urlFilteringPolicyType := new(UrlFilteringPolicy)
	return urlFilteringPolicyType
}

type UrlFilteringPolicyList struct {
	Extra *common.RbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*UrlFilteringPolicy `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewUrlFilteringPolicyList() *UrlFilteringPolicyList {
	urlFilteringPolicyListType := new(UrlFilteringPolicyList)
	return urlFilteringPolicyListType
}

func NewUrlFilteringPolicyListWithDefaults() *UrlFilteringPolicyList {
	urlFilteringPolicyListType := new(UrlFilteringPolicyList)
	return urlFilteringPolicyListType
}
