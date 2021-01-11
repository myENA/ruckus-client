package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGGuestAccessService struct {
	apiClient *VSZClient
}

func NewWSGGuestAccessService(c *VSZClient) *WSGGuestAccessService {
	s := new(WSGGuestAccessService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGGuestAccessService() *WSGGuestAccessService {
	return NewWSGGuestAccessService(ss.apiClient)
}

// AddRkszonesPortalsGuestByZoneId
//
// Use this API command to create new guest access of a zone.
//
// Operation ID: addRkszonesPortalsGuestByZoneId
// Operation path: /rkszones/{zoneId}/portals/guest
// Success code: 201 (Created)
//
// Request body:
//	 - body *WSGPortalServiceCreateGuestAccess
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGGuestAccessService) AddRkszonesPortalsGuestByZoneId(ctx context.Context, body *WSGPortalServiceCreateGuestAccess, zoneId string, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRkszonesPortalsGuestByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteRkszonesPortalsGuestById
//
// Use this API command to delete guest access of a zone.
//
// Operation ID: deleteRkszonesPortalsGuestById
// Operation path: /rkszones/{zoneId}/portals/guest/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesPortalsGuestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesPortalsGuestRedirectById
//
// Use this API command to set redirect to the URL that user intends to visit on guest access of a zone.
//
// Operation ID: deleteRkszonesPortalsGuestRedirectById
// Operation path: /rkszones/{zoneId}/portals/guest/{id}/redirect
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestRedirectById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesPortalsGuestRedirectById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteRkszonesPortalsGuestSmsGatewayById
//
// Use this API command to disable SMS gateway on guest access of a zone.
//
// Operation ID: deleteRkszonesPortalsGuestSmsGatewayById
// Operation path: /rkszones/{zoneId}/portals/guest/{id}/smsGateway
// Success code: 204 (No Content)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGuestAccessService) DeleteRkszonesPortalsGuestSmsGatewayById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteRkszonesPortalsGuestSmsGatewayById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindRkszonesPortalsGuestById
//
// Use this API command to retrieve guest access of a zone.
//
// Operation ID: findRkszonesPortalsGuestById
// Operation path: /rkszones/{zoneId}/portals/guest/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGuestAccessService) FindRkszonesPortalsGuestById(ctx context.Context, id string, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceGuestAccessAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceGuestAccessAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceGuestAccessAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesPortalsGuestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceGuestAccessAPIResponse), err
}

// FindRkszonesPortalsGuestByZoneId
//
// Use this API command to retrieve a list of guest access of a zone.
//
// Operation ID: findRkszonesPortalsGuestByZoneId
// Operation path: /rkszones/{zoneId}/portals/guest
// Success code: 200 (OK)
//
// Required parameters:
// - zoneId string
//		- required
func (s *WSGGuestAccessService) FindRkszonesPortalsGuestByZoneId(ctx context.Context, zoneId string, mutators ...RequestMutator) (*WSGPortalServiceListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGPortalServiceListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGPortalServiceListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindRkszonesPortalsGuestByZoneId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGPortalServiceListAPIResponse), err
}

// FindServicesGuestAccessByQueryCriteria
//
// Query Guest Access Profiles with specified filters.
//
// Operation ID: findServicesGuestAccessByQueryCriteria
// Operation path: /query/services/guestAccess
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGGuestAccessService) FindServicesGuestAccessByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGFindServicesGuestAccessByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// PartialUpdateRkszonesPortalsGuestById
//
// Use this API command to modify the configuration on guest access of a zone.
//
// Operation ID: partialUpdateRkszonesPortalsGuestById
// Operation path: /rkszones/{zoneId}/portals/guest/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *WSGPortalServiceModifyGuestAccess
//
// Required parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGGuestAccessService) PartialUpdateRkszonesPortalsGuestById(ctx context.Context, body *WSGPortalServiceModifyGuestAccess, id string, zoneId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateRkszonesPortalsGuestById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("id", id)
	req.PathParams.Set("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
