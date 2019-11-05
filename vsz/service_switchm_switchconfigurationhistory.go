package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/deploylog"
	"github.com/myENA/ruckus-client/vsz/types/switchm/deploylogitem"
)

type SwitchMSwitchConfigurationHistoryService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchConfigurationHistoryService(c *APIClient) *SwitchMSwitchConfigurationHistoryService {
	s := new(SwitchMSwitchConfigurationHistoryService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchConfigurationHistoryService() *SwitchMSwitchConfigurationHistoryService {
	serv := new(SwitchMSwitchConfigurationHistoryService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindConfigurationHistory
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistory(ctx context.Context) (*deploylog.ConfigurationHistoryQueryResult, error) {
}

// FindConfigurationHistoryByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*deploylog.ConfigurationHistoryQueryResult, error) {
}

// FindConfigurationHistoryDetail
//
// Use this API command to Retrieve Configuration History List.
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryDetail(ctx context.Context) (*deploylogitem.ConfigurationHistoryDetailQueryResult, error) {
}

// FindConfigurationHistoryDetailByQueryCriteria
//
// Use this API command to Query Configuration History List.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchConfigurationHistoryService) FindConfigurationHistoryDetailByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*deploylogitem.ConfigurationHistoryDetailQueryResult, error) {
}
