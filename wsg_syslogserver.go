package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGSyslogServerService struct {
	apiClient *VSZClient
}

func NewWSGSyslogServerService(c *VSZClient) *WSGSyslogServerService {
	s := new(WSGSyslogServerService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSyslogServerService() *WSGSyslogServerService {
	return NewWSGSyslogServerService(ss.apiClient)
}

// FindSystemSyslog
//
// Operation ID: findSystemSyslog
//
// Retrieve syslog server sertting.
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGSyslogServerService) FindSystemSyslog(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGSyslogServerSettingAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSyslogServerSettingAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSyslogServerSettingAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemSyslog, true)
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
	return resp.(*WSGSyslogServerSettingAPIResponse), err
}

// PartialUpdateSystemSyslog
//
// Operation ID: partialUpdateSystemSyslog
//
// Modify syslog server setting.
//
// Request Body:
//	 - body *WSGSyslogModifySyslogSettings
func (s *WSGSyslogServerService) PartialUpdateSystemSyslog(ctx context.Context, body *WSGSyslogModifySyslogSettings, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemSyslog, true)
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

// PartialUpdateSystemSyslogPrimaryServer
//
// Operation ID: partialUpdateSystemSyslogPrimaryServer
//
// Modify Primary Server of syslog.
//
// Request Body:
//	 - body *WSGSyslogPrimaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPrimaryServer(ctx context.Context, body *WSGSyslogPrimaryServer, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemSyslogPrimaryServer, true)
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

// PartialUpdateSystemSyslogPriority
//
// Operation ID: partialUpdateSystemSyslogPriority
//
// Modify Priority of syslog.
//
// Request Body:
//	 - body *WSGSyslogPriority
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogPriority(ctx context.Context, body *WSGSyslogPriority, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemSyslogPriority, true)
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

// PartialUpdateSystemSyslogSecondaryServer
//
// Operation ID: partialUpdateSystemSyslogSecondaryServer
//
// Modify Secondary Server of syslog.
//
// Request Body:
//	 - body *WSGSyslogSecondaryServer
func (s *WSGSyslogServerService) PartialUpdateSystemSyslogSecondaryServer(ctx context.Context, body *WSGSyslogSecondaryServer, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateSystemSyslogSecondaryServer, true)
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
