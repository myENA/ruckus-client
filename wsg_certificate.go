package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGCertificateService struct {
	apiClient *VSZClient
}

func NewWSGCertificateService(c *VSZClient) *WSGCertificateService {
	s := new(WSGCertificateService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGCertificateService() *WSGCertificateService {
	return NewWSGCertificateService(ss.apiClient)
}

// WSGCertificate
//
// Definition: certificate_certificate
type WSGCertificate struct {
	CertificasSigningRequest *WSGCommonGenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// Data of the certificate
	Data *string `json:"data,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the certificate
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Passphrase
	// Key passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	// PrivateKeyData
	// Private key data of the certificate
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// PublicKey
	// Public key data of the certificate
	PublicKey *string `json:"publicKey,omitempty"`

	// RootData
	// Root data of the certificate
	RootData *string `json:"rootData,omitempty"`
}

type WSGCertificateAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificate
}

func newWSGCertificateAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificate)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificate() *WSGCertificate {
	m := new(WSGCertificate)
	return m
}

// WSGCertificateList
//
// Definition: certificate_certificateList
type WSGCertificateList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateListAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateList
}

func newWSGCertificateListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateListAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateList() *WSGCertificateList {
	m := new(WSGCertificateList)
	return m
}

// WSGCertificateListType
//
// Definition: certificate_certificateListType
type WSGCertificateListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificate
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGCertificateListType() *WSGCertificateListType {
	m := new(WSGCertificateListType)
	return m
}

// WSGCertificatesSigningRequest
//
// Definition: certificate_certificatesSigningRequest
type WSGCertificatesSigningRequest struct {
	// City
	// City of the certificates signing request
	City *string `json:"city,omitempty"`

	// CommonName
	// Common name of the certificates signing request
	CommonName *string `json:"commonName,omitempty"`

	// CountryCode
	// Country code of the certificates signing request
	CountryCode *string `json:"countryCode,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Email
	// Email of the certificates signing request
	Email *string `json:"email,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Organization
	// Organization of the certificates signing request
	Organization *string `json:"organization,omitempty"`

	// OrganizationUnit
	// Organization unit of the certificates signing request
	OrganizationUnit *string `json:"organizationUnit,omitempty"`

	// State
	// State of the certificates signing request
	State *string `json:"state,omitempty"`
}

type WSGCertificatesSigningRequestAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificatesSigningRequest
}

func newWSGCertificatesSigningRequestAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificatesSigningRequestAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificatesSigningRequestAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificatesSigningRequest)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificatesSigningRequest() *WSGCertificatesSigningRequest {
	m := new(WSGCertificatesSigningRequest)
	return m
}

// WSGCertificateCertSetting
//
// Definition: certificate_certSetting
type WSGCertificateCertSetting struct {
	// ServiceCertificates
	// Certificate Setting of the service
	ServiceCertificates []*WSGCertificateServiceCertificate `json:"serviceCertificates,omitempty"`
}

type WSGCertificateCertSettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateCertSetting
}

func newWSGCertificateCertSettingAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateCertSettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateCertSettingAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateCertSetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateCertSetting() *WSGCertificateCertSetting {
	m := new(WSGCertificateCertSetting)
	return m
}

// WSGCertificateClientCert
//
// Definition: certificate_clientCert
type WSGCertificateClientCert struct {
	// Data
	// Data of the client certificate
	Data *string `json:"data,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the client certificate
	Information *string `json:"information,omitempty"`

	// IntermediateData
	// Intermediate data of the client certificate
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// PrivateKeyData
	// Private key data of the client certificate
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// PublicKey
	// Public key data of the client certificate
	PublicKey *string `json:"publicKey,omitempty"`

	// RootData
	// Root data of the client certificate
	RootData *string `json:"rootData,omitempty"`
}

type WSGCertificateClientCertAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateClientCert
}

func newWSGCertificateClientCertAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateClientCertAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateClientCertAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateClientCert)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateClientCert() *WSGCertificateClientCert {
	m := new(WSGCertificateClientCert)
	return m
}

// WSGCertificateClientCertList
//
// Definition: certificate_clientCertList
type WSGCertificateClientCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateClientCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateClientCertListAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateClientCertList
}

func newWSGCertificateClientCertListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateClientCertListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateClientCertListAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateClientCertList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateClientCertList() *WSGCertificateClientCertList {
	m := new(WSGCertificateClientCertList)
	return m
}

// WSGCertificateClientCertListType
//
// Definition: certificate_clientCertListType
type WSGCertificateClientCertListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the client certificate
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGCertificateClientCertListType() *WSGCertificateClientCertListType {
	m := new(WSGCertificateClientCertListType)
	return m
}

