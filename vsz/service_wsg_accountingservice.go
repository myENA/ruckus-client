package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/service"
)

type WSGAccountingServiceService struct {
	client *Client
}

func NewWSGAccountingServiceService(client *Client) *WSGAccountingServiceService {
	s := new(WSGAccountingServiceService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAccountingServiceService() *WSGAccountingServiceService {
	serv := new(WSGAccountingServiceService)
	serv.client = ss.client
	return serv
}

// AddServicesAcctTestById
//
// Use this API command to test an accounting service.
//
// Request Body:
//	 - body *service.TestingConfig
func (s *WSGAccountingServiceService) AddServicesAcctTestById(ctx context.Context, body *service.TestingConfig, pId string) error {
}

// DeleteServicesAcct
//
// Use this API command to delete a list of accounting service.
//
// Request Body:
//	 - body *service.DeleteBulkAccountingService
func (s *WSGAccountingServiceService) DeleteServicesAcct(ctx context.Context, body *service.DeleteBulkAccountingService) (*common.EmptyResult, error) {
}

// DeleteServicesAcctById
//
// Use this API command to delete an accounting service.
func (s *WSGAccountingServiceService) DeleteServicesAcctById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAcctRadiusSecondaryById
//
// Use this API command to disable secondary RADIUS server of a RADIUS accounting service.
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusSecondaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// DeleteServicesAcctRadiusStandbyPrimaryById
//
// Use this API command to disable Standby secondary RADIUS server of a RADIUS Accounting service.
func (s *WSGAccountingServiceService) DeleteServicesAcctRadiusStandbyPrimaryById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindServicesAcctByQueryCriteria
//
// Use this API command to retrieve a list of accounting services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.CommonAccountingServiceList, error) {
}

// FindServicesAcctRadius
//
// Use this API command to retrieve a list of RADIUS accounting services.
func (s *WSGAccountingServiceService) FindServicesAcctRadius(ctx context.Context) (*service.RadiusAccountingServiceList, error) {
}

// FindServicesAcctRadiusById
//
// Use this API command to retrieve a RADIUS accounting service.
func (s *WSGAccountingServiceService) FindServicesAcctRadiusById(ctx context.Context, pId string) (*service.RadiusAccountingService, error) {
}

// FindServicesAcctRadiusByQueryCriteria
//
// Use this API command to retrieve a list of Radius accounting services by query criteria.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGAccountingServiceService) FindServicesAcctRadiusByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*service.RadiusAccountingServiceList, error) {
}
