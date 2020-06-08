package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDevicepolicyindomainlevelService struct {
	apiClient *VSZClient
}

func NewWSGDevicepolicyindomainlevelService(c *VSZClient) *WSGDevicepolicyindomainlevelService {
	s := new(WSGDevicepolicyindomainlevelService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDevicepolicyindomainlevelService() *WSGDevicepolicyindomainlevelService {
	return NewWSGDevicepolicyindomainlevelService(ss.apiClient)
}

// AddDevicePolicy
//
// Use this API command to create a Device Policy profile.
//
// Request Body:
//	 - body *WSGDomaindevicepolicyCreateDomainDevicePolicy
func (s *WSGDevicepolicyindomainlevelService) AddDevicePolicy(ctx context.Context, body *WSGDomaindevicepolicyCreateDomainDevicePolicy) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDevicePolicy, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDevicePolicy
//
// Use this API command to delete a list of Device Policy profile.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGDevicepolicyindomainlevelService) DeleteDevicePolicy(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDevicePolicy, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteDevicePolicyById
//
// Use this API command to delete a Device Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDevicepolicyindomainlevelService) DeleteDevicePolicyById(ctx context.Context, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDevicePolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindDevicePolicy
//
// Use this API command to retrieve list of Device Policy profiles.
//
// Optional Parameters:
// - domainId string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGDevicepolicyindomainlevelService) FindDevicePolicy(ctx context.Context, optionalParams map[string][]string) (*WSGDomaindevicepolicyProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDomaindevicepolicyProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDevicePolicy, true)
	if v, ok := optionalParams["domainId"]; ok && len(v) > 0 {
		req.SetQueryParameter("domainId", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDomaindevicepolicyProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDevicePolicyById
//
// Use this API command to retrieve a Device Policy profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDevicepolicyindomainlevelService) FindDevicePolicyById(ctx context.Context, id string) (*WSGDomaindevicepolicyProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDomaindevicepolicyProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDevicePolicyById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDomaindevicepolicyProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDevicePolicyByQueryCriteria
//
// Query Device Policy Profile with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDevicepolicyindomainlevelService) FindDevicePolicyByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindDevicePolicyByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDevicePolicyById
//
// Use this API command to update a Device Policy profile.
//
// Request Body:
//	 - body *WSGDomaindevicepolicyModifyDomainDevicePolicy
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDevicepolicyindomainlevelService) UpdateDevicePolicyById(ctx context.Context, body *WSGDomaindevicepolicyModifyDomainDevicePolicy, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDevicePolicyById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}