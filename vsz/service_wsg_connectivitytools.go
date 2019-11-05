package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"github.com/myENA/ruckus-client/vsz/types/wsg/tool"
	"gopkg.in/go-playground/validator.v9"
)

type WSGConnectivityToolsService struct {
	apiClient *APIClient
	validate  *validator.Validate
}

func NewWSGConnectivityToolsService(c *APIClient) *WSGConnectivityToolsService {
	s := new(WSGConnectivityToolsService)
	s.apiClient = c
	s.validate = validator.New()
	return s
}

func (ss *WSGService) WSGConnectivityToolsService() *WSGConnectivityToolsService {
	return NewWSGConnectivityToolsService(ss.apiClient)
}

// AddToolSpeedflex
//
// Use this API command to start the SpeedFlex test.
//
// Request Body:
//	 - body *tool.SpeedFlex
func (s *WSGConnectivityToolsService) AddToolSpeedflex(ctx context.Context, body *tool.SpeedFlex) (*tool.TestResult, error) {
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

// FindToolPing
//
// Use this API command to run the PING test on an AP.
//
// Query Parameters:
// - qApMac string
//		- required
// - qTargetIP string
//		- required
func (s *WSGConnectivityToolsService) FindToolPing(ctx context.Context, qApMac string, qTargetIP string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindToolSpeedflexByWcid
//
// Use this API command to retrieve existing SpeedFlex test results.
//
// Path Parameters:
// - pWcid string
//		- required
func (s *WSGConnectivityToolsService) FindToolSpeedflexByWcid(ctx context.Context, pWcid string) (*tool.TestResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindToolTraceRoute
//
// Use this API command to run the traceroute test on an AP.
//
// Query Parameters:
// - qApMac string
//		- required
// - qTargetIP string
//		- required
// - qTimeoutInSec string
func (s *WSGConnectivityToolsService) FindToolTraceRoute(ctx context.Context, qApMac string, qTargetIP string, qTimeoutInSec string) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}
