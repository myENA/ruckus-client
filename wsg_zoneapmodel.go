package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGZoneAPModelApModel
//
// Definition: zoneApmodel_apModel
type WSGZoneAPModelApModel struct {
	CellularSettings *WSGAPModelCellularSettings `json:"cellularSettings,omitempty"`

	ExternalAntenna24 *WSGAPModelExternalAntenna `json:"externalAntenna24,omitempty"`

	ExternalAntenna50 *WSGAPModelExternalAntenna `json:"externalAntenna50,omitempty"`

	// InternalHeaterEnabled
	// Constraints:
	//    - nullable
	InternalHeaterEnabled *bool `json:"internalHeaterEnabled,omitempty"`

	Lacp *WSGAPModelLacpSetting `json:"lacp,omitempty"`

	LanPorts []*WSGZoneAPModelLanPortSetting `json:"lanPorts,omitempty"`

	// LedMode
	// Constraints:
	//    - nullable
	//    - oneof:[CableModem,AccessPoint,CableModem_AccessPoint,CableModem_AccessPoint_DEFAULT,ActiveSurgeProtector,ActiveSurgeProtector_ModemOnline_DEFAULT,Off]
	LedMode *string `json:"ledMode,omitempty"`

	// LedStatusEnabled
	// Constraints:
	//    - nullable
	LedStatusEnabled *bool `json:"ledStatusEnabled,omitempty"`

	Lldp *WSGAPModelLldpSetting `json:"lldp,omitempty"`

	PoeModeSetting *WSGCommonPoeModeSetting `json:"poeModeSetting,omitempty"`

	// PoeOutPortEnabled
	// Constraints:
	//    - nullable
	PoeOutPortEnabled *bool `json:"poeOutPortEnabled,omitempty"`

	// PoeTxChain
	// Option to use 1, 2 or 4 Tx chains while AP power source is 802.3af PoE
	// Constraints:
	//    - nullable
	//    - default:2
	PoeTxChain *int `json:"poeTxChain,omitempty"`

	// RadioBand
	// Band switch between 2.4GHz and 5GHz is provided in single radio AP ZF-7321, ZF-7321-U, and ZF-7441.
	// Constraints:
	//    - nullable
	//    - oneof:[2.4GHz,5GHz]
	RadioBand *string `json:"radioBand,omitempty"`

	// UsbPowerEnable
	// Constraints:
	//    - nullable
	UsbPowerEnable *bool `json:"usbPowerEnable,omitempty"`
}

type WSGZoneAPModelApModelAPIResponse struct {
	*RawAPIResponse
	Data *WSGZoneAPModelApModel
}

func newWSGZoneAPModelApModelAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGZoneAPModelApModelAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGZoneAPModelApModelAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGZoneAPModelApModel)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGZoneAPModelApModel() *WSGZoneAPModelApModel {
	m := new(WSGZoneAPModelApModel)
	return m
}

// WSGZoneAPModelLanPortSetting
//
// Definition: zoneApmodel_lanPortSetting
type WSGZoneAPModelLanPortSetting struct {
	// Enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`

	EthPortProfile *WSGCommonGenericRef `json:"ethPortProfile,omitempty"`

	// PortName
	// Constraints:
	//    - required
	//    - oneof:[LAN1,LAN2,LAN3,LAN4,LAN5]
	PortName *string `json:"portName"`
}

func NewWSGZoneAPModelLanPortSetting() *WSGZoneAPModelLanPortSetting {
	m := new(WSGZoneAPModelLanPortSetting)
	return m
}
