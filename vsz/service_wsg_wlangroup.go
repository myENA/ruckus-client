package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlangroup"
)

type WSGWLANGroupService struct {
	apiClient *APIClient
}

func NewWSGWLANGroupService(c *APIClient) *WSGWLANGroupService {
	s := new(WSGWLANGroupService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGWLANGroupService() *WSGWLANGroupService {
	serv := new(WSGWLANGroupService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddRkszonesWlangroupsByZoneId
//
// Use this API command to create a new WLAN group.
//
// Request Body:
//	 - body *wlangroup.CreateWlanGroup
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) AddRkszonesWlangroupsByZoneId(ctx context.Context, body *wlangroup.CreateWlanGroup, pZoneId string) (*common.CreateResult, error) {
}

// AddRkszonesWlangroupsMembersById
//
// Use this API command to add a member to a WLAN group.
//
// Request Body:
//	 - body *wlangroup.WlanMember
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) AddRkszonesWlangroupsMembersById(ctx context.Context, body *wlangroup.WlanMember, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsById
//
// Use this API command to delete a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsMembersByMemberId
//
// Use this API command to remove a member from a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsMembersNasIdByMemberId
//
// Use this API command to disable a member NAS-ID override of a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersNasIdByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

// DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId
//
// Use this API command to disable a member VLAN override of a WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) DeleteRkszonesWlangroupsMembersVlanOverrideByMemberId(ctx context.Context, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}

// FindRkszonesWlangroupsById
//
// Use this API command to retrieve the WLAN group.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) FindRkszonesWlangroupsById(ctx context.Context, pId string, pZoneId string) (*wlangroup.WlanGroup, error) {
}

// FindRkszonesWlangroupsByZoneId
//
// Use this API command to retrieve the list of WLAN groups within a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGWLANGroupService) FindRkszonesWlangroupsByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlangroup.WlanGroupList, error) {
}

// PartialUpdateRkszonesWlangroupsById
//
// Use this API command to modify the basic information of a WLAN group.
//
// Request Body:
//	 - body *wlangroup.ModifyWlanGroup
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) PartialUpdateRkszonesWlangroupsById(ctx context.Context, body *wlangroup.ModifyWlanGroup, pId string, pZoneId string) (*common.EmptyResult, error) {
}

// PartialUpdateRkszonesWlangroupsMembersByMemberId
//
// Use this API command to modify a member of a WLAN group.
//
// Request Body:
//	 - body *wlangroup.ModifyWlanGroupMember
//
// Path Parameters:
// - pId string
//		- required
// - pMemberId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANGroupService) PartialUpdateRkszonesWlangroupsMembersByMemberId(ctx context.Context, body *wlangroup.ModifyWlanGroupMember, pId string, pMemberId string, pZoneId string) (*common.EmptyResult, error) {
}
