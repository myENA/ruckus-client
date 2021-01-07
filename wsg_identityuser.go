package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
)

type WSGIdentityUserService struct {
	apiClient *VSZClient
}

func NewWSGIdentityUserService(c *VSZClient) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGIdentityUserService() *WSGIdentityUserService {
	return NewWSGIdentityUserService(ss.apiClient)
}

// AddIdentityUserList
//
// Operation ID: addIdentityUserList
//
// Use this API command to retrieve a list of identity user.
//
// Request Body:
//	 - body *WSGIdentityQueryCriteria
func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *WSGIdentityQueryCriteria, mutators ...RequestMutator) (*WSGIdentityUserListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityUserListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityUserList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGIdentityUserListAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityUserListAPIResponse), err
}

// AddIdentityUsers
//
// Operation ID: addIdentityUsers
//
// Use this API command to create identity user.
//
// Request Body:
//	 - body *WSGIdentityCreateUser
func (s *WSGIdentityUserService) AddIdentityUsers(ctx context.Context, body *WSGIdentityCreateUser, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddIdentityUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteIdentityUsers
//
// Operation ID: deleteIdentityUsers
//
// Use this API command to delete multiple identity users.
//
// Request Body:
//	 - body *WSGIdentityDeleteBulk
func (s *WSGIdentityUserService) DeleteIdentityUsers(ctx context.Context, body *WSGIdentityDeleteBulk, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityUsers, true)
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

// DeleteIdentityUsersById
//
// Operation ID: deleteIdentityUsersById
//
// Use this API command to delete identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) DeleteIdentityUsersById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteIdentityUsersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindIdentityUsers
//
// Operation ID: findIdentityUsers
//
// Use this API command to retrieve a list of identity user.
//
// Optional Parameters:
// - createdOnFrom string
//		- nullable
// - createdOnTo string
//		- nullable
// - displayName string
//		- nullable
// - email string
//		- nullable
// - firstName string
//		- nullable
// - index string
//		- nullable
// - isDisabled string
//		- nullable
// - lastName string
//		- nullable
// - listSize string
//		- nullable
// - phone string
//		- nullable
// - timeZone string
//		- nullable
// - userName string
//		- nullable
// - userSource string
//		- nullable
// - userType string
//		- nullable
func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, optionalParams map[string][]string, mutators ...RequestMutator) (*WSGIdentityUserListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityUserListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityUsers, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	if v, ok := optionalParams["createdOnFrom"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("createdOnFrom", v)
	}
	if v, ok := optionalParams["createdOnTo"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("createdOnTo", v)
	}
	if v, ok := optionalParams["displayName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("displayName", v)
	}
	if v, ok := optionalParams["email"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("email", v)
	}
	if v, ok := optionalParams["firstName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("firstName", v)
	}
	if v, ok := optionalParams["index"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("index", v)
	}
	if v, ok := optionalParams["isDisabled"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("isDisabled", v)
	}
	if v, ok := optionalParams["lastName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("lastName", v)
	}
	if v, ok := optionalParams["listSize"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("listSize", v)
	}
	if v, ok := optionalParams["phone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("phone", v)
	}
	if v, ok := optionalParams["timeZone"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("timeZone", v)
	}
	if v, ok := optionalParams["userName"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("userName", v)
	}
	if v, ok := optionalParams["userSource"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("userSource", v)
	}
	if v, ok := optionalParams["userType"]; ok && len(v) > 0 {
		req.QueryParams.SetStrings("userType", v)
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityUserListAPIResponse), err
}

// FindIdentityUsersAaaserver
//
// Operation ID: findIdentityUsersAaaserver
//
// Use this API command to retrieve a list of aaa server.
func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityAaaServerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityAaaServerListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityAaaServerListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityUsersAaaserver, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityAaaServerListAPIResponse), err
}

// FindIdentityUsersById
//
// Operation ID: findIdentityUsersById
//
// Use this API command to retrieve identity user.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGIdentityUserConfigurationAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityUserConfigurationAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityUserConfigurationAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityUsersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityUserConfigurationAPIResponse), err
}

// FindIdentityUsersCountries
//
// Operation ID: findIdentityUsersCountries
//
// Use this API command to retrieve a list of countries.
func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityCountryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityCountryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityCountryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityUsersCountries, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityCountryListAPIResponse), err
}

// FindIdentityUsersPackages
//
// Operation ID: findIdentityUsersPackages
//
// Use this API command to retrieve a list of packages.
func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context, mutators ...RequestMutator) (*WSGIdentityPackageListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGIdentityPackageListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGIdentityPackageListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindIdentityUsersPackages, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGIdentityPackageListAPIResponse), err
}

// PartialUpdateIdentityUsersById
//
// Operation ID: partialUpdateIdentityUsersById
//
// Use this API command to modify the configuration of identity user.
//
// Request Body:
//	 - body *WSGIdentityModifyUser
//
// Required Parameters:
// - id string
//		- required
func (s *WSGIdentityUserService) PartialUpdateIdentityUsersById(ctx context.Context, body *WSGIdentityModifyUser, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPatch, RouteWSGPartialUpdateIdentityUsersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
