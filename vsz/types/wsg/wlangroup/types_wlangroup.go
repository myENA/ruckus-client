package wlangroup

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type CreateWlanGroup struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ModifyWlanGroup struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type ModifyWlanGroupMember struct {
	// AccessVlan
	// Access VLAN
	AccessVlan *int `json:"accessVlan,omitempty"`

	// NasId
	// NAS-ID
	NasId *string `json:"nasId,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}

type WlanGroup struct {
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

	// Id
	// Identifier of the WLAN group
	Id *string `json:"id,omitempty"`

	// Members
	// Members of the WLAN group
	Members []*WlanMember `json:"members,omitempty"`

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

	// ZoneId
	// Identifier of the zone to which the WLAN group belongs
	ZoneId *string `json:"zoneId,omitempty"`
}

type WlanGroupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WlanGroup `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WlanMember struct {
	// AccessVlan
	// Access VLAN
	AccessVlan *int `json:"accessVlan,omitempty"`

	// Id
	// Identifier of the WLAN
	Id *string `json:"id,omitempty"`

	// NasId
	// NAS-ID
	NasId *string `json:"nasId,omitempty"`

	VlanPooling *common.GenericRef `json:"vlanPooling,omitempty"`
}