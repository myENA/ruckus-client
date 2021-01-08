package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// SwitchMSwitchModelModelList
//
// Definition: switchModel_modelList
type SwitchMSwitchModelModelList struct {
	// Model
	// Switch Model
	Model *string `json:"model,omitempty"`
}

func NewSwitchMSwitchModelModelList() *SwitchMSwitchModelModelList {
	m := new(SwitchMSwitchModelModelList)
	return m
}

// SwitchMSwitchModelResult
//
// Definition: switchModel_result
type SwitchMSwitchModelResult struct {
	// Extra
	// Extra field
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// The first data index for current reulst
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of remaining data
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSwitchModelModelList `json:"list,omitempty"`

	// RawDataTotalCount
	// Total Data Count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total Data Count
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMSwitchModelResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSwitchModelResult
}

func newSwitchMSwitchModelResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMSwitchModelResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMSwitchModelResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMSwitchModelResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMSwitchModelResult() *SwitchMSwitchModelResult {
	m := new(SwitchMSwitchModelResult)
	return m
}
