package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGClientIsolationWhitelistService struct {
	apiClient *VSZClient
}

func NewWSGClientIsolationWhitelistService(c *VSZClient) *WSGClientIsolationWhitelistService {
	s := new(WSGClientIsolationWhitelistService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGClientIsolationWhitelistService() *WSGClientIsolationWhitelistService {
	return NewWSGClientIsolationWhitelistService(ss.apiClient)
}

// AddRkszonesClientIsolationWhitelistByZoneId
//
// Create a new ClientIsolationWhitelist.
//
// Operation ID: addRkszonesClientIsolationWhitelistByZoneId
// Operation path: /rkszones/{zoneId}/clientIsolationWhitelist
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGProfileCreateClientIsolationWhitelist
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) AddRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, body *WSGProfileCreateClientIsolationWhitelist, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddRkszonesClientIsolationWhitelistByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesClientIsolationWhitelist
//
// Use this API command to delete Bulk Client Isolation Whitelist.
//
// Operation ID: deleteRkszonesClientIsolationWhitelist
// Operation path: /rkszones/clientIsolationWhitelist
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelist(ctx context.Context, body *WSGCommonBulkDeleteRequest, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelist, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesClientIsolationWhitelistById
//
// Delete a Client Isolation Whitelist.
//
// Operation ID: deleteRkszonesClientIsolationWhitelistById
// Operation path: /rkszones/clientIsolationWhitelist/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
func (s *WSGClientIsolationWhitelistService) DeleteRkszonesClientIsolationWhitelistById(ctx context.Context, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteWSGDeleteRkszonesClientIsolationWhitelistById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesClientIsolationWhitelistById
//
// Retrieve an Client Isolation Whitelist.
//
// Operation ID: findRkszonesClientIsolationWhitelistById
// Operation path: /rkszones/{zoneId}/clientIsolationWhitelist/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGProfileClientIsolationWhitelistAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileClientIsolationWhitelistAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileClientIsolationWhitelistAPIResponse), err
}

// FindRkszonesClientIsolationWhitelistByZoneId
//
// Retrieve a list of Client Isolation Whitelist.
//
// Operation ID: findRkszonesClientIsolationWhitelistByZoneId
// Operation path: /rkszones/{zoneId}/clientIsolationWhitelist
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) FindRkszonesClientIsolationWhitelistByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGProfileClientIsolationWhitelistArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileClientIsolationWhitelistArrayAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindRkszonesClientIsolationWhitelistByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
}

// FindServicesClientIsolationWhitelistByQueryCriteria
//
// Retrieve a list of Client Isolation Whitelist.
//
// Operation ID: findServicesClientIsolationWhitelistByQueryCriteria
// Operation path: /query/services/clientIsolationWhitelist
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGClientIsolationWhitelistService) FindServicesClientIsolationWhitelistByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*WSGProfileClientIsolationWhitelistArrayAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGProfileClientIsolationWhitelistArrayAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindServicesClientIsolationWhitelistByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGProfileClientIsolationWhitelistArrayAPIResponse), err
}

// PartialUpdateRkszonesClientIsolationWhitelistById
//
// Modify a specific Client Isolation Whitelist basic.
//
// Operation ID: partialUpdateRkszonesClientIsolationWhitelistById
// Operation path: /rkszones/{zoneId}/clientIsolationWhitelist/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGProfileModifyClientIsolationWhitelist
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGClientIsolationWhitelistService) PartialUpdateRkszonesClientIsolationWhitelistById(ctx context.Context, body *WSGProfileModifyClientIsolationWhitelist, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPatch, RouteWSGPartialUpdateRkszonesClientIsolationWhitelistById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
