package bigdog

// API Version: v9_0

type WSGAPRoutineConfigIntervalReq struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`
}

func NewWSGAPRoutineConfigIntervalReq() *WSGAPRoutineConfigIntervalReq {
	m := new(WSGAPRoutineConfigIntervalReq)
	return m
}

type WSGAPRoutineConfigIntervalRsp struct {
	ApRoutineConfigInterval *int `json:"apRoutineConfigInterval,omitempty"`

	LowerBound *int `json:"lowerBound,omitempty"`

	UpperBound *int `json:"upperBound,omitempty"`
}

func NewWSGAPRoutineConfigIntervalRsp() *WSGAPRoutineConfigIntervalRsp {
	m := new(WSGAPRoutineConfigIntervalRsp)
	return m
}
