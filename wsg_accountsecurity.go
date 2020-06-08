package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGAccountsecurityService struct {
	apiClient *VSZClient
}

func NewWSGAccountsecurityService(c *VSZClient) *WSGAccountsecurityService {
	s := new(WSGAccountsecurityService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccountsecurityService() *WSGAccountsecurityService {
	return NewWSGAccountsecurityService(ss.apiClient)
}

// AddAccountSecurity
//
// Use this API command to create the account security proile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileCreate
func (s *WSGAccountsecurityService) AddAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileCreate) (*WSGCommonCreateResultIdName, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResultIdName
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddAccountSecurity, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResultIdName()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteAccountSecurity
//
// Use this API command to selete the account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDeleteList
func (s *WSGAccountsecurityService) DeleteAccountSecurity(ctx context.Context, body *WSGAccountSecurityProfileDeleteList) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAccountSecurity, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// DeleteAccountSecurityById
//
// Use this API command to delete the account security profile by id.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountsecurityService) DeleteAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileDelete, id string) (*WSGCommonCreateResultIdName, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResultIdName
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResultIdName()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAccountSecurity
//
// Use this API command to get account security profiles.
func (s *WSGAccountsecurityService) FindAccountSecurity(ctx context.Context) (*WSGAccountSecurityProfileProfileListResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAccountSecurityProfileProfileListResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAccountSecurity, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAccountSecurityProfileProfileListResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindAccountSecurityById
//
// Use this API command to retrieve the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileGetById
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountsecurityService) FindAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileGetById, id string) (*WSGAccountSecurityProfileGetByIdResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAccountSecurityProfileGetByIdResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAccountSecurityProfileGetByIdResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// PartialUpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountsecurityService) PartialUpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, id string) (*WSGCommonCreateResultIdName, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResultIdName
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResultIdName()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateAccountSecurityById
//
// Use this API command to modify the specific account security profile.
//
// Request Body:
//	 - body *WSGAccountSecurityProfileUpdate
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAccountsecurityService) UpdateAccountSecurityById(ctx context.Context, body *WSGAccountSecurityProfileUpdate, id string) (*WSGCommonCreateResultIdName, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGCommonCreateResultIdName
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateAccountSecurityById, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGCommonCreateResultIdName()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}