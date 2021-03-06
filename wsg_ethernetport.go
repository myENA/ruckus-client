package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGEthernetPortCreateEthernetPortProfile
//
// Definition: ethernetPort_createEthernetPortProfile
type WSGEthernetPortCreateEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallAVCEnabled
	// Application Recognition & Control enabled
	FirewallAVCEnabled *bool `json:"firewallAVCEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUrlFilteringPolicyEnabled
	// URL Filtering Policy enabled
	FirewallUrlFilteringPolicyEnabled *bool `json:"firewallUrlFilteringPolicyEnabled,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy Id
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

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
	Type *string `json:"type"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - default:1
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty"`

	// UserSidePortEnabled
	// User side port enabled.
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	// X8021X
	// Constraints:
	//    - required
	X8021X *WSGAPModelLanPort8021X `json:"_8021X"`
}

func NewWSGEthernetPortCreateEthernetPortProfile() *WSGEthernetPortCreateEthernetPortProfile {
	m := new(WSGEthernetPortCreateEthernetPortProfile)
	return m
}

// WSGEthernetPortProfile
//
// Definition: ethernetPort_ethernetPortProfile
type WSGEthernetPortProfile struct {
	// AntiSpoofingEnabled
	// Anti-Spoofing enabled
	AntiSpoofingEnabled *bool `json:"antiSpoofingEnabled,omitempty"`

	// ArpRequestRateLimit
	// ARP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallAVCEnabled
	// Application Recognition & Control enabled
	FirewallAVCEnabled *bool `json:"firewallAVCEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUrlFilteringPolicyEnabled
	// URL Filtering Policy enabled
	FirewallUrlFilteringPolicyEnabled *bool `json:"firewallUrlFilteringPolicyEnabled,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty"`

	// Id
	// identifier of the ethernet port profile
	Id *string `json:"id,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy Id
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// Type
	// type of the ethernet port profile
	// Constraints:
	//    - oneof:[AccessPort,TrunkPort,GeneralPort]
	Type *string `json:"type,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty"`

	// UserSidePortEnabled
	// User side port enabled.
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

type WSGEthernetPortProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGEthernetPortProfile
}

func newWSGEthernetPortProfileAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGEthernetPortProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGEthernetPortProfileAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGEthernetPortProfile)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGEthernetPortProfile() *WSGEthernetPortProfile {
	m := new(WSGEthernetPortProfile)
	return m
}

// WSGEthernetPortModifyEthernetPortProfile
//
// Definition: ethernetPort_modifyEthernetPortProfile
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
	ArpRequestRateLimit *int `json:"arpRequestRateLimit,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DhcpRequestRateLimit
	// DHCP packets request rate limit, default value will be 15 if both rate limit not being set.
	// Constraints:
	//    - min:0
	//    - max:100
	DhcpRequestRateLimit *int `json:"dhcpRequestRateLimit,omitempty"`

	DynamicVlanEnabled *bool `json:"dynamicVlanEnabled,omitempty"`

	// FirewallAVCEnabled
	// Application Recognition & Control enabled
	FirewallAVCEnabled *bool `json:"firewallAVCEnabled,omitempty"`

	// FirewallProfileId
	// Firewall Profile Id
	FirewallProfileId *string `json:"firewallProfileId,omitempty"`

	// FirewallUrlFilteringPolicyEnabled
	// URL Filtering Policy enabled
	FirewallUrlFilteringPolicyEnabled *bool `json:"firewallUrlFilteringPolicyEnabled,omitempty"`

	// GuestVlan
	// Constraints:
	//    - min:1
	//    - max:4094
	GuestVlan *int `json:"guestVlan,omitempty"`

	IpsecProfile *WSGCommonGenericRef `json:"ipsecProfile,omitempty"`

	// L2AccessControlPolicyId
	// L2 Access Control Policy Id
	L2AccessControlPolicyId *string `json:"l2AccessControlPolicyId,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	Qinq *WSGCommonQinq `json:"qinq,omitempty"`

	// TunnelEnabled
	// tunnelEnabled of the ethernet port profile for AccessPort type
	TunnelEnabled *bool `json:"tunnelEnabled,omitempty"`

	TunnelProfile *WSGCommonGenericRef `json:"tunnelProfile,omitempty"`

	// UntagId
	// VLAN Untag ID of the ethernet port profile
	// Constraints:
	//    - min:1
	//    - max:4094
	UntagId *int `json:"untagId,omitempty"`

	// UserSidePortEnabled
	// User side port enabled.
	UserSidePortEnabled *bool `json:"userSidePortEnabled,omitempty"`

	// UserSidePortMaxClient
	// Number of wired clients allowed to connect to a particular user side port, default value will be 8 if the value not being set.
	// Constraints:
	//    - default:8
	//    - min:1
	//    - max:32
	UserSidePortMaxClient *int `json:"userSidePortMaxClient,omitempty"`

	// VlanMembers
	// VLAN Members of the ethernet port profile
	VlanMembers *string `json:"vlanMembers,omitempty"`

	X8021X *WSGAPModelLanPort8021X `json:"_8021X,omitempty"`
}

func NewWSGEthernetPortModifyEthernetPortProfile() *WSGEthernetPortModifyEthernetPortProfile {
	m := new(WSGEthernetPortModifyEthernetPortProfile)
	return m
}

// WSGEthernetPortProfileList
//
// Definition: ethernetPort_profileList
type WSGEthernetPortProfileList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGEthernetPortProfileListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGEthernetPortProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGEthernetPortProfileList
}

func newWSGEthernetPortProfileListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGEthernetPortProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGEthernetPortProfileListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGEthernetPortProfileList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGEthernetPortProfileList() *WSGEthernetPortProfileList {
	m := new(WSGEthernetPortProfileList)
	return m
}

// WSGEthernetPortProfileListType
//
// Definition: ethernetPort_profileListType
type WSGEthernetPortProfileListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGEthernetPortProfileListType() *WSGEthernetPortProfileListType {
	m := new(WSGEthernetPortProfileListType)
	return m
}
