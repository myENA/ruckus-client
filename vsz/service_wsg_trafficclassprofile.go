package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
	"gopkg.in/go-playground/validator.v9"
)

type WSGTrafficClassProfileService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGTrafficClassProfileService(c *APIClient) *WSGTrafficClassProfileService {
	s := new(WSGTrafficClassProfileService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGTrafficClassProfileService() *WSGTrafficClassProfileService {
	return NewWSGTrafficClassProfileService(ss.apiClient)
}

// AddRkszonesTrafficClassProfileByZoneId
//
// Use this API command to create a new Traffic Class Profile of a zone.
//
// Request Body:
//	 - body *profile.CreateTrafficClassProfile
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) AddRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *profile.CreateTrafficClassProfile, pZoneId string) (*common.CreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}

// DeleteRkszonesTrafficClassProfileById
//
// Use this API command to delete a Traffic Class Profile of a zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileById(ctx context.Context, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteRkszonesTrafficClassProfileByZoneId
//
// Use this API command to bulk delete Traffic Class Profiles of a zone.
//
// Request Body:
//	 - body *common.BulkDeleteRequest
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) DeleteRkszonesTrafficClassProfileByZoneId(ctx context.Context, body *common.BulkDeleteRequest, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// FindRkszonesTrafficClassProfileById
//
// Use this API command to retrieve a Traffic Class Profile of zone.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileById(ctx context.Context, pId string, pZoneId string) (*common.TrafficClassProfileRef, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindRkszonesTrafficClassProfileByZoneId
//
// Use this API command to retrieve a list of Traffic Class Profile of a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) FindRkszonesTrafficClassProfileByZoneId(ctx context.Context, pZoneId string) (*profile.TrafficClassProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindServicesTrafficClassProfileByQueryCriteria
//
// Retrieve a list of Traffic Class Profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGTrafficClassProfileService) FindServicesTrafficClassProfileByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.TrafficClassProfileList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
}

// PartialUpdateRkszonesTrafficClassProfileById
//
// Use this API command to modify Traffic Class Profile of a zone.
//
// Request Body:
//	 - body *profile.CreateTrafficClassProfile
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGTrafficClassProfileService) PartialUpdateRkszonesTrafficClassProfileById(ctx context.Context, body *profile.CreateTrafficClassProfile, pId string, pZoneId string) (*common.EmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
	if body == nil {
		return nil, errors.New("body cannot be empty")
	}
	if err := s.validate.StructCtx(ctx, body); err != nil {
		return nil, err
	}
}
