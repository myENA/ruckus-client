package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGVendorSpecificAttributeProfileService struct {
	apiClient *VSZClient
}

func NewWSGVendorSpecificAttributeProfileService(c *VSZClient) *WSGVendorSpecificAttributeProfileService {
	s := new(WSGVendorSpecificAttributeProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGVendorSpecificAttributeProfileService() *WSGVendorSpecificAttributeProfileService {
	return NewWSGVendorSpecificAttributeProfileService(ss.apiClient)
}

type WSGVendorSpecificAttributeProfileCreateResult struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileCreateResult() *WSGVendorSpecificAttributeProfileCreateResult {
	m := new(WSGVendorSpecificAttributeProfileCreateResult)
	return m
}

type WSGVendorSpecificAttributeProfileDeleteBulk struct {
	IdList WSGCommonIdList `json:"idList,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileDeleteBulk() *WSGVendorSpecificAttributeProfileDeleteBulk {
	m := new(WSGVendorSpecificAttributeProfileDeleteBulk)
	return m
}

type WSGVendorSpecificAttributeProfileGet struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes,omitempty"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// DomainId
	// Domain Id
	DomainId *string `json:"domainId,omitempty"`

	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// ZoneId
	// Zone Id
	ZoneId *string `json:"zoneId,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileGet() *WSGVendorSpecificAttributeProfileGet {
	m := new(WSGVendorSpecificAttributeProfileGet)
	return m
}

type WSGVendorSpecificAttributeProfileList struct {
	// FirstIndex
	// Index of the first profile returned out of the profile list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more profiles after the list that is currently displayed
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Information list of the vendor specific attribute profile
	List []*WSGVendorSpecificAttributeProfileListType `json:"list,omitempty"`

	// TotalCount
	// Total number of the profiles
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileList() *WSGVendorSpecificAttributeProfileList {
	m := new(WSGVendorSpecificAttributeProfileList)
	return m
}

type WSGVendorSpecificAttributeProfileListType struct {
	// Id
	// Identifier of the vendor specific attribute profile
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileListType() *WSGVendorSpecificAttributeProfileListType {
	m := new(WSGVendorSpecificAttributeProfileListType)
	return m
}

type WSGVendorSpecificAttributeProfilePersist struct {
	// Attributes
	// Vendor specific attribute list for Radius protocol
	// Constraints:
	//    - required
	Attributes []*WSGVendorSpecificAttributeProfileVendorSpecificAttribute `json:"attributes"`

	Description *WSGCommonDescription `json:"description,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`
}

func NewWSGVendorSpecificAttributeProfilePersist() *WSGVendorSpecificAttributeProfilePersist {
	m := new(WSGVendorSpecificAttributeProfilePersist)
	return m
}

type WSGVendorSpecificAttributeProfileQueryCriteriaResult struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGVendorSpecificAttributeProfileGet `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGVendorSpecificAttributeProfileQueryCriteriaResult() *WSGVendorSpecificAttributeProfileQueryCriteriaResult {
	m := new(WSGVendorSpecificAttributeProfileQueryCriteriaResult)
	return m
}

type WSGVendorSpecificAttributeProfileVendorSpecificAttribute struct {
	// KeyId
	// Key ID of vendor specific attribute
	// Constraints:
	//    - required
	KeyId *int `json:"keyId"`

	// SupportedRadiusProtocol
	// The radius protocol to which this given vendor specific attribute will attach
	// Constraints:
	//    - required
	//    - oneof:[ACCOUNTING,AUTHENTICATION,BOTH_AUTHENTICATION_AND_ACCOUNTING]
	SupportedRadiusProtocol *string `json:"supportedRadiusProtocol"`

	// Type
	// Type of vendor specific attribute
	// Constraints:
	//    - required
	//    - oneof:[INTEGER,STRING]
	Type *string `json:"type"`

	// Value
	// Value of vendor specific attribute
	// Constraints:
	//    - required
	Value *string `json:"value"`

	// VendorId
	// Vendor ID of vendor specific attribute
	// Constraints:
	//    - required
	VendorId *int `json:"vendorId"`
}

func NewWSGVendorSpecificAttributeProfileVendorSpecificAttribute() *WSGVendorSpecificAttributeProfileVendorSpecificAttribute {
	m := new(WSGVendorSpecificAttributeProfileVendorSpecificAttribute)
	return m
}

// AddRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Create a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) AddRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, zoneId string) (*WSGVendorSpecificAttributeProfileCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorSpecificAttributeProfileCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorSpecificAttributeProfileCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to delete a vendor specific attribute profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Use this API command to delete a list of vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfileDeleteBulk
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) DeleteRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, body *WSGVendorSpecificAttributeProfileDeleteBulk, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesVendorSpecificAttributeProfilesById
//
// Get a vendor specific attribute profile by ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, id string, zoneId string) (*WSGVendorSpecificAttributeProfileGet, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorSpecificAttributeProfileGet
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorSpecificAttributeProfileGet()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria
//
// Use this API command to retrieve a list of vendor specific attribute profile by query criteria.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGVendorSpecificAttributeProfileQueryCriteriaResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorSpecificAttributeProfileQueryCriteriaResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorSpecificAttributeProfileQueryCriteriaResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesVendorSpecificAttributeProfilesByZoneId
//
// Get a ID list of vendor specific attribute profile in this Zone.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) FindRkszonesVendorSpecificAttributeProfilesByZoneId(ctx context.Context, zoneId string) (*WSGVendorSpecificAttributeProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGVendorSpecificAttributeProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesVendorSpecificAttributeProfilesByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGVendorSpecificAttributeProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateRkszonesVendorSpecificAttributeProfilesById
//
// Use this API command to modify entire information of a vendor specific attribute profile.
//
// Request Body:
//	 - body *WSGVendorSpecificAttributeProfilePersist
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGVendorSpecificAttributeProfileService) UpdateRkszonesVendorSpecificAttributeProfilesById(ctx context.Context, body *WSGVendorSpecificAttributeProfilePersist, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesVendorSpecificAttributeProfilesById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
