package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGEthernetPortProfileService struct {
	apiClient *VSZClient
}

func NewWSGEthernetPortProfileService(c *VSZClient) *WSGEthernetPortProfileService {
	s := new(WSGEthernetPortProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGEthernetPortProfileService() *WSGEthernetPortProfileService {
	return NewWSGEthernetPortProfileService(ss.apiClient)
}

// AddRkszonesProfileEthernetPortByZoneId
//
// Create a new Ethernet Port Porfile.
//
// Request Body:
//	 - body *WSGEthernetPortCreateEthernetPortProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) AddRkszonesProfileEthernetPortByZoneId(ctx context.Context, body *WSGEthernetPortCreateEthernetPortProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesProfileEthernetPortByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesProfileEthernetPortById
//
// Delete Ethernet Port Porfile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) DeleteRkszonesProfileEthernetPortById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesProfileEthernetPortById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindRkszonesProfileEthernetPortById
//
// Retrieve a Ethernet Port Porfile.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortById(ctx context.Context, id string, zoneId string) (*WSGEthernetPortProfile, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGEthernetPortProfile
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesProfileEthernetPortById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEthernetPortProfile()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesProfileEthernetPortByZoneId
//
// Retrieve a list of Ethernet Port Porfiles within a zone.
//
// Required Parameters:
// - zoneId string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGEthernetPortProfileService) FindRkszonesProfileEthernetPortByZoneId(ctx context.Context, zoneId string, optionalParams map[string][]string) (*WSGEthernetPortProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGEthernetPortProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesProfileEthernetPortByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.SetQueryParameter("index", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.SetQueryParameter("listSize", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGEthernetPortProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesEthernetPortProfileByQueryCriteria
//
// Query Ethernet Port Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGEthernetPortProfileService) FindServicesEthernetPortProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (interface{}, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesEthernetPortProfileByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesProfileEthernetPortById
//
// Modify a specific Ethernet Port Porfile.
//
// Request Body:
//	 - body *WSGEthernetPortModifyEthernetPortProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) PartialUpdateRkszonesProfileEthernetPortById(ctx context.Context, body *WSGEthernetPortModifyEthernetPortProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesProfileEthernetPortById, true)
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

// UpdateRkszonesProfileEthernetPortById
//
// Modify a specific Ethernet Port Porfile.
//
// Request Body:
//	 - body *WSGEthernetPortModifyEthernetPortProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGEthernetPortProfileService) UpdateRkszonesProfileEthernetPortById(ctx context.Context, body *WSGEthernetPortModifyEthernetPortProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateRkszonesProfileEthernetPortById, true)
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
