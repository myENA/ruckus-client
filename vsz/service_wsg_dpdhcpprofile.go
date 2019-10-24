package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/dpprofile"
)

type WSGDPDHCPProfileService struct {
	client *Client
}

func NewWSGDPDHCPProfileService(client *Client) *WSGDPDHCPProfileService {
	s := new(WSGDPDHCPProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGDPDHCPProfileService() *WSGDPDHCPProfileService {
	serv := new(WSGDPDHCPProfileService)
	serv.client = ss.client
	return serv
}

// FindDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context) (*dpprofile.DpDhcpProfileBasicBOList, error) {
}

// FindDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileBasicBO, error) {
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, pHostId string, pId string) (*dpprofile.DpDhcpProfileHostBO, error) {
}

// FindDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileHostBOList, error) {
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBOList, error) {
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, pId string, pSpaceId string) (*dpprofile.DpDhcpProfileOptionSpaceApplyToBO, error) {
}

// FindDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, pId string) (*dpprofile.DpDhcpProfilePoolBOList, error) {
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, pId string, pPoolId string) (*dpprofile.DpDhcpProfilePoolBO, error) {
}
