package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIAPDetailsReportService struct {
	apiClient *SCIClient
}

func NewSCIAPDetailsReportService(c *SCIClient) *SCIAPDetailsReportService {
	s := new(SCIAPDetailsReportService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIAPDetailsReportService() *SCIAPDetailsReportService {
	return NewSCIAPDetailsReportService(ss.apiClient)
}

// SCIAPDetailsReport5trendChartData
//
// Definition: APDetailsReport_APDetailsReport_5_trendChart_Data
type SCIAPDetailsReport5trendChartData [][]*SCIAPDetailsReport5trendChartDataTypeType

func MakeSCIAPDetailsReport5trendChartData() SCIAPDetailsReport5trendChartData {
	m := make(SCIAPDetailsReport5trendChartData, 0)
	return m
}

// SCIAPDetailsReport5trendChartDataType
//
// Definition: APDetailsReport_APDetailsReport_5_trendChart_DataType
type SCIAPDetailsReport5trendChartDataType []*SCIAPDetailsReport5trendChartDataTypeType

func MakeSCIAPDetailsReport5trendChartDataType() SCIAPDetailsReport5trendChartDataType {
	m := make(SCIAPDetailsReport5trendChartDataType, 0)
	return m
}

// SCIAPDetailsReport5trendChartDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_5_trendChart_DataTypeType
type SCIAPDetailsReport5trendChartDataTypeType struct {
	AirtimeBusyAvg *float64 `json:"airtimeBusyAvg,omitempty"`

	AirtimeIdleAvg *float64 `json:"airtimeIdleAvg,omitempty"`

	AirtimeRxAvg *float64 `json:"airtimeRxAvg,omitempty"`

	AirtimeTxAvg *float64 `json:"airtimeTxAvg,omitempty"`

	AirtimeUtilizationAvg *float64 `json:"airtimeUtilizationAvg,omitempty"`

	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIAPDetailsReport5trendChartDataTypeType() *SCIAPDetailsReport5trendChartDataTypeType {
	m := new(SCIAPDetailsReport5trendChartDataTypeType)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableData
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_Data
type SCIAPDetailsReport8topAppsByTrafficTableData []*SCIAPDetailsReport8topAppsByTrafficTableDataType

func MakeSCIAPDetailsReport8topAppsByTrafficTableData() SCIAPDetailsReport8topAppsByTrafficTableData {
	m := make(SCIAPDetailsReport8topAppsByTrafficTableData, 0)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableDataType
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_DataType
type SCIAPDetailsReport8topAppsByTrafficTableDataType struct {
	App *string `json:"app,omitempty"`

	Index *int `json:"index,omitempty"`

	Port *string `json:"port,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport8topAppsByTrafficTableDataType() *SCIAPDetailsReport8topAppsByTrafficTableDataType {
	m := new(SCIAPDetailsReport8topAppsByTrafficTableDataType)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_MetaData
type SCIAPDetailsReport8topAppsByTrafficTableMetaData struct {
	MaxValues *SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *float64 `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIAPDetailsReport8topAppsByTrafficTableMetaData() *SCIAPDetailsReport8topAppsByTrafficTableMetaData {
	m := new(SCIAPDetailsReport8topAppsByTrafficTableMetaData)
	return m
}

// SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_8_topAppsByTrafficTable_MetaDataMaxValuesType
type SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType() *SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport8topAppsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport14topTableData
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_Data
type SCIAPDetailsReport14topTableData []*SCIAPDetailsReport14topTableDataType

func MakeSCIAPDetailsReport14topTableData() SCIAPDetailsReport14topTableData {
	m := make(SCIAPDetailsReport14topTableData, 0)
	return m
}

// SCIAPDetailsReport14topTableDataType
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_DataType
type SCIAPDetailsReport14topTableDataType struct {
	ClientIp *string `json:"clientIp,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Index *int `json:"index,omitempty"`

	Manufacturer *string `json:"manufacturer,omitempty"`

	OsType *string `json:"osType,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewSCIAPDetailsReport14topTableDataType() *SCIAPDetailsReport14topTableDataType {
	m := new(SCIAPDetailsReport14topTableDataType)
	return m
}

// SCIAPDetailsReport14topTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_MetaData
type SCIAPDetailsReport14topTableMetaData struct {
	MaxValues *SCIAPDetailsReport14topTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIAPDetailsReport14topTableMetaData() *SCIAPDetailsReport14topTableMetaData {
	m := new(SCIAPDetailsReport14topTableMetaData)
	return m
}

// SCIAPDetailsReport14topTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_14_topTable_MetaDataMaxValuesType
type SCIAPDetailsReport14topTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport14topTableMetaDataMaxValuesType() *SCIAPDetailsReport14topTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport14topTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport15trendChartData
//
// Definition: APDetailsReport_APDetailsReport_15_trendChart_Data
type SCIAPDetailsReport15trendChartData [][]*SCIAPDetailsReport15trendChartDataTypeType

func MakeSCIAPDetailsReport15trendChartData() SCIAPDetailsReport15trendChartData {
	m := make(SCIAPDetailsReport15trendChartData, 0)
	return m
}

// SCIAPDetailsReport15trendChartDataType
//
// Definition: APDetailsReport_APDetailsReport_15_trendChart_DataType
type SCIAPDetailsReport15trendChartDataType []*SCIAPDetailsReport15trendChartDataTypeType

func MakeSCIAPDetailsReport15trendChartDataType() SCIAPDetailsReport15trendChartDataType {
	m := make(SCIAPDetailsReport15trendChartDataType, 0)
	return m
}

// SCIAPDetailsReport15trendChartDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_15_trendChart_DataTypeType
type SCIAPDetailsReport15trendChartDataTypeType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	Uniqueusers *float64 `json:"unique_users,omitempty"`

	Uniqueusers24 *float64 `json:"unique_users_2-4,omitempty"`

	Uniqueusers5 *float64 `json:"unique_users_5,omitempty"`
}

func NewSCIAPDetailsReport15trendChartDataTypeType() *SCIAPDetailsReport15trendChartDataTypeType {
	m := new(SCIAPDetailsReport15trendChartDataTypeType)
	return m
}

// SCIAPDetailsReport22trafficTrendData
//
// Definition: APDetailsReport_APDetailsReport_22_trafficTrend_Data
type SCIAPDetailsReport22trafficTrendData [][]*SCIAPDetailsReport22trafficTrendDataTypeType

func MakeSCIAPDetailsReport22trafficTrendData() SCIAPDetailsReport22trafficTrendData {
	m := make(SCIAPDetailsReport22trafficTrendData, 0)
	return m
}

// SCIAPDetailsReport22trafficTrendDataType
//
// Definition: APDetailsReport_APDetailsReport_22_trafficTrend_DataType
type SCIAPDetailsReport22trafficTrendDataType []*SCIAPDetailsReport22trafficTrendDataTypeType

func MakeSCIAPDetailsReport22trafficTrendDataType() SCIAPDetailsReport22trafficTrendDataType {
	m := make(SCIAPDetailsReport22trafficTrendDataType, 0)
	return m
}

// SCIAPDetailsReport22trafficTrendDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_22_trafficTrend_DataTypeType
type SCIAPDetailsReport22trafficTrendDataTypeType struct {
	End *string `json:"end,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	Start *string `json:"start,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`
}

func NewSCIAPDetailsReport22trafficTrendDataTypeType() *SCIAPDetailsReport22trafficTrendDataTypeType {
	m := new(SCIAPDetailsReport22trafficTrendDataTypeType)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableData
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_Data
type SCIAPDetailsReport40topSsidsByTrafficTableData []*SCIAPDetailsReport40topSsidsByTrafficTableDataType

func MakeSCIAPDetailsReport40topSsidsByTrafficTableData() SCIAPDetailsReport40topSsidsByTrafficTableData {
	m := make(SCIAPDetailsReport40topSsidsByTrafficTableData, 0)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableDataType
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_DataType
type SCIAPDetailsReport40topSsidsByTrafficTableDataType struct {
	ApCount *float64 `json:"apCount,omitempty"`

	ClientCount *float64 `json:"clientCount,omitempty"`

	Index *int `json:"index,omitempty"`

	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	SessionCount *float64 `json:"sessionCount,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAPDetailsReport40topSsidsByTrafficTableDataType() *SCIAPDetailsReport40topSsidsByTrafficTableDataType {
	m := new(SCIAPDetailsReport40topSsidsByTrafficTableDataType)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_MetaData
type SCIAPDetailsReport40topSsidsByTrafficTableMetaData struct {
	MaxValues *SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType `json:"maxValues,omitempty"`

	Name *string `json:"name,omitempty"`

	Percentage *int `json:"percentage,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	Traffic *int `json:"traffic,omitempty"`
}

func NewSCIAPDetailsReport40topSsidsByTrafficTableMetaData() *SCIAPDetailsReport40topSsidsByTrafficTableMetaData {
	m := new(SCIAPDetailsReport40topSsidsByTrafficTableMetaData)
	return m
}

// SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_40_topSsidsByTrafficTable_MetaDataMaxValuesType
type SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType struct {
	MgmtRxBytes *int `json:"mgmtRxBytes,omitempty"`

	MgmtTraffic *int `json:"mgmtTraffic,omitempty"`

	MgmtTxBytes *int `json:"mgmtTxBytes,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`

	UserRxBytes *int `json:"userRxBytes,omitempty"`

	UserTraffic *int `json:"userTraffic,omitempty"`

	UserTxBytes *int `json:"userTxBytes,omitempty"`
}

func NewSCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType() *SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport40topSsidsByTrafficTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport76apPerformanceData
//
// Definition: APDetailsReport_APDetailsReport_76_apPerformance_Data
type SCIAPDetailsReport76apPerformanceData [][]*SCIAPDetailsReport76apPerformanceDataTypeType

func MakeSCIAPDetailsReport76apPerformanceData() SCIAPDetailsReport76apPerformanceData {
	m := make(SCIAPDetailsReport76apPerformanceData, 0)
	return m
}

// SCIAPDetailsReport76apPerformanceDataType
//
// Definition: APDetailsReport_APDetailsReport_76_apPerformance_DataType
type SCIAPDetailsReport76apPerformanceDataType []*SCIAPDetailsReport76apPerformanceDataTypeType

func MakeSCIAPDetailsReport76apPerformanceDataType() SCIAPDetailsReport76apPerformanceDataType {
	m := make(SCIAPDetailsReport76apPerformanceDataType, 0)
	return m
}

// SCIAPDetailsReport76apPerformanceDataTypeType
//
// Definition: APDetailsReport_APDetailsReport_76_apPerformance_DataTypeType
type SCIAPDetailsReport76apPerformanceDataTypeType struct {
	Avg *float64 `json:"avg,omitempty"`

	Max *float64 `json:"max,omitempty"`

	MaxLabel *string `json:"maxLabel,omitempty"`

	Min *float64 `json:"min,omitempty"`

	MinLabel *string `json:"minLabel,omitempty"`
}

func NewSCIAPDetailsReport76apPerformanceDataTypeType() *SCIAPDetailsReport76apPerformanceDataTypeType {
	m := new(SCIAPDetailsReport76apPerformanceDataTypeType)
	return m
}

// SCIAPDetailsReport78apStatsOverviewData
//
// Definition: APDetailsReport_APDetailsReport_78_apStatsOverview_Data
type SCIAPDetailsReport78apStatsOverviewData []*SCIAPDetailsReport78apStatsOverviewDataType

func MakeSCIAPDetailsReport78apStatsOverviewData() SCIAPDetailsReport78apStatsOverviewData {
	m := make(SCIAPDetailsReport78apStatsOverviewData, 0)
	return m
}

// SCIAPDetailsReport78apStatsOverviewDataType
//
// Definition: APDetailsReport_APDetailsReport_78_apStatsOverview_DataType
type SCIAPDetailsReport78apStatsOverviewDataType struct {
	TotalClientCount *float64 `json:"totalClientCount,omitempty"`

	TotalRxTraffic *int `json:"totalRxTraffic,omitempty"`

	TotalSsids *int `json:"totalSsids,omitempty"`

	TotalTraffic *int `json:"totalTraffic,omitempty"`

	TotalTxTraffic *int `json:"totalTxTraffic,omitempty"`
}

func NewSCIAPDetailsReport78apStatsOverviewDataType() *SCIAPDetailsReport78apStatsOverviewDataType {
	m := new(SCIAPDetailsReport78apStatsOverviewDataType)
	return m
}

// SCIAPDetailsReport79apUptimeHistoryData
//
// Definition: APDetailsReport_APDetailsReport_79_apUptimeHistory_Data
type SCIAPDetailsReport79apUptimeHistoryData []*SCIAPDetailsReport79apUptimeHistoryDataType

func MakeSCIAPDetailsReport79apUptimeHistoryData() SCIAPDetailsReport79apUptimeHistoryData {
	m := make(SCIAPDetailsReport79apUptimeHistoryData, 0)
	return m
}

// SCIAPDetailsReport79apUptimeHistoryDataType
//
// Definition: APDetailsReport_APDetailsReport_79_apUptimeHistory_DataType
type SCIAPDetailsReport79apUptimeHistoryDataType struct {
	End *string `json:"end,omitempty"`

	Start *string `json:"start,omitempty"`

	Status *int `json:"status,omitempty"`
}

func NewSCIAPDetailsReport79apUptimeHistoryDataType() *SCIAPDetailsReport79apUptimeHistoryDataType {
	m := new(SCIAPDetailsReport79apUptimeHistoryDataType)
	return m
}

// SCIAPDetailsReport79apUptimeHistoryMetaData
//
// Definition: APDetailsReport_APDetailsReport_79_apUptimeHistory_MetaData
type SCIAPDetailsReport79apUptimeHistoryMetaData struct {
	TotalDowntime *int `json:"totalDowntime,omitempty"`

	TotalUptime *int `json:"totalUptime,omitempty"`
}

func NewSCIAPDetailsReport79apUptimeHistoryMetaData() *SCIAPDetailsReport79apUptimeHistoryMetaData {
	m := new(SCIAPDetailsReport79apUptimeHistoryMetaData)
	return m
}

// SCIAPDetailsReport81sessionsTableData
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_Data
type SCIAPDetailsReport81sessionsTableData []*SCIAPDetailsReport81sessionsTableDataType

func MakeSCIAPDetailsReport81sessionsTableData() SCIAPDetailsReport81sessionsTableData {
	m := make(SCIAPDetailsReport81sessionsTableData, 0)
	return m
}

// SCIAPDetailsReport81sessionsTableDataType
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_DataType
type SCIAPDetailsReport81sessionsTableDataType struct {
	AuthMethod *string `json:"authMethod,omitempty"`

	ClientMac *string `json:"clientMac,omitempty"`

	DisconnectTime *string `json:"disconnectTime,omitempty"`

	FirstConnection *string `json:"firstConnection,omitempty"`

	Hostname *string `json:"hostname,omitempty"`

	Radio *string `json:"radio,omitempty"`

	RxBytes *int `json:"rxBytes,omitempty"`

	SessionDuration *int `json:"sessionDuration,omitempty"`

	Ssid *string `json:"ssid,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport81sessionsTableDataType() *SCIAPDetailsReport81sessionsTableDataType {
	m := new(SCIAPDetailsReport81sessionsTableDataType)
	return m
}

// SCIAPDetailsReport81sessionsTableMetaData
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_MetaData
type SCIAPDetailsReport81sessionsTableMetaData struct {
	MaxValues *SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType `json:"maxValues,omitempty"`
}

func NewSCIAPDetailsReport81sessionsTableMetaData() *SCIAPDetailsReport81sessionsTableMetaData {
	m := new(SCIAPDetailsReport81sessionsTableMetaData)
	return m
}

// SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType
//
// Definition: APDetailsReport_APDetailsReport_81_sessionsTable_MetaDataMaxValuesType
type SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType struct {
	RxBytes *int `json:"rxBytes,omitempty"`

	Traffic *int `json:"traffic,omitempty"`

	TxBytes *int `json:"txBytes,omitempty"`
}

func NewSCIAPDetailsReport81sessionsTableMetaDataMaxValuesType() *SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType {
	m := new(SCIAPDetailsReport81sessionsTableMetaDataMaxValuesType)
	return m
}

// SCIAPDetailsReport82rssTrendData
//
// Definition: APDetailsReport_APDetailsReport_82_rssTrend_Data
type SCIAPDetailsReport82rssTrendData []*SCIAPDetailsReport82rssTrendDataType

func MakeSCIAPDetailsReport82rssTrendData() SCIAPDetailsReport82rssTrendData {
	m := make(SCIAPDetailsReport82rssTrendData, 0)
	return m
}

// SCIAPDetailsReport82rssTrendDataType
//
// Definition: APDetailsReport_APDetailsReport_82_rssTrend_DataType
type SCIAPDetailsReport82rssTrendDataType struct {
	AvgRss *float64 `json:"avgRss,omitempty"`

	End *string `json:"end,omitempty"`

	MaxRss *float64 `json:"maxRss,omitempty"`

	MinRss *float64 `json:"minRss,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIAPDetailsReport82rssTrendDataType() *SCIAPDetailsReport82rssTrendDataType {
	m := new(SCIAPDetailsReport82rssTrendDataType)
	return m
}

// SCIAPDetailsReport83snrTrendData
//
// Definition: APDetailsReport_APDetailsReport_83_snrTrend_Data
type SCIAPDetailsReport83snrTrendData []*SCIAPDetailsReport83snrTrendDataType

func MakeSCIAPDetailsReport83snrTrendData() SCIAPDetailsReport83snrTrendData {
	m := make(SCIAPDetailsReport83snrTrendData, 0)
	return m
}

// SCIAPDetailsReport83snrTrendDataType
//
// Definition: APDetailsReport_APDetailsReport_83_snrTrend_DataType
type SCIAPDetailsReport83snrTrendDataType struct {
	AvgSnr *float64 `json:"avgSnr,omitempty"`

	End *string `json:"end,omitempty"`

	MaxSnr *float64 `json:"maxSnr,omitempty"`

	MinSnr *float64 `json:"minSnr,omitempty"`

	Start *string `json:"start,omitempty"`
}

func NewSCIAPDetailsReport83snrTrendDataType() *SCIAPDetailsReport83snrTrendDataType {
	m := new(SCIAPDetailsReport83snrTrendDataType)
	return m
}

// SCIAPDetailsReport84alarmsTableData
//
// Definition: APDetailsReport_APDetailsReport_84_alarmsTable_Data
type SCIAPDetailsReport84alarmsTableData []*SCIAPDetailsReport84alarmsTableDataType

func MakeSCIAPDetailsReport84alarmsTableData() SCIAPDetailsReport84alarmsTableData {
	m := make(SCIAPDetailsReport84alarmsTableData, 0)
	return m
}

// SCIAPDetailsReport84alarmsTableDataType
//
// Definition: APDetailsReport_APDetailsReport_84_alarmsTable_DataType
type SCIAPDetailsReport84alarmsTableDataType struct {
	AlarmCode *string `json:"alarmCode,omitempty"`

	AlarmState *string `json:"alarmState,omitempty"`

	AlarmType *string `json:"alarmType,omitempty"`

	AlarmUUID *string `json:"alarmUUID,omitempty"`

	Category *string `json:"category,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Severity *string `json:"severity,omitempty"`

	Xtime *string `json:"__time,omitempty"`
}

func NewSCIAPDetailsReport84alarmsTableDataType() *SCIAPDetailsReport84alarmsTableDataType {
	m := new(SCIAPDetailsReport84alarmsTableDataType)
	return m
}

// SCIAPDetailsReport85eventsTableData
//
// Definition: APDetailsReport_APDetailsReport_85_eventsTable_Data
type SCIAPDetailsReport85eventsTableData []*SCIAPDetailsReport85eventsTableDataType

func MakeSCIAPDetailsReport85eventsTableData() SCIAPDetailsReport85eventsTableData {
	m := make(SCIAPDetailsReport85eventsTableData, 0)
	return m
}

// SCIAPDetailsReport85eventsTableDataType
//
// Definition: APDetailsReport_APDetailsReport_85_eventsTable_DataType
type SCIAPDetailsReport85eventsTableDataType struct {
	Category *string `json:"category,omitempty"`

	EventCode *string `json:"eventCode,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	Reason *string `json:"reason,omitempty"`

	Severity *string `json:"severity,omitempty"`

	Xtime *string `json:"__time,omitempty"`
}

func NewSCIAPDetailsReport85eventsTableDataType() *SCIAPDetailsReport85eventsTableDataType {
	m := new(SCIAPDetailsReport85eventsTableDataType)
	return m
}

// ReportAPDetailsReport5TrendChart
//
// Operation ID: report_APDetailsReport_5_trendChart
//
// AP Details Report - Airtime Utilization Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport5TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport5trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport5trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport5TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport5trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport7Top10ApplicationsByTrafficVolume
//
// Operation ID: report_APDetailsReport_7_top10ApplicationsByTrafficVolume
//
// AP Details Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport7Top10ApplicationsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport7Top10ApplicationsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport7top10ApplicationsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport8TopAppsByTrafficTable
//
// Operation ID: report_APDetailsReport_8_topAppsByTrafficTable
//
// AP Details Report - Top Applications by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport8TopAppsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport8TopAppsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport8topAppsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport14TopTable
//
// Operation ID: report_APDetailsReport_14_topTable
//
// AP Details Report - Clients Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport14TopTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport14topTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport14topTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport14TopTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport14topTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport15TrendChart
//
// Operation ID: report_APDetailsReport_15_trendChart
//
// AP Details Report - Unique Clients Trend Over Time
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport15TrendChart(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport15trendChart200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport15trendChart200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport15TrendChart, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport15trendChart200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport22TrafficTrend
//
// Operation ID: report_APDetailsReport_22_trafficTrend
//
// AP Details Report - Traffic Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport22TrafficTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport22trafficTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport22trafficTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport22TrafficTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport22trafficTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport40TopSsidsByTrafficTable
//
// Operation ID: report_APDetailsReport_40_topSsidsByTrafficTable
//
// AP Details Report - Top SSIDs by Traffic
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport40TopSsidsByTrafficTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport40TopSsidsByTrafficTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport40topSsidsByTrafficTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport75ApSummary
//
// Operation ID: report_APDetailsReport_75_apSummary
//
// AP Details Report - Summary
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport75ApSummary(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport75apSummary200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport75apSummary200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport75ApSummary, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport75apSummary200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport76ApPerformance
//
// Operation ID: report_APDetailsReport_76_apPerformance
//
// AP Details Report - Performance
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport76ApPerformance(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport76apPerformance200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport76apPerformance200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport76ApPerformance, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport76apPerformance200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport77ApDetails
//
// Operation ID: report_APDetailsReport_77_apDetails
//
// AP Details Report - Details
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport77ApDetails(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport77apDetails200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport77apDetails200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport77ApDetails, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport77apDetails200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport78ApStatsOverview
//
// Operation ID: report_APDetailsReport_78_apStatsOverview
//
// AP Details Report - Stats
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport78ApStatsOverview(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport78apStatsOverview200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport78apStatsOverview200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport78ApStatsOverview, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport78apStatsOverview200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport79ApUptimeHistory
//
// Operation ID: report_APDetailsReport_79_apUptimeHistory
//
// AP Details Report - Uptime History
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport79ApUptimeHistory(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport79apUptimeHistory200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport79apUptimeHistory200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport79ApUptimeHistory, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport79apUptimeHistory200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport80Top10ClientsByTrafficVolume
//
// Operation ID: report_APDetailsReport_80_top10ClientsByTrafficVolume
//
// AP Details Report - Top 10 Clients by Traffic Volume
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport80Top10ClientsByTrafficVolume(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport80Top10ClientsByTrafficVolume, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport80top10ClientsByTrafficVolume200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport81SessionsTable
//
// Operation ID: report_APDetailsReport_81_sessionsTable
//
// AP Details Report - Sessions
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport81SessionsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport81sessionsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport81sessionsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport81SessionsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport81sessionsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport82RssTrend
//
// Operation ID: report_APDetailsReport_82_rssTrend
//
// AP Details Report - RSS Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport82RssTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport82rssTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport82rssTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport82RssTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport82rssTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport83SnrTrend
//
// Operation ID: report_APDetailsReport_83_snrTrend
//
// AP Details Report - SNR Trend
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport83SnrTrend(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport83snrTrend200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport83snrTrend200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport83SnrTrend, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport83snrTrend200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport84AlarmsTable
//
// Operation ID: report_APDetailsReport_84_alarmsTable
//
// AP Details Report - Alarms
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport84AlarmsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport84alarmsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport84alarmsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport84AlarmsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport84alarmsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport85EventsTable
//
// Operation ID: report_APDetailsReport_85_eventsTable
//
// AP Details Report - Events
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport85EventsTable(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport85eventsTable200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport85eventsTable200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport85EventsTable, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport85eventsTable200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport95Anomalies
//
// Operation ID: report_APDetailsReport_95_anomalies
//
// AP Details Report - Anomalies
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport95Anomalies(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport95anomalies200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport95anomalies200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport95Anomalies, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport95anomalies200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ReportAPDetailsReport110ApAnomaly
//
// Operation ID: report_APDetailsReport_110_apAnomaly
//
// AP Details Report - Anomalies for the Past 30 Days
//
// Request Body:
//	 - body *SCICommonReportQueryBody
func (s *SCIAPDetailsReportService) ReportAPDetailsReport110ApAnomaly(ctx context.Context, body *SCICommonReportQueryBody, mutators ...RequestMutator) (*SCIReportAPDetailsReport110apAnomaly200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCIReportAPDetailsReport110apAnomaly200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIReportAPDetailsReport110ApAnomaly, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSCIReportAPDetailsReport110apAnomaly200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
