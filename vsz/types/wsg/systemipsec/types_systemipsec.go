package systemipsec

// API Version: v8_1

type GetResult struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []*Proposal `json:"espProposals,omitempty"`

	// EspRekeyDisabled
	// Disable rekey mechanisam of Encapsulating Security Payload
	EspRekeyDisabled *bool `json:"espRekeyDisabled,omitempty"`

	// EspRekeyTime
	// Rekey time of Encapsulating Security Payload
	EspRekeyTime *int `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Rekey time unit of Encapsulating Security Payload
	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	// IkeProposals
	// Proposal of Internet Key Exchange
	IkeProposals []*Proposal `json:"ikeProposals,omitempty"`

	// IkeRekeyDisabled
	// Disable rekey mechanisam of Internet Key Exchange
	IkeRekeyDisabled *bool `json:"ikeRekeyDisabled,omitempty"`

	// IkeRekeyTime
	// Rekey time of Internet Key Exchange
	IkeRekeyTime *int `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Rekey time unit of Internet Key Exchange
	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	// IpSecEnabled
	// Enable System IPSec
	IpSecEnabled *bool `json:"ipSecEnabled,omitempty"`

	OcspEnabled *bool `json:"ocspEnabled,omitempty"`

	OcspServerUri *string `json:"ocspServerUri,omitempty"`

	// PreSharedKey
	// Pre-shared key
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// RemoteId
	// The IP of IPSec receiver
	RemoteId *string `json:"remoteId,omitempty"`

	// ScgCertId
	// SCG client certification id.
	ScgCertId *string `json:"scgCertId,omitempty"`

	// SecurityGateway
	// Security gateway IP
	SecurityGateway *string `json:"securityGateway,omitempty"`

	// SubnetMask
	// Subnet Mask of security gateway
	SubnetMask *string `json:"subnetMask,omitempty"`

	TrustChainProfileId *string `json:"trustChainProfileId,omitempty"`
}

func NewGetResult() *GetResult {
	getResultType := new(GetResult)
	return getResultType
}

func NewGetResultWithDefaults() *GetResult {
	getResultType := new(GetResult)
	return getResultType
}

type Proposal struct {
	// AuthAlg
	// Authentication algorithm
	// Constraints:
	//    - required
	AuthAlg *string `json:"authAlg" validate:"required"`

	// EncAlg
	// Encrytion algorithm
	// Constraints:
	//    - required
	EncAlg *string `json:"encAlg" validate:"required"`
}

func NewProposal() *Proposal {
	proposalType := new(Proposal)
	return proposalType
}

func NewProposalWithDefaults() *Proposal {
	proposalType := new(Proposal)
	return proposalType
}

type Update struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []*Proposal `json:"espProposals,omitempty"`

	// EspRekeyDisabled
	// Disable rekey mechanisam of Encapsulating Security Payload
	EspRekeyDisabled *bool `json:"espRekeyDisabled,omitempty"`

	// EspRekeyTime
	// Rekey time of Encapsulating Security Payload
	EspRekeyTime *int `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Rekey time unit of Encapsulating Security Payload
	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	// IkeProposals
	// Proposal of Internet Key Exchange
	IkeProposals []*Proposal `json:"ikeProposals,omitempty"`

	// IkeRekeyDisabled
	// Disable rekey mechanisam of Internet Key Exchange
	IkeRekeyDisabled *bool `json:"ikeRekeyDisabled,omitempty"`

	// IkeRekeyTime
	// Rekey time of Internet Key Exchange
	IkeRekeyTime *int `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Rekey time unit of Internet Key Exchange
	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	// IpSecEnabled
	// Enable System IPSec
	// Constraints:
	//    - required
	IpSecEnabled *bool `json:"ipSecEnabled" validate:"required"`

	OcspEnabled *bool `json:"ocspEnabled,omitempty"`

	OcspServerUri *string `json:"ocspServerUri,omitempty"`

	// PreSharedKey
	// Pre-shared key
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// RemoteId
	// The IP of IPSec receiver
	RemoteId *string `json:"remoteId,omitempty"`

	// ScgCertId
	// SCG client certification id.
	ScgCertId *string `json:"scgCertId,omitempty"`

	// SecurityGateway
	// Security gateway IP
	SecurityGateway *string `json:"securityGateway,omitempty"`

	// SubnetMask
	// Subnet Mask of security gateway
	SubnetMask *string `json:"subnetMask,omitempty"`

	TrustChainProfileId *string `json:"trustChainProfileId,omitempty"`
}

func NewUpdate() *Update {
	updateType := new(Update)
	return updateType
}

func NewUpdateWithDefaults() *Update {
	updateType := new(Update)
	return updateType
}
