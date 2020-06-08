package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDhcpService struct {
	apiClient *VSZClient
}

func NewWSGDhcpService(c *VSZClient) *WSGDhcpService {
	s := new(WSGDhcpService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDhcpService() *WSGDhcpService {
	return NewWSGDhcpService(ss.apiClient)
}

// AddRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to create DHCP Pool.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDhcpService) AddRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGProfileCreateDhcpProfile, zoneId string) (*WSGCommonCreateResult, *APIResponseMeta, error) {
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpProfileByZoneId, true)
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

// AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId
//
// Use this API command to get the DHCP/NAT service IP assignment when selecting with "Enable on Multiple APs". In the Manually Select AP mode (the manualSelect is true), the body should contain the selected APs (include the siteAps array). Otherwise, there is no need to include the selected APs in the Auto Select AP mode (see samples).
//
// Request Body:
//	 - body *WSGCommonDoAssignIp
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDhcpService) AddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId(ctx context.Context, body *WSGCommonDoAssignIp, zoneId string) (*WSGCommonDhcpSiteConfigListRef, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddRkszonesDhcpSiteDhcpSiteConfigDoAssignIpByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to delete DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDhcpService) DeleteRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to delete multiple DHCP Pools.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDhcpService) DeleteRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, body *WSGCommonBulkDeleteRequest, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteRkszonesDhcpSiteDhcpProfileByZoneId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindDhcpDataDhcpMsgStatsByApMac
//
// Use this API command to get AP DHCP Message Statistic.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDhcpService) FindDhcpDataDhcpMsgStatsByApMac(ctx context.Context, apMac string) (*WSGDHCPMessageStatsDhcpMsgStats, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPMessageStatsDhcpMsgStats
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpMsgStatsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDHCPMessageStatsDhcpMsgStats()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDhcpDataDhcpPoolsByApMac
//
// Use this API command to get AP DHCP Pools Usage.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGDhcpService) FindDhcpDataDhcpPoolsByApMac(ctx context.Context, apMac string) (*WSGDHCPPools, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPPools
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByApMac, true)
	req.SetPathParameter("apMac", apMac)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDHCPPools()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDhcpDataDhcpPoolsByPoolIndex
//
// Use this API command to get AP DHCP Pool Usage by pool's index.
//
// Required Parameters:
// - apMac string
//		- required
// - poolIndex string
//		- required
func (s *WSGDhcpService) FindDhcpDataDhcpPoolsByPoolIndex(ctx context.Context, apMac string, poolIndex string) (*WSGDHCPPoolsDhcpPoolInfo, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDHCPPoolsDhcpPoolInfo
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDhcpDataDhcpPoolsByPoolIndex, true)
	req.SetPathParameter("apMac", apMac)
	req.SetPathParameter("poolIndex", poolIndex)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGDHCPPoolsDhcpPoolInfo()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to get DHCP Pool by pool's ID.
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDhcpService) FindRkszonesDhcpSiteDhcpProfileById(ctx context.Context, id string, zoneId string) (*WSGCommonDhcpProfileRef, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpProfileRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileById, true)
	req.SetPathParameter("id", id)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonDhcpProfileRef()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpProfileByZoneId
//
// Use this API command to get DHCP Pool list.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDhcpService) FindRkszonesDhcpSiteDhcpProfileByZoneId(ctx context.Context, zoneId string) (*WSGProfileDhcpProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpProfileByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileDhcpProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesDhcpSiteDhcpSiteConfigByZoneId
//
// Use this API command to get DHCP Configuration.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGDhcpService) FindRkszonesDhcpSiteDhcpSiteConfigByZoneId(ctx context.Context, zoneId string) (*WSGCommonDhcpSiteConfigListRef, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonDhcpSiteConfigListRef
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindRkszonesDhcpSiteDhcpSiteConfigByZoneId, true)
	req.SetPathParameter("zoneId", zoneId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonDhcpSiteConfigListRef()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindRkszonesServicesDhcpSiteConfigByQueryCriteria
//
// Use this API command to modify DHCP/NAT service configuration of Domain.
//
// Request Body:
//	 - body *WSGZoneQueryCriteria
func (s *WSGDhcpService) FindRkszonesServicesDhcpSiteConfigByQueryCriteria(ctx context.Context, body *WSGZoneQueryCriteria) (*WSGZoneDhcpSiteConfigList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGZoneDhcpSiteConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindRkszonesServicesDhcpSiteConfigByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGZoneDhcpSiteConfigList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindServicesDhcpProfileByQueryCriteria
//
// Query DHCP Profiles with specified filters.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGDhcpService) FindServicesDhcpProfileByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileDhcpProfileList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGProfileDhcpProfileList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindServicesDhcpProfileByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGProfileDhcpProfileList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateRkszonesDhcpSiteDhcpProfileById
//
// Use this API command to modify DHCP Pool by pool's ID.
//
// Request Body:
//	 - body *WSGProfileCreateDhcpProfile
//
// Required Parameters:
// - id string
//		- required
// - zoneId string
//		- required
func (s *WSGDhcpService) PartialUpdateRkszonesDhcpSiteDhcpProfileById(ctx context.Context, body *WSGProfileCreateDhcpProfile, id string, zoneId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateRkszonesDhcpSiteDhcpProfileById, true)
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