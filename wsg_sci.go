package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSCIService struct {
	apiClient *VSZClient
}

func NewWSGSCIService(c *VSZClient) *WSGSCIService {
	s := new(WSGSCIService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	return NewWSGSCIService(ss.apiClient)
}

// WSGSCICreateSciProfile
//
// Definition: sci_createSciProfile
type WSGSCICreateSciProfile struct {
	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser"`
}

func NewWSGSCICreateSciProfile() *WSGSCICreateSciProfile {
	m := new(WSGSCICreateSciProfile)
	return m
}

// WSGSCIDeleteSciProfile
//
// Definition: sci_deleteSciProfile
type WSGSCIDeleteSciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	Id *string `json:"id"`
}

func NewWSGSCIDeleteSciProfile() *WSGSCIDeleteSciProfile {
	m := new(WSGSCIDeleteSciProfile)
	return m
}

// WSGSCIDeleteSciProfileList
//
// Definition: sci_deleteSciProfileList
type WSGSCIDeleteSciProfileList struct {
	List []*WSGSCIDeleteSciProfile `json:"list,omitempty"`
}

func NewWSGSCIDeleteSciProfileList() *WSGSCIDeleteSciProfileList {
	m := new(WSGSCIDeleteSciProfileList)
	return m
}

// WSGSCIModifyEventCode
//
// Definition: sci_modifyEventCode
type WSGSCIModifyEventCode struct {
	// SciAcceptedEventCodes
	// Constraints:
	//    - required
	SciAcceptedEventCodes []int `json:"sciAcceptedEventCodes"`
}

func NewWSGSCIModifyEventCode() *WSGSCIModifyEventCode {
	m := new(WSGSCIModifyEventCode)
	return m
}

// WSGSCIModifySciEnabled
//
// Definition: sci_modifySciEnabled
type WSGSCIModifySciEnabled struct {
	// SciEnabled
	// Is SZ/SCI interface enabled or disabled
	// Constraints:
	//    - required
	SciEnabled *bool `json:"sciEnabled"`
}

func NewWSGSCIModifySciEnabled() *WSGSCIModifySciEnabled {
	m := new(WSGSCIModifySciEnabled)
	return m
}

