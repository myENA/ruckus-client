package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGNorthboundDataStreamingService struct {
	apiClient *VSZClient
}

func NewWSGNorthboundDataStreamingService(c *VSZClient) *WSGNorthboundDataStreamingService {
	s := new(WSGNorthboundDataStreamingService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGNorthboundDataStreamingService() *WSGNorthboundDataStreamingService {
	return NewWSGNorthboundDataStreamingService(ss.apiClient)
}

// WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile
//
// Definition: northboundDataStreaming_createNorthboundDataStreamingProfile
type WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile struct {
	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerHost *string `json:"serverHost"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerPort *string `json:"serverPort"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	SystemId *string `json:"systemId"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	User *string `json:"user"`
}

func NewWSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile)
	return m
}

// WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes
//
// Definition: northboundDataStreaming_modifyNorthboundDataStreamingEventCodes
type WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes struct {
	// NorthboundDataStreamingAcceptedEventCodes
	// Constraints:
	//    - required
	NorthboundDataStreamingAcceptedEventCodes []int `json:"northboundDataStreamingAcceptedEventCodes"`
}

func NewWSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes() *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes {
	m := new(WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes)
	return m
}

// WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile
//
// Definition: northboundDataStreaming_modifyNorthboundDataStreamingProfile
type WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile struct {
	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	Password *string `json:"password"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerHost *string `json:"serverHost"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	ServerPort *string `json:"serverPort"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	SystemId *string `json:"systemId"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	// Constraints:
	//    - required
	User *string `json:"user"`
}

func NewWSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile)
	return m
}

// WSGNorthboundDataStreamingEventCodes
//
// Definition: northboundDataStreaming_northboundDataStreamingEventCodes
type WSGNorthboundDataStreamingEventCodes struct {
	// FirstIndex
	// Index of the first event code returned from the complete event code set
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates whether there are more Northbound Data Streaming accepted event codes after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGNorthboundDataStreamingEventCodesListType `json:"list,omitempty"`

	// TotalCount
	// Total Northbound Data Streaming accepted event code count
	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGNorthboundDataStreamingEventCodesAPIResponse struct {
	*RawAPIResponse
	Data *WSGNorthboundDataStreamingEventCodes
}

func newWSGNorthboundDataStreamingEventCodesAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGNorthboundDataStreamingEventCodesAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGNorthboundDataStreamingEventCodesAPIResponse) Hydrate() error {
	r.Data = new(WSGNorthboundDataStreamingEventCodes)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGNorthboundDataStreamingEventCodes() *WSGNorthboundDataStreamingEventCodes {
	m := new(WSGNorthboundDataStreamingEventCodes)
	return m
}

// WSGNorthboundDataStreamingEventCodesListType
//
// Definition: northboundDataStreaming_northboundDataStreamingEventCodesListType
type WSGNorthboundDataStreamingEventCodesListType struct {
	// Code
	// Northbound Data Streaming accepted event code
	Code *int `json:"code,omitempty"`

	// Type
	// Northbound Data Streaming accepted event type
	Type *string `json:"type,omitempty"`
}

func NewWSGNorthboundDataStreamingEventCodesListType() *WSGNorthboundDataStreamingEventCodesListType {
	m := new(WSGNorthboundDataStreamingEventCodesListType)
	return m
}

// WSGNorthboundDataStreamingProfile
//
// Definition: northboundDataStreaming_northboundDataStreamingProfile
type WSGNorthboundDataStreamingProfile struct {
	// ConnectionStatus
	// Connection status of the Northbound Data Streaming Profile
	ConnectionStatus *string `json:"connectionStatus,omitempty"`

	// DataTypes
	// Subscribed data types of the Northbound Data Streaming Profile
	// Constraints:
	//    - nullable
	DataTypes []string `json:"dataTypes,omitempty"`

	// Id
	// UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Id *string `json:"id,omitempty"`

	// Name
	// Profile name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Name *string `json:"name,omitempty"`

	// Password
	// Password of the Northbound Data Streaming profile for Northbound Data Streaming interface
	Password *string `json:"password,omitempty"`

	// ServerHost
	// Server host of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerHost *string `json:"serverHost,omitempty"`

	// ServerPort
	// Server port of the Northbound Data Streaming profile for Northbound Data Streaming interface
	ServerPort *string `json:"serverPort,omitempty"`

	// SystemId
	// System UUID of the Northbound Data Streaming profile for Northbound Data Streaming interface
	SystemId *string `json:"systemId,omitempty"`

	// User
	// User name of the Northbound Data Streaming profile for Northbound Data Streaming interface
	User *string `json:"user,omitempty"`
}

type WSGNorthboundDataStreamingProfileAPIResponse struct {
	*RawAPIResponse
	Data *WSGNorthboundDataStreamingProfile
}

func newWSGNorthboundDataStreamingProfileAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGNorthboundDataStreamingProfileAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGNorthboundDataStreamingProfileAPIResponse) Hydrate() error {
	r.Data = new(WSGNorthboundDataStreamingProfile)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGNorthboundDataStreamingProfile() *WSGNorthboundDataStreamingProfile {
	m := new(WSGNorthboundDataStreamingProfile)
	return m
}

// WSGNorthboundDataStreamingProfileList
//
// Definition: northboundDataStreaming_northboundDataStreamingProfileList
type WSGNorthboundDataStreamingProfileList struct {
	Extra *WSGNorthboundDataStreamingProfileListExtraType `json:"extra,omitempty"`

	List []*WSGNorthboundDataStreamingProfile `json:"list,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGNorthboundDataStreamingProfileList) UnmarshalJSON(b []byte) error {
	type _WSGNorthboundDataStreamingProfileList WSGNorthboundDataStreamingProfileList
	tmpType := new(_WSGNorthboundDataStreamingProfileList)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "list")
	*t = WSGNorthboundDataStreamingProfileList(*tmpType)
	return nil
}

func (t *WSGNorthboundDataStreamingProfileList) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.List != nil {
		tmp["list"] = t.List
	}
	return json.Marshal(tmp)
}