// WSGCertificateCreateCert
//
// Definition: certificate_createCert
type WSGCertificateCreateCert struct {
	CertificasSigningRequest *WSGCommonGenericRef `json:"certificasSigningRequest,omitempty"`

	// Data
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	// Constraints:
	//    - required
	Data *string `json:"data"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Passphrase
	// Key passphrase
	Passphrase *string `json:"passphrase,omitempty"`

	// PrivateKeyData
	// Incompatible with certificasSigningRequest. Must choose one to enter. The value must be in PEM format which is a Base64 encoded DER certificate.
	PrivateKeyData *string `json:"privateKeyData,omitempty"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER certificate.
	RootData *string `json:"rootData,omitempty"`
}

func NewWSGCertificateCreateCert() *WSGCertificateCreateCert {
	m := new(WSGCertificateCreateCert)
	return m
}

// WSGCertificateCreateClientCert
//
// Definition: certificate_createClientCert
type WSGCertificateCreateClientCert struct {
	// Data
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	Data *string `json:"data"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// IntermediateData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - nullable
	IntermediateData []string `json:"intermediateData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// PrivateKeyData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	// Constraints:
	//    - required
	PrivateKeyData *string `json:"privateKeyData"`

	// RootData
	// The value must be in PEM format which is a Base64 encoded DER client certificate.
	RootData *string `json:"rootData,omitempty"`
}

func NewWSGCertificateCreateClientCert() *WSGCertificateCreateClientCert {
	m := new(WSGCertificateCreateClientCert)
	return m
}

// WSGCertificateCreateCSR
//
// Definition: certificate_createCSR
type WSGCertificateCreateCSR struct {
	// City
	// City of the certificates signing request
	// Constraints:
	//    - required
	//    - max:128
	City *string `json:"city"`

	// CommonName
	// Constraints:
	//    - required
	CommonName *WSGCommonFQDN `json:"commonName"`

	// CountryCode
	// Country code of the certificates signing request
	// Constraints:
	//    - required
	//    - max:2
	CountryCode *string `json:"countryCode"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Email
	// Constraints:
	//    - required
	Email *WSGCommonEmail `json:"email"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// Organization
	// Organization of the certificates signing request
	// Constraints:
	//    - required
	//    - max:64
	Organization *string `json:"organization"`

	// OrganizationUnit
	// Organization unit of the certificates signing request
	// Constraints:
	//    - max:64
	OrganizationUnit *string `json:"organizationUnit,omitempty"`

	// State
	// State of the certificates signing request
	// Constraints:
	//    - required
	//    - max:128
	State *string `json:"state"`
}

func NewWSGCertificateCreateCSR() *WSGCertificateCreateCSR {
	m := new(WSGCertificateCreateCSR)
	return m
}

// WSGCertificateCreateTrustedCAChain
//
// Definition: certificate_createTrustedCAChain
type WSGCertificateCreateTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	// Constraints:
	//    - required
	RootCertData *string `json:"rootCertData"`
}

func NewWSGCertificateCreateTrustedCAChain() *WSGCertificateCreateTrustedCAChain {
	m := new(WSGCertificateCreateTrustedCAChain)
	return m
}

// WSGCertificateCsrList
//
// Definition: certificate_csrList
type WSGCertificateCsrList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateCsrListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateCsrListAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateCsrList
}

func newWSGCertificateCsrListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateCsrListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateCsrListAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateCsrList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateCsrList() *WSGCertificateCsrList {
	m := new(WSGCertificateCsrList)
	return m
}

// WSGCertificateCsrListType
//
// Definition: certificate_csrListType
type WSGCertificateCsrListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the certificates signing request
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGCertificateCsrListType() *WSGCertificateCsrListType {
	m := new(WSGCertificateCsrListType)
	return m
}

// WSGCertificateDeleteBulk
//
// Definition: certificate_deleteBulk
type WSGCertificateDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGCertificateDeleteBulk() *WSGCertificateDeleteBulk {
	m := new(WSGCertificateDeleteBulk)
	return m
}

// WSGCertificateModifyTrustedCAChain
//
// Definition: certificate_modifyTrustedCAChain
type WSGCertificateModifyTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewWSGCertificateModifyTrustedCAChain() *WSGCertificateModifyTrustedCAChain {
	m := new(WSGCertificateModifyTrustedCAChain)
	return m
}

// WSGCertificateServiceCertificate
//
// Definition: certificate_serviceCertificate
type WSGCertificateServiceCertificate struct {
	Certificate *WSGCommonGenericRef `json:"certificate,omitempty"`

	// Service
	// the service
	// Constraints:
	//    - required
	//    - oneof:[MANAGEMENT_WEB,AP_PORTAL,HOTSPOT,COMMUNICATOR]
	Service *string `json:"service"`
}

