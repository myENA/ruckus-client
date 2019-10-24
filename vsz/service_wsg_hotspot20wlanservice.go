package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspot20WLANServiceService struct {
	client *Client
}

func NewWSGHotspot20WLANServiceService(client *Client) *WSGHotspot20WLANServiceService {
	s := new(WSGHotspot20WLANServiceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGHotspot20WLANServiceService() *WSGHotspot20WLANServiceService {
	serv := new(WSGHotspot20WLANServiceService)
	serv.client = ss.client
	return serv
}

func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sById(ctx context.Context, pId string, pZoneId string) (*portalservice.Hotspot20WlanProfile, error) {
}

func (s *WSGHotspot20WLANServiceService) FindRkszonesHs20sByZoneId(ctx context.Context, pZoneId string) (*portalservice.PortalServiceList, error) {
}
