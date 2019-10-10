package portalservice

// API Version: v8_1

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type ConnectionCapability struct {
	// PortNumber
	// Port number of connection capability
	PortNumber *float64 `json:"portNumber" validate:"required,gte=0.000000,lte=65535.000000"`

	// ProtocolName
	// Protocol aame of connection capability
	ProtocolName *string `json:"protocolName" validate:"required"`

	// ProtocolNumber
	// Protocol number of connection capability
	ProtocolNumber *float64 `json:"protocolNumber" validate:"required,gte=0.000000,lte=254.000000"`

	// Status
	// Status of connection capability
	Status *string `json:"status" validate:"required,oneof=CLOSED OPEN UNKNOWN"`
}

type CreateGuestAccess struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession" validate:"required"`
}

type CreateHotspot20VenueProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkSpeedInKbps *LinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	Group *string `json:"group,omitempty" validate:"oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	Name *common.NormalName `json:"name" validate:"required"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *LinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames" validate:"required"`
}

type CreateHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	AccessNetworkType *string `json:"accessNetworkType" validate:"required,oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// IdentityProviders
	// Ddentity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption" validate:"required"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 WLAN profile
	Ipv4AddressType *string `json:"ipv4AddressType" validate:"required,oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	Ipv6AddressType *string `json:"ipv6AddressType" validate:"required,oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	Name *common.NormalName `json:"name" validate:"required"`

	Operator *common.GenericRef `json:"operator" validate:"required"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`
}

type CreateHotspotExternal struct {
	// BackupPortalUrl
	// Backup Portal URL of the Hotspot
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat" validate:"required"`

	Name *common.NormalName `json:"name" validate:"required"`

	PortalUrl *common.NormalURL `json:"portalUrl" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport" validate:"required,oneof=None Enabled"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type CreateHotspotInternal struct {
	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat" validate:"required"`

	Name *common.NormalName `json:"name" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport" validate:"required,oneof=None Enabled"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type CreateHotspotSmartClientOnly struct {
	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat" validate:"required"`

	Name *common.NormalName `json:"name" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client
	// application.
	SmartClientInfo *string `json:"smartClientInfo" validate:"required"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type CreateL2ACL struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all
	// stations listed below
	Restriction *string `json:"restriction" validate:"required,oneof=ALLOW BLOCK"`

	RuleMacs []common.Mac `json:"ruleMacs,omitempty"`
}

type CreateWebAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name" validate:"required"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	WebAuthenticationPortalCustomization *common.WebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
}

type CreateWechat struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl" validate:"required"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList" validate:"required"`

	Description *common.Description `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination" validate:"required"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"gte=1,lte=14399"`

	Name *common.NormalName `json:"name" validate:"required"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList" validate:"required"`
}

type DefaultConnectionCapability struct {
	// PortNumber
	// Port number of connection capability, cannot be modified
	PortNumber *float64 `json:"portNumber" validate:"required,gte=0.000000,lte=65535.000000"`

	// ProtocolName
	// Protocol aame of connection capability, cannot be modified
	ProtocolName *string `json:"protocolName" validate:"required"`

	// ProtocolNumber
	// Protocol number of connection capability, cannot be modified
	ProtocolNumber *float64 `json:"protocolNumber" validate:"required,gte=0.000000,lte=254.000000"`

	// Status
	// Status of connection capability
	Status *string `json:"status" validate:"required,oneof=CLOSED OPEN UNKNOWN"`
}

type DnatPortMapping struct {
	// DestPort
	// Destination port
	DestPort *int `json:"destPort,omitempty" validate:"gte=0,lte=65535"`

	// SourcePort
	// Source port
	SourcePort *int `json:"sourcePort,omitempty" validate:"gte=0,lte=65535"`
}

