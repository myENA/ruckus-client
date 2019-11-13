package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAccessPointAPPService struct {
	apiClient *APIClient
}

func NewWSGAccessPointAPPService(c *APIClient) *WSGAccessPointAPPService {
	s := new(WSGAccessPointAPPService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointAPPService() *WSGAccessPointAPPService {
	return NewWSGAccessPointAPPService(ss.apiClient)
}

// FindApsLineman
//
// Use this API command to retrieve the summary information of an AP. This is used by the Ruckus Wireless AP mobile app.
//
// Query Parameters:
// - qDomainId string
//		- nullable
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
// - qShowAlarm string
//		- nullable
// - qZoneId string
//		- nullable
func (s *WSGAccessPointAPPService) FindApsLineman(ctx context.Context, qDomainId string, qIndex string, qListSize string, qShowAlarm string, qZoneId string) (*WSGAPLinemanSummary, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsTotalCount
//
// Use this API command to retrieve the total AP count within a zone or a domain.
//
// Query Parameters:
// - qDomainId string
//		- nullable
// - qZoneId string
//		- nullable
func (s *WSGAccessPointAPPService) FindApsTotalCount(ctx context.Context, qDomainId string, qZoneId string) (interface{}, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLinemanWorkflow
//
// Use this API command to download the workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) FindLinemanWorkflow(ctx context.Context) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateLinemanWorkflow
//
// Use this API command to upload a workflow file used by the Ruckus Wireless AP mobile app.
func (s *WSGAccessPointAPPService) UpdateLinemanWorkflow(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}