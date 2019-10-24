package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/traffic"
)

type SwitchMSwitchTrafficService struct {
	client *Client
}

func NewSwitchMSwitchTrafficService(client *Client) *SwitchMSwitchTrafficService {
	s := new(SwitchMSwitchTrafficService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchTrafficService() *SwitchMSwitchTrafficService {
	serv := new(SwitchMSwitchTrafficService)
	serv.client = ss.client
	return serv
}

func (s *SwitchMSwitchTrafficService) AddTrafficTopPoeutilization(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopSwitchPoEUtilizationQueryResultList, error) {
}

func (s *SwitchMSwitchTrafficService) AddTrafficTopPorterror(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopPortErrorQueryResultList, error) {
}

func (s *SwitchMSwitchTrafficService) AddTrafficTopPortusage(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopPortTrafficUsageQueryResultList, error) {
}

func (s *SwitchMSwitchTrafficService) AddTrafficTopUsage(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopTrafficUsageQueryResultList, error) {
}

func (s *SwitchMSwitchTrafficService) AddTrafficTotalTrend(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TrafficQueryResultList, error) {
}
