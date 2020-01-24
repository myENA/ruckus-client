package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type SwitchMVLANSettingService struct {
	apiClient *APIClient
}

func NewSwitchMVLANSettingService(c *APIClient) *SwitchMVLANSettingService {
	s := new(SwitchMVLANSettingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMVLANSettingService() *SwitchMVLANSettingService {
	return NewSwitchMVLANSettingService(ss.apiClient)
}

// AddVlans
//
// Use this API command to Create the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVlanConfigCreateVlanConfig
func (s *SwitchMVLANSettingService) AddVlans(ctx context.Context, body *SwitchMVlanConfigCreateVlanConfig) (*SwitchMCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddVlans, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteVlans
//
// Use this API command to Delete the VLAN Config by Id list.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
func (s *SwitchMVLANSettingService) DeleteVlans(ctx context.Context, body *SwitchMCommonBulkDeleteRequest) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVlans, true)
	if err = req.SetBody(body); err != nil {
		return err
	}
}

// DeleteVlansById
//
// Use this API command to Delete the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingService) DeleteVlansById(ctx context.Context, id string) error {
	var (
		req *APIRequest
		err error
	)
	if err = ctx.Err(); err != nil {
		return err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteVlansById, true)
}

// FindVlans
//
// Use this API command to Retrieve the VLAN Config List.
func (s *SwitchMVLANSettingService) FindVlans(ctx context.Context) (*SwitchMVlanConfigQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMVlanConfigQueryResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVlans, true)
}

// FindVlansById
//
// Use this API command to Retrieve the VLAN Config.
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingService) FindVlansById(ctx context.Context, id string) (*SwitchMVlanConfig, error) {
	var (
		req  *APIRequest
		resp *SwitchMVlanConfig
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindVlansById, true)
}

// FindVlansByQueryCriteria
//
// Use this API command to Retrieve the VLAN Config list.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMVLANSettingService) FindVlansByQueryCriteria(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet) (*SwitchMVlanConfigQueryResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMVlanConfigQueryResult
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
	req = NewAPIRequest(http.MethodPost, RouteSwitchMFindVlansByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateVlansById
//
// Use this API command to Update the VLAN Config.
//
// Request Body:
//	 - body *SwitchMVlanConfigUpdateVlanConfig
//
// Required Parameters:
// - id string
//		- required
func (s *SwitchMVLANSettingService) UpdateVlansById(ctx context.Context, body *SwitchMVlanConfigUpdateVlanConfig, id string) (*SwitchMVlanConfigEmptyResult, error) {
	var (
		req  *APIRequest
		resp *SwitchMVlanConfigEmptyResult
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateVlansById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
