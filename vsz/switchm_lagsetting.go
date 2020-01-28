package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMLAGSettingService struct {
	apiClient *APIClient
}

func NewSwitchMLAGSettingService(c *APIClient) *SwitchMLAGSettingService {
	s := new(SwitchMLAGSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMLAGSettingService() *SwitchMLAGSettingService {
	return NewSwitchMLAGSettingService(ss.apiClient)
}

// AddLagConfigs
//
// Use this API command to Create LAG Config.
//
// Request Body:
//	 - body *SwitchMLagConfigCreate
func (s *SwitchMLAGSettingService) AddLagConfigs(ctx context.Context, body *SwitchMLagConfigCreate) (SwitchMLagConfigCreateResult, error) {
	var (
		req      *APIRequest
		resp     SwitchMLagConfigCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddLagConfigs, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = MakeSwitchMLagConfigCreateResult()
	return resp, handleResponse(req, http.StatusCreated, httpResp, resp, err)
}

// DeleteLagConfigs
//
// Use this API command to Delete LAG Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMLAGSettingService) DeleteLagConfigs(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteLagConfigs, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// DeleteLagConfigsById
//
// Use this API command to Delete LAG Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) DeleteLagConfigsById(ctx context.Context, id string) error {
	var (
		req      *APIRequest
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteLagConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	return handleResponse(req, http.StatusNoContent, httpResp, nil, err)
}

// FindLagConfigs
//
// Use this API command to Retrieve all LAG Config list.
func (s *SwitchMLAGSettingService) FindLagConfigs(ctx context.Context) (*SwitchMLagConfigList, error) {
	var (
		req      *APIRequest
		resp     *SwitchMLagConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindLagConfigs, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMLagConfigList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindLagConfigsById
//
// Use this API command to Retrieve Specific LAG Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) FindLagConfigsById(ctx context.Context, id string) (*SwitchMLagConfig, error) {
	var (
		req      *APIRequest
		resp     *SwitchMLagConfig
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindLagConfigsById, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMLagConfig()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// FindLagConfigsByQueryCriteria
//
// Use this API command to Retrieve LAG Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMLAGSettingService) FindLagConfigsByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMLagConfigList, error) {
	var (
		req      *APIRequest
		resp     *SwitchMLagConfigList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindLagConfigsByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMLagConfigList()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}

// UpdateLagConfigsById
//
// Use this API command to Update LAG Config.
//
// Request Body:
//	 - body *SwitchMLagConfigModify
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMLAGSettingService) UpdateLagConfigsById(ctx context.Context, body *SwitchMLagConfigModify, id string) (*SwitchMLagConfigEmptyResult, error) {
	var (
		req      *APIRequest
		resp     *SwitchMLagConfigEmptyResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateLagConfigsById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMLagConfigEmptyResult()
	return resp, handleResponse(req, http.StatusOK, httpResp, &resp, err)
}