func NewWSGCertificateServiceCertificate() *WSGCertificateServiceCertificate {
	m := new(WSGCertificateServiceCertificate)
	return m
}

// WSGCertificateServiceCertificates
//
// Definition: certificate_serviceCertificates
type WSGCertificateServiceCertificates []*WSGCertificateServiceCertificate

func MakeWSGCertificateServiceCertificates() WSGCertificateServiceCertificates {
	m := make(WSGCertificateServiceCertificates, 0)
	return m
}

// WSGCertificateTrustedCAChain
//
// Definition: certificate_trustedCAChain
type WSGCertificateTrustedCAChain struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificates
	Id *string `json:"id,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

type WSGCertificateTrustedCAChainAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateTrustedCAChain
}

func newWSGCertificateTrustedCAChainAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateTrustedCAChainAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateTrustedCAChainAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateTrustedCAChain)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateTrustedCAChain() *WSGCertificateTrustedCAChain {
	m := new(WSGCertificateTrustedCAChain)
	return m
}

// WSGCertificateTrustedCAChainCertList
//
// Definition: certificate_trustedCAChainCertList
type WSGCertificateTrustedCAChainCertList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGCertificateTrustedCAChainCertListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGCertificateTrustedCAChainCertListAPIResponse struct {
	*RawAPIResponse
	Data *WSGCertificateTrustedCAChainCertList
}

func newWSGCertificateTrustedCAChainCertListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGCertificateTrustedCAChainCertListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGCertificateTrustedCAChainCertListAPIResponse) Hydrate() error {
	r.Data = new(WSGCertificateTrustedCAChainCertList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGCertificateTrustedCAChainCertList() *WSGCertificateTrustedCAChainCertList {
	m := new(WSGCertificateTrustedCAChainCertList)
	return m
}

// WSGCertificateTrustedCAChainCertListType
//
// Definition: certificate_trustedCAChainCertListType
type WSGCertificateTrustedCAChainCertListType struct {
	Description *WSGCommonDescription `json:"description,omitempty"`

	// Id
	// Identifier of the trusted CA chain certificate
	Id *string `json:"id,omitempty"`

	// Information
	// Information of the certificates
	Information *string `json:"information,omitempty"`

	// InterCertData
	// Intermediate data of the trusted CA chain certificates
	// Constraints:
	//    - nullable
	InterCertData []string `json:"interCertData,omitempty"`

	ModifiedDateTime *WSGCommonNormalNameAllowBlank `json:"modifiedDateTime,omitempty"`

	ModifierUsername *WSGCommonNormalNameAllowBlank `json:"modifierUsername,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// RootCertData
	// Root data of the trusted CA chain certificates
	RootCertData *string `json:"rootCertData,omitempty"`
}

func NewWSGCertificateTrustedCAChainCertListType() *WSGCertificateTrustedCAChainCertListType {
	m := new(WSGCertificateTrustedCAChainCertListType)
	return m
}

// AddCertstoreCertificate
//
// Operation ID: addCertstoreCertificate
//
// Use this API command to create an installed certificate.
//
// Request Body:
//	 - body *WSGCertificateCreateCert
func (s *WSGCertificateService) AddCertstoreCertificate(ctx context.Context, body *WSGCertificateCreateCert, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddCertstoreCertificate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddCertstoreClientCert
//
// Operation ID: addCertstoreClientCert
//
// Use this API command to create a client certificate.
//
// Request Body:
//	 - body *WSGCertificateCreateClientCert
func (s *WSGCertificateService) AddCertstoreClientCert(ctx context.Context, body *WSGCertificateCreateClientCert, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddCertstoreClientCert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddCertstoreCsr
//
// Operation ID: addCertstoreCsr
//
// Use this API command to create a certificates signing request.
//
// Request Body:
//	 - body *WSGCertificateCreateCSR
func (s *WSGCertificateService) AddCertstoreCsr(ctx context.Context, body *WSGCertificateCreateCSR, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddCertstoreCsr, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddCertstoreTrustedCAChainCert
//
// Operation ID: addCertstoreTrustedCAChainCert
//
// Use this API command to create trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateCreateTrustedCAChain
func (s *WSGCertificateService) AddCertstoreTrustedCAChainCert(ctx context.Context, body *WSGCertificateCreateTrustedCAChain, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddCertstoreTrustedCAChainCert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteCertstoreCertificateById
//
// Operation ID: deleteCertstoreCertificateById
//
// Use this API command to delete an installed certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCertificateById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteCertstoreCertificateById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteCertstoreClientCertById
//
// Operation ID: deleteCertstoreClientCertById
//
// Use this API command to delete a client certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreClientCertById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteCertstoreClientCertById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteCertstoreCsrById
//
// Operation ID: deleteCertstoreCsrById
//
// Use this API command to delete a certificates signing request.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreCsrById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteCertstoreCsrById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteCertstoreTrustedCAChainCert
//
// Operation ID: deleteCertstoreTrustedCAChainCert
//
// Use this API command to delete bulk trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateDeleteBulk
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCert(ctx context.Context, body *WSGCertificateDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteCertstoreTrustedCAChainCert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteCertstoreTrustedCAChainCertById
//
// Operation ID: deleteCertstoreTrustedCAChainCertById
//
// Use this API command to delete a trusted CA chain certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) DeleteCertstoreTrustedCAChainCertById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteCertstoreTrustedCAChainCertById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindCertstoreCertificate
//
// Operation ID: findCertstoreCertificate
//
// Use this API command to retrieve list of installed certificates.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreCertificate(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCertificateListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreCertificate, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateListAPIResponse), err
}

// FindCertstoreCertificateById
//
// Operation ID: findCertstoreCertificateById
//
// Use this API command to retrieve an installed certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreCertificateById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGCertificateAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreCertificateById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateAPIResponse), err
}

// FindCertstoreClientCert
//
// Operation ID: findCertstoreClientCert
//
// Use this API command to retrieve list of client certificates.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreClientCert(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCertificateClientCertListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateClientCertListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateClientCertListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreClientCert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateClientCertListAPIResponse), err
}

// FindCertstoreClientCertById
//
// Operation ID: findCertstoreClientCertById
//
// Use this API command to retrieve a client certificate.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreClientCertById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGCertificateClientCertAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateClientCertAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateClientCertAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreClientCertById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateClientCertAPIResponse), err
}

// FindCertstoreCsr
//
// Operation ID: findCertstoreCsr
//
// Use this API command to retrieve list of certificates signing request.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreCsr(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCertificateCsrListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateCsrListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateCsrListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreCsr, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateCsrListAPIResponse), err
}

