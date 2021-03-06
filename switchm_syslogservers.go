package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"
)

type SwitchMSyslogServersService struct {
	apiClient *VSZClient
}

func NewSwitchMSyslogServersService(c *VSZClient) *SwitchMSyslogServersService {
	s := new(SwitchMSyslogServersService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSyslogServersService() *SwitchMSyslogServersService {
	return NewSwitchMSyslogServersService(ss.apiClient)
}

// SwitchMSyslogServersCreateSyslogServerConfig
//
// Definition: syslogServers_createSyslogServerConfig
type SwitchMSyslogServersCreateSyslogServerConfig struct {
	Ip *string `json:"ip,omitempty"`

	UdpPort *int `json:"udpPort,omitempty"`
}

func NewSwitchMSyslogServersCreateSyslogServerConfig() *SwitchMSyslogServersCreateSyslogServerConfig {
	m := new(SwitchMSyslogServersCreateSyslogServerConfig)
	return m
}

// SwitchMSyslogServersSyslogServer
//
// Definition: syslogServers_syslogServer
type SwitchMSyslogServersSyslogServer struct {
	CreatedTime *int `json:"createdTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	UdpPort *int `json:"udpPort,omitempty"`

	UpdatedTime *int `json:"updatedTime,omitempty"`

	UpdaterId *string `json:"updaterId,omitempty"`

	UpdaterUsername *string `json:"updaterUsername,omitempty"`
}

func NewSwitchMSyslogServersSyslogServer() *SwitchMSyslogServersSyslogServer {
	m := new(SwitchMSyslogServersSyslogServer)
	return m
}

// SwitchMSyslogServersQueryResult
//
// Definition: syslogServers_syslogServersQueryResult
type SwitchMSyslogServersQueryResult struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Indicator of whether there are more configs after the current displayed list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSyslogServersSyslogServer `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMSyslogServersQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSyslogServersQueryResult
}

func newSwitchMSyslogServersQueryResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSyslogServersQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSyslogServersQueryResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMSyslogServersQueryResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMSyslogServersQueryResult() *SwitchMSyslogServersQueryResult {
	m := new(SwitchMSyslogServersQueryResult)
	return m
}

// AddGroupSyslogServersByGroupId
//
// Use this API command to create a new Syslog server.
//
// Operation ID: addGroupSyslogServersByGroupId
// Operation path: /group/{groupId}/syslogServers
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMSyslogServersCreateSyslogServerConfig
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMSyslogServersService) AddGroupSyslogServersByGroupId(ctx context.Context, body *SwitchMSyslogServersCreateSyslogServerConfig, groupId string, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMCommonCreateResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMAddGroupSyslogServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("groupId", groupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteGroupSyslogServersById
//
// Use this API command to delete a Syslog server.
//
// Operation ID: deleteGroupSyslogServersById
// Operation path: /group/{groupId}/syslogServers/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMSyslogServersService) DeleteGroupSyslogServersById(ctx context.Context, groupId string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodDelete, RouteSwitchMDeleteGroupSyslogServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}

// FindGroupSyslogServersByGroupId
//
// Use this API command to get Syslog servers.
//
// Operation ID: findGroupSyslogServersByGroupId
// Operation path: /group/{groupId}/syslogServers
// Success code: 200 (OK)
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMSyslogServersService) FindGroupSyslogServersByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMSyslogServersQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newSwitchMSyslogServersQueryResultAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindGroupSyslogServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*SwitchMSyslogServersQueryResultAPIResponse), err
}

// UpdateGroupSyslogServersById
//
// Use this API command to update a Syslog server.
//
// Operation ID: updateGroupSyslogServersById
// Operation path: /group/{groupId}/syslogServers/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMSyslogServersSyslogServer
//
// Required parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMSyslogServersService) UpdateGroupSyslogServersById(ctx context.Context, body *SwitchMSyslogServersSyslogServer, groupId string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteSwitchMUpdateGroupSyslogServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*EmptyAPIResponse), err
}
