package bigdog

// API Version: v9_1

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

type SwitchMWiredClientsService struct {
	apiClient *VSZClient
}

func NewSwitchMWiredClientsService(c *VSZClient) *SwitchMWiredClientsService {
	s := new(SwitchMWiredClientsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMWiredClientsService() *SwitchMWiredClientsService {
	return NewSwitchMWiredClientsService(ss.apiClient)
}

// AddSwitchClients
//
// Operation ID: addSwitchClients
//
// Use this API command to retrieve all the wired clients connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClients(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchConnectedDevicesQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchConnectedDevicesQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSwitchConnectedDevicesQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchClients, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMSwitchConnectedDevicesQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchConnectedDevicesQueryListAPIResponse), err
}

// AddSwitchClientsAp
//
// Operation ID: addSwitchClientsAp
//
// Use this API command to retrieve all the Ruckus APs connected to switch, currently managed by SmartZone.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAp(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMSwitchConnectedAPsQueryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSwitchConnectedAPsQueryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSwitchConnectedAPsQueryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchClientsAp, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMSwitchConnectedAPsQueryListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSwitchConnectedAPsQueryListAPIResponse), err
}

// AddSwitchClientsAPExport
//
// Operation ID: addSwitchClientsAPExport
//
// Download CSV of AP's discovered via LLDP
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsAPExport(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchClientsAPExport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "application/x-www-form-urlencoded")
	req.Header.Set(headerKeyAccept, "*/*")
	if b, err := json.Marshal(body); err != nil {
		return resp.(*RawAPIResponse), err
	} else if err = req.SetBody(bytes.NewBufferString((url.Values{"json": []string{string(b)}}).Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddSwitchClientsExport
//
// Operation ID: addSwitchClientsExport
//
// Download CSV of wired clients discovered via LLDP
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMWiredClientsService) AddSwitchClientsExport(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddSwitchClientsExport, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "application/x-www-form-urlencoded")
	req.Header.Set(headerKeyAccept, "*/*")
	if b, err := json.Marshal(body); err != nil {
		return resp.(*RawAPIResponse), err
	} else if err = req.SetBody(bytes.NewBufferString((url.Values{"json": []string{string(b)}}).Encode())); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
