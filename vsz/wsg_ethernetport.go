package vsz

// API Version: v8_1

type WSGEthernetPortCreateEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty" validate:"gte=1,lte=4094"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - required
	//    - default:'TrunkPort'
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type" validate:"required,oneof=AccessPort TrunkPort GeneralPort"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty" validate:"gte=1,lte=4094"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	// X8021X
	// Constraints:
	//    - required
	X8021X *WSGAPModelLanPort8021X `json:"_8021X" validate:"required"`
}

type WSGEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty" validate:"gte=1,lte=4094"`

	// Id
	// identifier of the ethernet port profile
	Id *string `json:"id,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type,omitempty" validate:"oneof=AccessPort TrunkPort GeneralPort"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty" validate:"gte=1,lte=4094"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

type WSGEthernetPortModifyEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - default:15
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty" validate:"gte=0,lte=100"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty" validate:"gte=1,lte=4094"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty" validate:"gte=1,lte=4094"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

type WSGEthernetPortProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGEthernetPortProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGEthernetPortProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}