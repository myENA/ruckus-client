package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portalservice"
)

type WSGHotspotServiceService struct {
    client *Client
}

func NewWSGHotspotServiceService (client *Client) *WSGHotspotServiceService {
    s := new(WSGHotspotServiceService)
    s.client = client
    return s
}

func (ss *WSGService) WSGHotspotServiceService () *WSGHotspotServiceService {
    serv := new(WSGHotspotServiceService)
    serv.client = ss.client
    return serv
}

