package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/accountsecurityprofile"
)

type WSGAccountSecurityService struct {
	client *Client
}

func NewWSGAccountSecurityService(client *Client) *WSGAccountSecurityService {
	s := new(WSGAccountSecurityService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAccountSecurityService() *WSGAccountSecurityService {
	serv := new(WSGAccountSecurityService)
	serv.client = ss.client
	return serv
}

func (s *WSGAccountSecurityService) FindAccountSecurity(ctx context.Context) (*accountsecurityprofile.ProfileListResult, error) {
}

func (s *WSGAccountSecurityService) FindAccountSecurityById(ctx context.Context, body *accountsecurityprofile.GetById, pId string) (*accountsecurityprofile.GetByIdResult, error) {
}
