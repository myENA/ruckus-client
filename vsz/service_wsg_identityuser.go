package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/identity"
)

type WSGIdentityUserService struct {
	client *Client
}

func NewWSGIdentityUserService(client *Client) *WSGIdentityUserService {
	s := new(WSGIdentityUserService)
	s.client = client
	return s
}

func (ss *WSGService) WSGIdentityUserService() *WSGIdentityUserService {
	serv := new(WSGIdentityUserService)
	serv.client = ss.client
	return serv
}

func (s *WSGIdentityUserService) AddIdentityUserList(ctx context.Context, body *identity.QueryCriteria) (*identity.UserList, error) {
}

func (s *WSGIdentityUserService) FindIdentityUsers(ctx context.Context, qCreatedOnFrom string, qCreatedOnTo string, qDisplayName string, qEmail string, qFirstName string, qIndex string, qIsDisabled string, qLastName string, qListSize string, qPhone string, qTimeZone string, qUserName string, qUserSource string, qUserType string) (*identity.UserList, error) {
}

func (s *WSGIdentityUserService) FindIdentityUsersAaaserver(ctx context.Context) (*identity.AaaServerList, error) {
}

func (s *WSGIdentityUserService) FindIdentityUsersById(ctx context.Context, pId string) (*identity.UserConfiguration, error) {
}

func (s *WSGIdentityUserService) FindIdentityUsersCountries(ctx context.Context) (*identity.CountryList, error) {
}

func (s *WSGIdentityUserService) FindIdentityUsersPackages(ctx context.Context) (*identity.PackageList, error) {
}
