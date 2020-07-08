package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGDPDHCPProfileService struct {
	apiClient *VSZClient
}

func NewWSGDPDHCPProfileService(c *VSZClient) *WSGDPDHCPProfileService {
	s := new(WSGDPDHCPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPDHCPProfileService() *WSGDPDHCPProfileService {
	return NewWSGDPDHCPProfileService(ss.apiClient)
}

// AddDpDhcpProfiles
//
// Operation ID: addDpDhcpProfiles
//
// Use this API command to create basic DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
func (s *WSGDPDHCPProfileService) AddDpDhcpProfiles(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfiles, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddDpDhcpProfilesDpDhcpProfileHostsById
//
// Operation ID: addDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to create DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileHostsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileHostBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Operation ID: addDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to create DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// AddDpDhcpProfilesDpDhcpProfilePoolsById
//
// Operation ID: addDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to create DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfilePoolsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfiles
//
// Operation ID: deleteDpDhcpProfiles
//
// Use this API command to delete DP DHCP profiles.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfiles(ctx context.Context, body *WSGDPProfileBulkDelete, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfiles, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesById
//
// Operation ID: deleteDpDhcpProfilesById
//
// Use this API command to delete DP DHCP profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesById, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to delete DP DHCP profile - host.
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("hostId", hostId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to delete DP DHCP profile - hosts.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to delete DP DHCP profile - option43 spaces.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to delete DP DHCP profile - option43 space.
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to delete DP DHCP profile - pools.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Operation ID: deleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to delete DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	req.SetHeader(headerKeyContentType, "*/*")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// FindDpDhcpProfiles
//
// Operation ID: findDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfiles, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileBasicBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesById
//
// Operation ID: findDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("hostId", hostId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileHostBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileHostsById
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileHostBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Operation ID: findDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileOptionSpaceApplyToBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsById
//
// Operation ID: findDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBOList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBOList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsById, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfilePoolBOList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Operation ID: findDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesById
//
// Operation ID: updateDpDhcpProfilesById
//
// Use this API command to modify DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileBasicBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileBasicBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesById, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileBasicBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Operation ID: updateDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to modify DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, hostId string, id string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfileHostBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfileHostBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("hostId", hostId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfileHostBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Operation ID: updateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to update DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string, spaceId string, mutators ...RequestMutator) (*RawResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("spaceId", spaceId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawResponse)
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}

// UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Operation ID: updateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to modify DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string, poolId string, mutators ...RequestMutator) (*WSGDPProfileDpDhcpProfilePoolBO, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGDPProfileDpDhcpProfilePoolBO
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	req.SetPathParameter("poolId", poolId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGDPProfileDpDhcpProfilePoolBO()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