type GuestAccess struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the guest access profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// ZoneId
	// Identifier of the zone which the guest access profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type Hotspot struct {
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// Id
	// Identifier of the Hotspot
	Id *string `json:"id,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	// PortalType
	// Portal type of the Hotspot
	PortalType *string `json:"portalType,omitempty" validate:"oneof=Internal External"`

	PortalUrl *common.NormalURL `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client
	// application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport,omitempty" validate:"oneof=None Enabled SmartClientOnly"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type Hotspot20VeuneProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkSpeedInKbps *LinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	Group *string `json:"group,omitempty" validate:"oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	// Id
	// Identifier of the Hotspot 2.0 venue profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *LinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 venue profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type Hotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	AccessNetworkType *string `json:"accessNetworkType,omitempty" validate:"oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*ConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the Hotspot 2.0 WLAN profile
	Id *string `json:"id,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the v WLAN profile
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty" validate:"oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 WLAN profile
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty" validate:"oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	Name *common.NormalName `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`

	// ZoneId
	// Identifier of the zone which the Hotspot 2.0 WLAN profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type L2ACL struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// identifier of the L2 Access Control
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all
	// stations listed below
	Restriction *string `json:"restriction,omitempty" validate:"oneof=ALLOW BLOCK"`

	RuleMacs []common.Mac `json:"ruleMacs,omitempty"`

	// ZoneId
	// identifier of the zone which the L2 Access Control belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

// LinkSpeedInKbps
//
// Link Speed of the Hotspot 2.0 venue profile
type LinkSpeedInKbps float64

// MacAddressFormatSetting
//
// mac address format of redirection,the format define: 0(aabbccddeeff), 1(AA-BB-CC-DD-EE-FF), 2(AA:BB:CC:DD:EE:FF),
// 3(AABBCCDDEEFF), 4(aa-bb-cc-dd-ee-ff), 5(aa:bb:cc:dd:ee:ff)
type MacAddressFormatSetting int

type ModifyGuestAccess struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	SmsGateway *common.GenericRef `json:"smsGateway,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`
}

type ModifyHotspot struct {
	BackupPortalUrl *string `json:"backupPortalUrl,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// HttpsRedirect
	// HTTPS Redirect is disable or not
	HttpsRedirect *bool `json:"httpsRedirect,omitempty"`

	// InternalNode
	// Internal Node of the Hotspot
	InternalNode *string `json:"internalNode,omitempty"`

	Location *PortalLocation `json:"location,omitempty"`

	MacAddressFormat *MacAddressFormatSetting `json:"macAddressFormat,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalCustomization *common.PortalCustomization `json:"portalCustomization,omitempty"`

	PortalUrl *common.NormalURL `json:"portalUrl,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	// SignatureSigningKey
	// Signature Signing Key of the Hotspot
	SignatureSigningKey *string `json:"signatureSigningKey,omitempty"`

	// SmartClientInfo
	// Smart client info of the Hotspot. Type instructions for enabling users to log on using the Smart Client
	// application.
	SmartClientInfo *string `json:"smartClientInfo,omitempty"`

	// SmartClientSupport
	// Smart client support of the Hotspot
	SmartClientSupport *string `json:"smartClientSupport,omitempty" validate:"oneof=None Enabled"`

	// TrafficClassProfileId
	// Traffic Class Profile of the Hotspot
	TrafficClassProfileId *string `json:"trafficClassProfileId,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	// WalledGardens
	// Walled garden map set of the Hotspot. Unauthenticated users are allowed to access the following
	// destinations. Format: - IP (e.g. 10.11.12.13) - IP Range (e.g. 10.11.12.13-10.11.12.15) - CIDR (e.g.
	// 10.11.12.100/28) - IP and mask (e.g. 10.11.12.13 255.255.255.0) - Precise web site (e.g.
	// www.ruckus.com) - Web site with special regular expression like    - *.amazon.com    - *.com
	WalledGardens []string `json:"walledGardens,omitempty"`
}

type ModifyHotspot20VenueProfile struct {
	Description *common.Description `json:"description,omitempty"`

	DownlinkSpeedInKbps *LinkSpeedInKbps `json:"downlinkSpeedInKbps,omitempty"`

	// Group
	// Category group of the Hotspot 2.0 venue profile
	Group *string `json:"group,omitempty" validate:"oneof=Unspecified Assembly Business Educational FactoryAndIndustrial Institutional Mercantile Residential Storage UtilityAndMiscellaneous Vehicular Outdoor"`

	Name *common.NormalName `json:"name,omitempty"`

	// Type
	// Category type of the Hotspot 2.0 venue profile
	Type *string `json:"type,omitempty"`

	UplinkSpeedInKbps *LinkSpeedInKbps `json:"uplinkSpeedInKbps,omitempty"`

	VenueNames []*VenueName `json:"venueNames,omitempty"`
}