type WSGNorthboundDataStreamingProfileListAPIResponse struct {
	*RawAPIResponse
	Data *WSGNorthboundDataStreamingProfileList
}

func newWSGNorthboundDataStreamingProfileListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGNorthboundDataStreamingProfileListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGNorthboundDataStreamingProfileListAPIResponse) Hydrate() error {
	r.Data = new(WSGNorthboundDataStreamingProfileList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGNorthboundDataStreamingProfileList() *WSGNorthboundDataStreamingProfileList {
	m := new(WSGNorthboundDataStreamingProfileList)
	return m
}

// WSGNorthboundDataStreamingProfileListExtraType
//
// Definition: northboundDataStreaming_northboundDataStreamingProfileListExtraType
type WSGNorthboundDataStreamingProfileListExtraType struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled,omitempty"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled,omitempty"`

	// StreamingDomainIds
	// Domain Ids for 'streamingByDomainZoneEnabled' settings
	StreamingDomainIds []string `json:"streamingDomainIds,omitempty"`

	// StreamingZoneIds
	// Zone Ids for 'streamingByDomainZoneEnabled' settings
	StreamingZoneIds []string `json:"streamingZoneIds,omitempty"`
}

func NewWSGNorthboundDataStreamingProfileListExtraType() *WSGNorthboundDataStreamingProfileListExtraType {
	m := new(WSGNorthboundDataStreamingProfileListExtraType)
	return m
}

// WSGNorthboundDataStreamingSettings
//
// Definition: northboundDataStreaming_northboundDataStreamingSettings
type WSGNorthboundDataStreamingSettings struct {
	// NorthboundDataStreamingEnabled
	// Is Northbound Data Streaming enabled or disabled
	// Constraints:
	//    - required
	NorthboundDataStreamingEnabled *bool `json:"northboundDataStreamingEnabled"`

	// StreamingByDomainZoneEnabled
	// Is Northbound Data Streaming enabled by domain/zone settings
	// Constraints:
	//    - required
	StreamingByDomainZoneEnabled *bool `json:"streamingByDomainZoneEnabled"`

	// StreamingDomainIds
	// Domain Ids for 'streamingByDomainZoneEnabled' settings
	StreamingDomainIds []string `json:"streamingDomainIds,omitempty"`

	// StreamingZoneIds
	// Zone Ids for 'streamingByDomainZoneEnabled' settings
	StreamingZoneIds []string `json:"streamingZoneIds,omitempty"`
}

func NewWSGNorthboundDataStreamingSettings() *WSGNorthboundDataStreamingSettings {
	m := new(WSGNorthboundDataStreamingSettings)
	return m
}

// AddNorthboundDataStreamingProfile
//
// Operation ID: addNorthboundDataStreamingProfile
//
// Use this API command to create northbound Data Streaming Profile
//
// Request Body:
//	 - body *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile
func (s *WSGNorthboundDataStreamingService) AddNorthboundDataStreamingProfile(ctx context.Context, body *WSGNorthboundDataStreamingCreateNorthboundDataStreamingProfile, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddNorthboundDataStreamingProfile, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// DeleteNorthboundDataStreamingProfileById
//
// Operation ID: deleteNorthboundDataStreamingProfileById
//
// Use this API command to delete northbound Data Streaming Profile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) DeleteNorthboundDataStreamingProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteNorthboundDataStreamingProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindNorthboundDataStreamingEventCodes
//
// Operation ID: findNorthboundDataStreamingEventCodes
//
// Use this API command to retrieve NorthboundDataStreamingEventCodes.
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingEventCodes(ctx context.Context, mutators ...RequestMutator) (*WSGNorthboundDataStreamingEventCodesAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGNorthboundDataStreamingEventCodesAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGNorthboundDataStreamingEventCodesAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindNorthboundDataStreamingEventCodes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGNorthboundDataStreamingEventCodesAPIResponse), err
}

// FindNorthboundDataStreamingProfileById
//
// Operation ID: findNorthboundDataStreamingProfileById
//
// Use this API command to retrieve northbound Data Streaming Profile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGNorthboundDataStreamingProfileAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGNorthboundDataStreamingProfileAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGNorthboundDataStreamingProfileAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGNorthboundDataStreamingProfileAPIResponse), err
}

// FindNorthboundDataStreamingProfileList
//
// Operation ID: findNorthboundDataStreamingProfileList
//
// Use this API command to retrieve northbound Data Streaming Profile List
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileList(ctx context.Context, mutators ...RequestMutator) (*WSGNorthboundDataStreamingProfileListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGNorthboundDataStreamingProfileListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGNorthboundDataStreamingProfileListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindNorthboundDataStreamingProfileList, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGNorthboundDataStreamingProfileListAPIResponse), err
}

// UpdateNorthboundDataStreamingEventCodes
//
// Operation ID: updateNorthboundDataStreamingEventCodes
//
// Use this API command to modify NorthboundDataStreamingEventCodes.
//
// Request Body:
//	 - body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingEventCodes(ctx context.Context, body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingEventCodes, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingEventCodes, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateNorthboundDataStreamingProfileById
//
// Operation ID: updateNorthboundDataStreamingProfileById
//
// Use this API command to update northbound Data Streaming Profile
//
// Request Body:
//	 - body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile
//
// Required Parameters:
// - id string
//		- required
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingProfileById(ctx context.Context, body *WSGNorthboundDataStreamingModifyNorthboundDataStreamingProfile, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingProfileById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateNorthboundDataStreamingSettings
//
// Operation ID: updateNorthboundDataStreamingSettings
//
// Use this API command to modify Northbound Data Streaming Settings.
//
// Request Body:
//	 - body *WSGNorthboundDataStreamingSettings
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingSettings(ctx context.Context, body *WSGNorthboundDataStreamingSettings, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateNorthboundDataStreamingSettings, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
