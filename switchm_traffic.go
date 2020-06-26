package bigdog

// API Version: v9_0

import (
	"context"
	"net/http"
)

type SwitchMTrafficService struct {
	apiClient *VSZClient
}

func NewSwitchMTrafficService(c *VSZClient) *SwitchMTrafficService {
	s := new(SwitchMTrafficService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTrafficService() *SwitchMTrafficService {
	return NewSwitchMTrafficService(ss.apiClient)
}

// SwitchMTrafficTopPortErrorQueryResultList
//
// Definition: traffic_topPortErrorQueryResultList
type SwitchMTrafficTopPortErrorQueryResultList struct {
	// Extra
	// Extra information for top port error
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port error returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port error after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port error count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port error count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopPortErrorQueryResultList() *SwitchMTrafficTopPortErrorQueryResultList {
	m := new(SwitchMTrafficTopPortErrorQueryResultList)
	return m
}

// SwitchMTrafficTopPortTrafficUsageQueryResultList
//
// Definition: traffic_topPortTrafficUsageQueryResultList
type SwitchMTrafficTopPortTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top port traffic usage
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top port traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top port traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top port traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top port traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopPortTrafficUsageQueryResultList() *SwitchMTrafficTopPortTrafficUsageQueryResultList {
	m := new(SwitchMTrafficTopPortTrafficUsageQueryResultList)
	return m
}

// SwitchMTrafficTopSwitchPoEUtilizationQueryResultList
//
// Definition: traffic_topSwitchPoEUtilizationQueryResultList
type SwitchMTrafficTopSwitchPoEUtilizationQueryResultList struct {
	// Extra
	// Extra information for top PoE utilization
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top PoE usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top PoE usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// PoE utilization count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total PoE utilization count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopSwitchPoEUtilizationQueryResultList() *SwitchMTrafficTopSwitchPoEUtilizationQueryResultList {
	m := new(SwitchMTrafficTopSwitchPoEUtilizationQueryResultList)
	return m
}

// SwitchMTrafficTopTrafficUsageQueryResultList
//
// Definition: traffic_topTrafficUsageQueryResultList
type SwitchMTrafficTopTrafficUsageQueryResultList struct {
	// Extra
	// Extra information for top traffic usage
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first top traffic usage returned out of the complete list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more top traffic usage after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTrafficUsage `json:"list,omitempty"`

	// RawDataTotalCount
	// Top traffic usage count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of top traffic usage count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficTopTrafficUsageQueryResultList() *SwitchMTrafficTopTrafficUsageQueryResultList {
	m := new(SwitchMTrafficTopTrafficUsageQueryResultList)
	return m
}

// SwitchMTraffic
//
// Definition: traffic_traffic
type SwitchMTraffic struct {
	// Rx
	// RX traffic of the switch
	Rx interface{} `json:"rx,omitempty"`

	// Timestamp
	// Timestamp of the switch traffic
	Timestamp *string `json:"timestamp,omitempty"`

	// Total
	// Total traffic of the switch
	Total interface{} `json:"total,omitempty"`

	// Tx
	// TX traffic of the switch
	Tx interface{} `json:"tx,omitempty"`
}

func NewSwitchMTraffic() *SwitchMTraffic {
	m := new(SwitchMTraffic)
	return m
}

// SwitchMTrafficQueryResultList
//
// Definition: traffic_trafficQueryResultList
type SwitchMTrafficQueryResultList struct {
	// Extra
	// Extra information for traffic list
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first traffic list returned out of the complete traffic list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicates if there are more traffic list after the currently displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMTraffic `json:"list,omitempty"`

	// RawDataTotalCount
	// Total traffic count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total of traffic list count
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewSwitchMTrafficQueryResultList() *SwitchMTrafficQueryResultList {
	m := new(SwitchMTrafficQueryResultList)
	return m
}

// SwitchMTrafficUsage
//
// Definition: traffic_trafficUsage
type SwitchMTrafficUsage struct {
	// Id
	// Identifier of the Traffic Usage
	Id *string `json:"id,omitempty"`

	// Key
	// Interface of the Traffic Usage
	Key *string `json:"key,omitempty"`

	// Value
	// Total Tx and Rx of Traffic Usage
	Value *float64 `json:"value,omitempty"`
}

func NewSwitchMTrafficUsage() *SwitchMTrafficUsage {
	m := new(SwitchMTrafficUsage)
	return m
}

// AddTrafficTopPoeutilization
//
// Operation ID: addTrafficTopPoeutilization
//
// Use this API command retrieve the top 10 switches by the PoE utilization.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPoeutilization(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopSwitchPoEUtilizationQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopSwitchPoEUtilizationQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPoeutilization, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTrafficTopSwitchPoEUtilizationQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTopPorterror
//
// Operation ID: addTrafficTopPorterror
//
// Use this API command to get the top 10 switches by the porterror.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPorterror(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopPortErrorQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopPortErrorQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPorterror, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTrafficTopPortErrorQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTopPortusage
//
// Operation ID: addTrafficTopPortusage
//
// Use this API command to get the top 10 ports by the traffic.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopPortusage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopPortTrafficUsageQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopPortTrafficUsageQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopPortusage, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTrafficTopPortTrafficUsageQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTopUsage
//
// Operation ID: addTrafficTopUsage
//
// Use this API command to retrieve Top Swich/Port usage data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTopUsage(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficTopTrafficUsageQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficTopTrafficUsageQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTopUsage, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTrafficTopTrafficUsageQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// AddTrafficTotalTrend
//
// Operation ID: addTrafficTotalTrend
//
// Use this API command to retrieve Swich/Port trend data based on the time duration.
//
// Request Body:
//	 - body *SwitchMCommonQueryCriteriaSuperSet
func (s *SwitchMTrafficService) AddTrafficTotalTrend(ctx context.Context, body *SwitchMCommonQueryCriteriaSuperSet, mutators ...RequestMutator) (*SwitchMTrafficQueryResultList, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTrafficQueryResultList
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddTrafficTotalTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTrafficQueryResultList()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
