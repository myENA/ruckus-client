package bigdog

// API Version: v9_1

// WSGAPInfo
//
// Definition: apInfo_apInfo
type WSGAPInfo struct {
	ApMac *WSGCommonMac `json:"apMac,omitempty"`

	ApName *WSGAPInfoApName `json:"apName,omitempty"`

	// LastDetected
	// Timestamp of the AP
	LastDetected *int `json:"lastDetected,omitempty"`

	// MainDetector
	// To identify whether this is main instance for UI
	MainDetector *bool `json:"mainDetector,omitempty"`

	// RogueType
	// Rogue type
	RogueType *string `json:"rogueType,omitempty"`

	// Rssi
	// RSSI of the rogue AP
	Rssi *string `json:"rssi,omitempty"`

	// ZoneName
	// Zone name
	ZoneName *string `json:"zoneName,omitempty"`
}

func NewWSGAPInfo() *WSGAPInfo {
	m := new(WSGAPInfo)
	return m
}

// WSGAPInfoApName
//
// Definition: apInfo_apName
//
// Constraints:
//    - max:64
//    - min:2
type WSGAPInfoApName string

func NewWSGAPInfoApName() *WSGAPInfoApName {
	m := new(WSGAPInfoApName)
	return m
}
