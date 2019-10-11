package alarmlist

// API Version: v8_1

type AlarmQueryResultList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*SingleAlarm `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewAlarmQueryResultList() *AlarmQueryResultList {
	alarmQueryResultListType := new(AlarmQueryResultList)
	return alarmQueryResultListType
}

func NewAlarmQueryResultListWithDefaults() *AlarmQueryResultList {
	alarmQueryResultListType := new(AlarmQueryResultList)
	return alarmQueryResultListType
}

type SingleAlarm struct {
	Acknowledged *string `json:"acknowledged,omitempty"`

	AckTime *float64 `json:"ackTime,omitempty"`

	AckUser *string `json:"ackUser,omitempty"`

	Activity *string `json:"activity,omitempty"`

	AlarmCode *int `json:"alarmCode,omitempty"`

	AlarmState *string `json:"alarmState,omitempty"`

	AlarmType *string `json:"alarmType,omitempty"`

	Category *string `json:"category,omitempty"`

	ClearComment *string `json:"clearComment,omitempty"`

	ClearTime *float64 `json:"clearTime,omitempty"`

	ClearUser *string `json:"clearUser,omitempty"`

	Id *string `json:"id,omitempty"`

	InsertionTime *float64 `json:"insertionTime,omitempty"`

	Severity *string `json:"severity,omitempty"`
}

func NewSingleAlarm() *SingleAlarm {
	singleAlarmType := new(SingleAlarm)
	return singleAlarmType
}

func NewSingleAlarmWithDefaults() *SingleAlarm {
	singleAlarmType := new(SingleAlarm)
	return singleAlarmType
}
