package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGHistoricalClientConnectionDiagnosticService struct {
	apiClient *VSZClient
}

func NewWSGHistoricalClientConnectionDiagnosticService(c *VSZClient) *WSGHistoricalClientConnectionDiagnosticService {
	s := new(WSGHistoricalClientConnectionDiagnosticService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGHistoricalClientConnectionDiagnosticService() *WSGHistoricalClientConnectionDiagnosticService {
	return NewWSGHistoricalClientConnectionDiagnosticService(ss.apiClient)
}

// WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry
//
// Definition: historicalClientConnectionDiagnostic_clientConnectionFailureTypeCountEntry
type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry struct {
	DisplayName *string `json:"displayName,omitempty"`

	Key *string `json:"key,omitempty"`

	Value *float64 `json:"value,omitempty"`
}

func NewWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry() *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry {
	m := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry)
	return m
}

// WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList
//
// Definition: historicalClientConnectionDiagnostic_clientConnectionFailureTypeCountList
type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountEntry `json:"list,omitempty"`

	// RawDataTotalCount
	// No idea.
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse struct {
	*RawAPIResponse
	Data *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList
}

func newWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse) Hydrate() error {
	r.Data = new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList() *WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList {
	m := new(WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountList)
	return m
}

// HccdCount
//
// Operation ID: hccdCount
//
// Retrieve total count of of specified type of diagnostic entries
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - type_ string
//		- required
//		- oneof:[failureType]
func (s *WSGHistoricalClientConnectionDiagnosticService) HccdCount(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGHccdCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
	}
	req.PathParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
}

// HccdTypeCount
//
// Operation ID: hccdTypeCount
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required Parameters:
// - type_ string
//		- required
//		- oneof:[association,authentication,eap,radius,dhcp,userAuthentication]
func (s *WSGHistoricalClientConnectionDiagnosticService) HccdTypeCount(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, type_ string, mutators ...RequestMutator) (*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGHccdTypeCount, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
	}
	req.PathParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGHistoricalClientConnectionDiagnosticClientConnectionFailureTypeCountListAPIResponse), err
}
