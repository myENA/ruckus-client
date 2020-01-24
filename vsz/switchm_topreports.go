package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMTopReportsService struct {
	apiClient *APIClient
}

func NewSwitchMTopReportsService(c *APIClient) *SwitchMTopReportsService {
	s := new(SwitchMTopReportsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTopReportsService() *SwitchMTopReportsService {
	return NewSwitchMTopReportsService(ss.apiClient)
}

// AddSwitchTopByFirmware
//
// Use this API command to retrieves top N switch count based on firmware version.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByFirmware(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTopSwitchesByFirmwareQueryResultList, error) {
	var (
		req  *APIRequest
		resp *SwitchMSwitchTopSwitchesByFirmwareQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchTopByFirmware, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddSwitchTopByModel
//
// Use this API command to retrieve top N switch count based on switch model.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByModel(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMSwitchTopSwitchesByModelQueryResultList, error) {
	var (
		req  *APIRequest
		resp *SwitchMSwitchTopSwitchesByModelQueryResultList
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddSwitchTopByModel, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
