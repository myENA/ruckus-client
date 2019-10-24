package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wificalling"
)

type WSGWiFiCallingPolicyService struct {
	client *Client
}

func NewWSGWiFiCallingPolicyService(client *Client) *WSGWiFiCallingPolicyService {
	s := new(WSGWiFiCallingPolicyService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWiFiCallingPolicyService() *WSGWiFiCallingPolicyService {
	serv := new(WSGWiFiCallingPolicyService)
	serv.client = ss.client
	return serv
}

func (s *WSGWiFiCallingPolicyService) FindWifiCallingByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*wificalling.WifiCallingPolicyList, error) {
}

func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicy(ctx context.Context, qDomainId string, qIndex string, qListSize string) (*wificalling.WifiCallingPolicyList, error) {
}

func (s *WSGWiFiCallingPolicyService) FindWifiCallingWifiCallingPolicyById(ctx context.Context, pId string) (*wificalling.WifiCallingPolicy, error) {
}
