package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGL2AccessControlinDomainLevelService struct {
	apiClient *VSZClient
}

func NewWSGL2AccessControlinDomainLevelService(c *VSZClient) *WSGL2AccessControlinDomainLevelService {
	s := new(WSGL2AccessControlinDomainLevelService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL2AccessControlinDomainLevelService() *WSGL2AccessControlinDomainLevelService {
	return NewWSGL2AccessControlinDomainLevelService(ss.apiClient)
}

// AddL2AccessControls
//
// Use this API command to create a new L2 Access Control.
//
// Operation ID: addL2AccessControls
// Operation path: /l2AccessControls
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGL2AccessControlCreateL2AccessControl
func (s *WSGL2AccessControlinDomainLevelService) AddL2AccessControls(ctx context.Context, body *WSGL2AccessControlCreateL2AccessControl, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddL2AccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteL2AccessControls
//
// Use this API command to delete a list of L2 Access Control.
//
// Operation ID: deleteL2AccessControls
// Operation path: /l2AccessControls
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGL2AccessControlinDomainLevelService) DeleteL2AccessControls(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteL2AccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteL2AccessControlsById
//
// Use this API command to delete an L2 Access Control.
//
// Operation ID: deleteL2AccessControlsById
// Operation path: /l2AccessControls/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) DeleteL2AccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteL2AccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindL2AccessControls
//
// Use this API command to retrieve a list of L2 Access Control.
//
// Operation ID: findL2AccessControls
// Operation path: /l2AccessControls
// Success code: 200 (OK)
//
// Optional parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControls(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGL2AccessControlListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGL2AccessControlListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindL2AccessControls, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGL2AccessControlListAPIResponse), err
}

// FindL2AccessControlsById
//
// Use this API command to retrieve an L2 Access Control.
//
// Operation ID: findL2AccessControlsById
// Operation path: /l2AccessControls/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControlsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGL2AccessControlAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGL2AccessControlAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindL2AccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGL2AccessControlAPIResponse), err
}

// FindL2AccessControlsByQueryCriteria
//
// Query L2 Access Control with specified filters.
//
// Operation ID: findL2AccessControlsByQueryCriteria
// Operation path: /l2AccessControls/query
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGL2AccessControlinDomainLevelService) FindL2AccessControlsByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGL2AccessControlListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGL2AccessControlListAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindL2AccessControlsByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGL2AccessControlListAPIResponse), err
}

// UpdateL2AccessControlsById
//
// Use this API command to modify a specific L2 Access Control.
//
// Operation ID: updateL2AccessControlsById
// Operation path: /l2AccessControls/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGL2AccessControlModifyL2AccessControl
//
// Required parameters:
// - id string
//		- required
func (s *WSGL2AccessControlinDomainLevelService) UpdateL2AccessControlsById(ctx context.Context, body *WSGL2AccessControlModifyL2AccessControl, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateL2AccessControlsById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
