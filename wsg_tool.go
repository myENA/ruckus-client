package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGToolSpeedFlex
//
// Definition: tool_speedFlex
type WSGToolSpeedFlex struct {
	ClientIp *WSGCommonIpAddress `json:"clientIp,omitempty"`

	ClientMac *WSGCommonMac `json:"clientMac,omitempty"`

	// Model
	// Test model
	// Constraints:
	//    - oneof:[AP,CLIENT,TRACE,HOP,NULL]
	Model *string `json:"model,omitempty"`

	// Protocol
	// Protocol used in the SpeedFlex test
	// Constraints:
	//    - required
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol"`

	ServerIp *WSGCommonIpAddress `json:"serverIp,omitempty"`

	ServerMac *WSGCommonMac `json:"serverMac,omitempty"`

	// Syspmtu
	// Default: 1500
	Syspmtu *int `json:"syspmtu,omitempty"`

	// Tool
	// SpeedFlex tool
	// Constraints:
	//    - required
	//    - oneof:[ZAP_DOWN,ZAP_UP]
	Tool *string `json:"tool"`
}

func NewWSGToolSpeedFlex() *WSGToolSpeedFlex {
	m := new(WSGToolSpeedFlex)
	return m
}

// WSGToolTestResult
//
// Definition: tool_testResult
type WSGToolTestResult struct {
	// Downlink
	// Downlink
	Downlink *int `json:"downlink,omitempty"`

	// Etf
	// ETF
	Etf *int `json:"etf,omitempty"`

	// Latency
	// Latency
	Latency *int `json:"latency,omitempty"`

	// PacketLoss
	// Packet loss
	PacketLoss *int `json:"packetLoss,omitempty"`

	// ResultId
	// Result ID
	ResultId *int `json:"resultId,omitempty"`

	// Uplink
	// Uplink
	Uplink *int `json:"uplink,omitempty"`

	Wcid *string `json:"wcid,omitempty"`
}

type WSGToolTestResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGToolTestResult
}

func newWSGToolTestResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGToolTestResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGToolTestResultAPIResponse) Hydrate() error {
	r.Data = new(WSGToolTestResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGToolTestResult() *WSGToolTestResult {
	m := new(WSGToolTestResult)
	return m
}
