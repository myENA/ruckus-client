package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGL3RoamingService struct {
	apiClient *APIClient
}

func NewWSGL3RoamingService(c *APIClient) *WSGL3RoamingService {
	s := new(WSGL3RoamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGL3RoamingService() *WSGL3RoamingService {
	return NewWSGL3RoamingService(ss.apiClient)
}

// FindProfilesTunnelL3Roaming
//
// Use this API command to retrieve L3 Roaming basic configuration.
func (s *WSGL3RoamingService) FindProfilesTunnelL3Roaming(ctx context.Context) (*WSGProfileGetL3RoamingConfig, error) {
	var (
		req  *APIRequest
		resp *WSGProfileGetL3RoamingConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindProfilesTunnelL3Roaming, true)
}

// FindProfilesTunnelL3RoamingByQueryCriteria
//
// Use this API command to retrieve L3 Roaming configuration.
func (s *WSGL3RoamingService) FindProfilesTunnelL3RoamingByQueryCriteria(ctx context.Context) (*WSGProfileGetL3RoamingConfig, error) {
	var (
		req  *APIRequest
		resp *WSGProfileGetL3RoamingConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGFindProfilesTunnelL3RoamingByQueryCriteria, true)
}

// PartialUpdateProfilesTunnelL3Roaming
//
// Use this API command to modify L3 Roaming basic configuration.
//
// Request Body:
//	 - body *WSGProfileUpdateL3RoamingConfig
func (s *WSGL3RoamingService) PartialUpdateProfilesTunnelL3Roaming(ctx context.Context, body *WSGProfileUpdateL3RoamingConfig) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateProfilesTunnelL3Roaming, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
