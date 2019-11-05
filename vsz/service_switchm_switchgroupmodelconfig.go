package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/groupmodelconfig"
)

type SwitchMSwitchGroupModelConfigService struct {
	apiClient *APIClient
}

func NewSwitchMSwitchGroupModelConfigService(c *APIClient) *SwitchMSwitchGroupModelConfigService {
	s := new(SwitchMSwitchGroupModelConfigService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchGroupModelConfigService() *SwitchMSwitchGroupModelConfigService {
	serv := new(SwitchMSwitchGroupModelConfigService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindGroupModelConfigsByQueryCriteria
//
// Use this API command to retrieve the list of group model configs.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchGroupModelConfigService) FindGroupModelConfigsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*groupmodelconfig.GroupModelConfigQueryResult, error) {
}

// UpdateGroupModelConfigsByGroupId
//
// Use this API command to add or remove the model family of a group config.
//
// Request Body:
//	 - body *groupmodelconfig.SelectedIds
//
// Path Parameters:
// - pGroupId string
//		- required
func (s *SwitchMSwitchGroupModelConfigService) UpdateGroupModelConfigsByGroupId(ctx context.Context, body *groupmodelconfig.SelectedIds, pGroupId string) (*groupmodelconfig.UpdateGroupConfigResultList, error) {
}
