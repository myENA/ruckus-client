package bigdog

// API Version: v9_1

import (
	"context"
	"net/http"
	"time"
)

type WSGResourceMonitoringService struct {
	apiClient *VSZClient
}

func NewWSGResourceMonitoringService(c *VSZClient) *WSGResourceMonitoringService {
	s := new(WSGResourceMonitoringService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGResourceMonitoringService() *WSGResourceMonitoringService {
	return NewWSGResourceMonitoringService(ss.apiClient)
}

// FindResourceMonitoringSummaryByQueryCriteria
//
// Retrieve monitoring status of resource
//
// Operation ID: findResourceMonitoringSummaryByQueryCriteria
// Operation path: /monitor/{resource}
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGCommonQueryCriteriaSuperSet
//
// Required parameters:
// - resource string
//		- required
//		- oneof:[client,ap]
func (s *WSGResourceMonitoringService) FindResourceMonitoringSummaryByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet, resource string, mutators ...RequestMutator) (*WSGCommonMonitoringSummaryAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		execDur  time.Duration
		resp     APIResponse
		err      error

		respFn = newWSGCommonMonitoringSummaryAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteWSGFindResourceMonitoringSummaryByQueryCriteria, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	req.SetBody(body)
	req.PathParams.Set("resource", resource)
	httpResp, execDur, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, execDur, respFn, s.apiClient.autoHydrate, s.apiClient.ev, err)
	return resp.(*WSGCommonMonitoringSummaryAPIResponse), err
}