// WSGSCIModifySciProfile
//
// Definition: sci_modifySciProfile
type WSGSCIModifySciProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciPassword *string `json:"sciPassword"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciProfile *string `json:"sciProfile"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerHost *string `json:"sciServerHost"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciServerPort *string `json:"sciServerPort"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciSystemId *string `json:"sciSystemId"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	// Constraints:
	//    - required
	SciUser *string `json:"sciUser"`
}

func NewWSGSCIModifySciProfile() *WSGSCIModifySciProfile {
	m := new(WSGSCIModifySciProfile)
	return m
}

// WSGSCIEventCode
//
// Definition: sci_sciEventCode
type WSGSCIEventCode struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more SCI accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSCIEventCodeListType `json:"list,omitempty"`

	// TotalCount
	// Total SCI accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSCIEventCode() *WSGSCIEventCode {
	m := new(WSGSCIEventCode)
	return m
}

// WSGSCIEventCodeListType
//
// Definition: sci_sciEventCodeListType
type WSGSCIEventCodeListType struct {
	// Code
	// SCI accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// SCI accepted event type
	Type *string `json:"type,omitempty"`
}

func NewWSGSCIEventCodeListType() *WSGSCIEventCodeListType {
	m := new(WSGSCIEventCodeListType)
	return m
}

// WSGSCIProfile
//
// Definition: sci_sciProfile
type WSGSCIProfile struct {
	// Id
	// UUID of the SCI profile for SZ/SCI interface
	Id *string `json:"id,omitempty"`

	// SciPassword
	// SCI password of the SCI profile for SZ/SCI interface
	SciPassword *string `json:"sciPassword,omitempty"`

	// SciPriority
	// Priority of the SCI profile for SZ/SCI interface
	SciPriority *int `json:"sciPriority,omitempty"`

	// SciProfile
	// Profile name of the SCI profile for SZ/SCI interface
	SciProfile *string `json:"sciProfile,omitempty"`

	// SciServerHost
	// SCI server host of the SCI profile for SZ/SCI interface
	SciServerHost *string `json:"sciServerHost,omitempty"`

	// SciServerPort
	// SCI server port of the SCI profile for SZ/SCI interface
	SciServerPort *string `json:"sciServerPort,omitempty"`

	// SciSystemId
	// SCI system UUID of the SCI profile for SZ/SCI interface
	SciSystemId *string `json:"sciSystemId,omitempty"`

	// SciUser
	// SCI user name of the SCI profile for SZ/SCI interface
	SciUser *string `json:"sciUser,omitempty"`
}

func NewWSGSCIProfile() *WSGSCIProfile {
	m := new(WSGSCIProfile)
	return m
}

// WSGSCIProfileListExtraType
//
// Definition: sci_sciProfileListExtraType
type WSGSCIProfileListExtraType struct {
	// SciEnabled
	// SCI password of the SCI profile for SZ/SCI interface
	SciEnabled *bool `json:"sciEnabled,omitempty"`
}

func NewWSGSCIProfileListExtraType() *WSGSCIProfileListExtraType {
	m := new(WSGSCIProfileListExtraType)
	return m
}

// AddSciSciEventCode
//
// Operation ID: addSciSciEventCode
//
// Use this API command to modify SciAcceptedEventCodes.
//
// Request Body:
//	 - body *WSGSCIModifyEventCode
func (s *WSGSCIService) AddSciSciEventCode(ctx context.Context, body *WSGSCIModifyEventCode, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSciSciEventCode, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddSciSciProfile
//
// Operation ID: addSciSciProfile
//
// Use this API command to create sciProfile.
//
// Request Body:
//	 - body *WSGSCICreateSciProfile
func (s *WSGSCIService) AddSciSciProfile(ctx context.Context, body *WSGSCICreateSciProfile, mutators ...RequestMutator) (*WSGCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddSciSciProfile, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteSciSciProfile
//
// Operation ID: deleteSciSciProfile
//
// Use this API command to delete sciProfile list.
//
// Request Body:
//	 - body *WSGSCIDeleteSciProfileList
func (s *WSGSCIService) DeleteSciSciProfile(ctx context.Context, body *WSGSCIDeleteSciProfileList, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSciSciProfile, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteSciSciProfileById
//
// Operation ID: deleteSciSciProfileById
//
// Use this API command to delete sciProfile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) DeleteSciSciProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteSciSciProfileById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// FindSciSciEventCode
//
// Operation ID: findSciSciEventCode
//
// Use this API command to retrieve SciAcceptedEventCodes.
func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context, mutators ...RequestMutator) (*WSGSCIEventCode, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCIEventCode
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSciSciEventCode, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSCIEventCode()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSciSciProfile
//
// Operation ID: findSciSciProfile
//
// Use this API command to retrieve sciProfile list.
func (s *WSGSCIService) FindSciSciProfile(ctx context.Context, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSciSciProfile, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindSciSciProfileById
//
// Operation ID: findSciSciProfileById
//
// Use this API command to retrieve sciProfile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGSCIProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSCIProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSciSciProfileById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSCIProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateSciSciEnabled
//
// Operation ID: partialUpdateSciSciEnabled
//
// Use this API command to modify SCI settings is enabled or not.
//
// Request Body:
//	 - body *WSGSCIModifySciEnabled
func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context, body *WSGSCIModifySciEnabled, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSciSciEnabled, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// PartialUpdateSciSciProfileById
//
// Operation ID: partialUpdateSciSciProfileById
//
// Use this API command to modify sciProfile.
//
// Request Body:
//	 - body *WSGSCIModifySciProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGSCIService) PartialUpdateSciSciProfileById(ctx context.Context, body *WSGSCIModifySciProfile, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateSciSciProfileById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