// FindCertstoreCsrById
//
// Operation ID: findCertstoreCsrById
//
// Use this API command to retrieve a certificates signing request.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreCsrById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGCertificatesSigningRequestAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificatesSigningRequestAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificatesSigningRequestAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreCsrById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificatesSigningRequestAPIResponse), err
}

// FindCertstoreSetting
//
// Operation ID: findCertstoreSetting
//
// Use this API command to retrieve certificate setting.
func (s *WSGCertificateService) FindCertstoreSetting(ctx context.Context, mutators ...RequestMutator) (*WSGCertificateCertSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateCertSettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateCertSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateCertSettingAPIResponse), err
}

// FindCertstoreTrustedCAChainCert
//
// Operation ID: findCertstoreTrustedCAChainCert
//
// Use this API command to retrieve list of installed trusted CA chain certificates.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCert(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGCertificateTrustedCAChainCertListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateTrustedCAChainCertListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateTrustedCAChainCertListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreTrustedCAChainCert, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateTrustedCAChainCertListAPIResponse), err
}

// FindCertstoreTrustedCAChainCertById
//
// Operation ID: findCertstoreTrustedCAChainCertById
//
// Use this API command to retrieve an installed trusted CA chain certificates.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) FindCertstoreTrustedCAChainCertById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGCertificateTrustedCAChainAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCertificateTrustedCAChainAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCertificateTrustedCAChainAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindCertstoreTrustedCAChainCertById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCertificateTrustedCAChainAPIResponse), err
}

// PartialUpdateCertstoreSetting
//
// Operation ID: partialUpdateCertstoreSetting
//
// Use this API command to Modify the Certificate Setting.
//
// Request Body:
//	 - body *WSGCertificateCertSetting
func (s *WSGCertificateService) PartialUpdateCertstoreSetting(ctx context.Context, body *WSGCertificateCertSetting, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateCertstoreSetting, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateCertstoreSettingServiceCertificates
//
// Operation ID: partialUpdateCertstoreSettingServiceCertificates
//
// Use this API command to Modify serviceCertificates of the Certificate Setting.
//
// Request Body:
//	 - body WSGCertificateServiceCertificates
func (s *WSGCertificateService) PartialUpdateCertstoreSettingServiceCertificates(ctx context.Context, body WSGCertificateServiceCertificates, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateCertstoreSettingServiceCertificates, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateCertstoreTrustedCAChainCertById
//
// Operation ID: partialUpdateCertstoreTrustedCAChainCertById
//
// Use this API command to patch a trusted CA chain certificates.
//
// Request Body:
//	 - body *WSGCertificateModifyTrustedCAChain
//
// Required Parameters:
// - id string
//		- required
func (s *WSGCertificateService) PartialUpdateCertstoreTrustedCAChainCertById(ctx context.Context, body *WSGCertificateModifyTrustedCAChain, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateCertstoreTrustedCAChainCertById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