type ModifyHotspot20WlanProfile struct {
	// AccessNetworkType
	// Access network type of the Hotspot 2.0 WLAN profile
	AccessNetworkType *string `json:"accessNetworkType,omitempty" validate:"oneof=CHARGEABLE_PUBLIC FREE_PUBLIC PERSONAL_DEVICE PRIVATE PRIVATE_WITH_GUEST TEST WILDCARD"`

	// ConnectionCapabilities
	// Default connection capabilities of the Hotspot 2.0 WLAN profile
	ConnectionCapabilities []*DefaultConnectionCapability `json:"connectionCapabilities,omitempty"`

	// CustomConnectionCapabilities
	// Custom connection capabilities of the Hotspot 2.0 WLAN profile
	CustomConnectionCapabilities []*ConnectionCapability `json:"customConnectionCapabilities,omitempty"`

	DefaultIdentityProvider *common.GenericRef `json:"defaultIdentityProvider,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// IdentityProviders
	// Identity providers of the Hotspot 2.0 WLAN profile
	IdentityProviders []*common.GenericRef `json:"identityProviders,omitempty"`

	// InternetOption
	// Internet option of the Hotspot 2.0 WLAN profile
	InternetOption *bool `json:"internetOption,omitempty"`

	// Ipv4AddressType
	// IPv4 address type of the Hotspot 2.0 Wlan profile
	Ipv4AddressType *string `json:"ipv4AddressType,omitempty" validate:"oneof=UNAVAILABLE PUBLIC PORT_RESTRICTED SINGLE_NATED_PRIVATE DOUBLE_NATED_PRIVATE PORT_RESTRICTED_AND_SINGLE_NATED PORT_RESTRICTED_AND_DOUBLE_NATED UNKNOWN"`

	// Ipv6AddressType
	// IPv6 address type of the Hotspot 2.0 Wlan profile
	Ipv6AddressType *string `json:"ipv6AddressType,omitempty" validate:"oneof=UNAVAILABLE AVAILABLE UNKNOWN"`

	Name *common.NormalName `json:"name,omitempty"`

	Operator *common.GenericRef `json:"operator,omitempty"`

	SignupSsid *common.GenericRef `json:"signupSsid,omitempty"`
}

type ModifyL2ACL struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	// Restriction
	// restriction of the L2 Access Control, ALLOW: Only allow all stations listed below, BLOCK:Only block all
	// stations listed below
	Restriction *string `json:"restriction,omitempty" validate:"oneof=ALLOW BLOCK"`

	RuleMacs []common.Mac `json:"ruleMacs,omitempty"`
}

type ModifyWebAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalLanguage *common.PortalLanguage `json:"portalLanguage,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *common.WebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`
}

type ModifyWechat struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"gte=1,lte=14399"`

	Name *common.NormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}

type PortalLocation struct {
	// Id
	// Portal location id
	Id *string `json:"id,omitempty"`

	// Name
	// Portal location name
	Name *string `json:"name,omitempty"`
}

type PortalRedirect struct {
	Url *common.NormalURL `json:"url,omitempty"`
}

type PortalServiceList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*PortalServiceListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type PortalServiceListType struct {
	// Id
	// Identifier of the service
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`
}

type UserSession struct {
	// GracePeriodInMin
	// Grace period in minutes
	GracePeriodInMin *int `json:"gracePeriodInMin,omitempty" validate:"gte=1,lte=14399"`

	// TimeoutInMin
	// Time out value in minutes
	TimeoutInMin *int `json:"timeoutInMin,omitempty" validate:"gte=2,lte=14400"`
}

type VenueName struct {
	Language *common.LanguageName `json:"language" validate:"required"`

	// Name
	// Venue name
	Name *string `json:"name" validate:"required"`
}

type WebAuthentication struct {
	Description *common.Description `json:"description,omitempty"`

	// Id
	// Identifier of the web authentication profile
	Id *string `json:"id,omitempty"`

	Name *common.NormalName `json:"name,omitempty"`

	PortalLanguage *common.PortalLanguage `json:"portalLanguage,omitempty"`

	Redirect *PortalRedirect `json:"redirect,omitempty"`

	UserSession *UserSession `json:"userSession,omitempty"`

	WebAuthenticationPortalCustomization *common.WebAuthenticationPortalCustomization `json:"webAuthenticationPortalCustomization,omitempty"`

	// ZoneId
	// Identifier of the zone which the web authentication profile belongs to
	ZoneId *string `json:"zoneId,omitempty"`
}

type WechatConfiguration struct {
	// AuthUrl
	// Authentication URL of the wechat profile
	AuthUrl *string `json:"authUrl,omitempty"`

	// BlackList
	// Black list of the wechat profile
	BlackList *string `json:"blackList,omitempty"`

	Description *common.Description `json:"description,omitempty"`

	// DnatDestination
	// DNAT destination of the wechat profile
	DnatDestination *string `json:"dnatDestination,omitempty"`

	// DnatPortMapping
	// DNAT Port Mapping of the wechat profile
	DnatPortMapping []*DnatPortMapping `json:"dnatPortMapping,omitempty"`

	// GracePeriod
	// Grace period of the wechat profile
	GracePeriod *int `json:"gracePeriod,omitempty" validate:"gte=1,lte=14399"`

	Name *common.NormalName `json:"name,omitempty"`

	// WhiteList
	// White list of the wechat profile
	WhiteList []string `json:"whiteList,omitempty"`
}
