package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGConnectivityToolsService struct {
	apiClient *VSZClient
}

func NewWSGConnectivityToolsService(c *VSZClient) *WSGConnectivityToolsService {
	s := new(WSGConnectivityToolsService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGConnectivityToolsService() *WSGConnectivityToolsService {
	return NewWSGConnectivityToolsService(ss.apiClient)
}

// AddToolSpeedflex
//
// Use this API command to start the SpeedFlex test.
//
// Operation ID: addToolSpeedflex
// Operation path: /tool/speedflex
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGToolSpeedFlex
func (s *WSGConnectivityToolsService) AddToolSpeedflex(ctx context.Context, body *WSGToolSpeedFlex, mutators ...RequestMutator) (*WSGToolTestResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGToolTestResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGAddToolSpeedflex, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGToolTestResultAPIResponse), err
}

// FindToolPing
//
// Use this API command to run the PING test on an AP.
//
// Operation ID: findToolPing
// Operation path: /tool/ping
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
// - targetIP string
//		- required
func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, apMac string, targetIP string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindToolPing, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("apMac", apMac)
	req.QueryParams.Set("targetIP", targetIP)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}

// FindToolSpeedflexByWcid
//
// Use this API command to retrieve existing SpeedFlex test results.
//
// Operation ID: findToolSpeedflexByWcid
// Operation path: /tool/speedflex/{wcid}
// Success code: 200 (OK)
//
// Required parameters:
// - wcid string
//		- required
func (s *WSGConnectivityToolsService) FindToolSpeedflexByWcid(ctx context.Context, wcid string, mutators ...RequestMutator) (*WSGToolTestResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGToolTestResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindToolSpeedflexByWcid, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("wcid", wcid)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGToolTestResultAPIResponse), err
}

// FindToolTraceRoute
//
// Use this API command to run the traceroute test on an AP.
//
// Operation ID: findToolTraceRoute
// Operation path: /tool/traceRoute
// Success code: 200 (OK)
//
// Required parameters:
// - apMac string
//		- required
// - targetIP string
//		- required
//
// Optional parameters:
// - timeoutInSec string
//		- nullable
func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, apMac string, targetIP string, optionalParams map[string][]string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindToolTraceRoute, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("apMac", apMac)
	req.QueryParams.Set("targetIP", targetIP)
	if v, ok := optionalParams["timeoutInSec"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timeoutInSec", v)
	}
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*RawAPIResponse), err
}
