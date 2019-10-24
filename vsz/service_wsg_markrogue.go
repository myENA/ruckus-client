package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGMarkRogueService struct {
	client *Client
}

func NewWSGMarkRogueService(client *Client) *WSGMarkRogueService {
	s := new(WSGMarkRogueService)
	s.client = client
	return s
}

func (ss *WSGService) WSGMarkRogueService() *WSGMarkRogueService {
	serv := new(WSGMarkRogueService)
	serv.client = ss.client
	return serv
}

// AddRogueMarkIgnore
//
// Mark a rogue AP as ignore.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkIgnore(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
}

// AddRogueMarkMalicious
//
// Mark a rogue AP as malicious.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkMalicious(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
}

// AddRogueMarkRogue
//
// Mark a rogue AP as rogue.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkRogue(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
}

// AddRogueUnMark
//
// Use this API command to remove the manual admin classification marking.
//
// Request Body:
//	 - body *ap.ModifyRogueType
func (s *WSGMarkRogueService) AddRogueUnMark(ctx context.Context, body *ap.ModifyRogueType) (*common.EmptyResult, error) {
}

// FindRogueMarkKnown
//
// Get Known Rogue AP list.
func (s *WSGMarkRogueService) FindRogueMarkKnown(ctx context.Context) (*ap.ModifyRogueType, error) {
}
