package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type WSGAdministrationService struct {
	apiClient *VSZClient
}

func NewWSGAdministrationService(c *VSZClient) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	return NewWSGAdministrationService(ss.apiClient)
}

// WSGAdministrationActiveDirectoryServer
//
// Definition: administration_activeDirectoryServer
type WSGAdministrationActiveDirectoryServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Port number of Active Directory Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// WindowsDomainName
	// Windows Domain Name of Active Directory Server object
	// Constraints:
	//    - required
	WindowsDomainName *string `json:"windowsDomainName"`
}

func NewWSGAdministrationActiveDirectoryServer() *WSGAdministrationActiveDirectoryServer {
	m := new(WSGAdministrationActiveDirectoryServer)
	return m
}

// WSGAdministrationApPatchHistory
//
// Definition: administration_apPatchHistory
type WSGAdministrationApPatchHistory struct {
	// ApFwVersion
	// apFwVersion of the AP Patch history
	ApFwVersion *string `json:"apFwVersion,omitempty"`

	// ApModelList
	// AP Models of the AP Patch history
	ApModelList []string `json:"apModelList,omitempty"`

	// FileName
	// file name of the AP Patch history
	FileName *string `json:"fileName,omitempty"`

	// StartDateTime
	// startDateTime of the AP Patch history
	StartDateTime *string `json:"startDateTime,omitempty"`
}

func NewWSGAdministrationApPatchHistory() *WSGAdministrationApPatchHistory {
	m := new(WSGAdministrationApPatchHistory)
	return m
}

// WSGAdministrationApPatchHistoryList
//
// Definition: administration_apPatchHistoryList
type WSGAdministrationApPatchHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationApPatchHistory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationApPatchHistoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApPatchHistoryList
}

func newWSGAdministrationApPatchHistoryListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApPatchHistoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApPatchHistoryListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationApPatchHistoryList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationApPatchHistoryList() *WSGAdministrationApPatchHistoryList {
	m := new(WSGAdministrationApPatchHistoryList)
	return m
}

// WSGAdministrationApPatchInfo
//
// Definition: administration_apPatchInfo
type WSGAdministrationApPatchInfo struct {
	// ApModels
	// AP Models of the upload file
	ApModels []string `json:"apModels,omitempty"`

	// ApVersion
	// ApFwVersion of the upload file
	ApVersion *string `json:"apVersion,omitempty"`

	// FileName
	// file name of the upload file
	FileName *string `json:"fileName,omitempty"`

	// FileSize
	// file size(Byte) of the upload file
	FileSize *int `json:"fileSize,omitempty"`
}

type WSGAdministrationApPatchInfoAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApPatchInfo
}

func newWSGAdministrationApPatchInfoAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApPatchInfoAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApPatchInfoAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationApPatchInfo)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationApPatchInfo() *WSGAdministrationApPatchInfo {
	m := new(WSGAdministrationApPatchInfo)
	return m
}

// WSGAdministrationApPatchStatus
//
// Definition: administration_apPatchStatus
type WSGAdministrationApPatchStatus struct {
	ClusterOperationProgress *WSGClusterBladeClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

type WSGAdministrationApPatchStatusAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApPatchStatus
}

func newWSGAdministrationApPatchStatusAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApPatchStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApPatchStatusAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationApPatchStatus)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationApPatchStatus() *WSGAdministrationApPatchStatus {
	m := new(WSGAdministrationApPatchStatus)
	return m
}

// WSGAdministrationApplicationLogAndStatus
//
// Definition: administration_applicationLogAndStatus
type WSGAdministrationApplicationLogAndStatus struct {
	// ApplicationName
	// Application name
	ApplicationName *string `json:"applicationName,omitempty"`

	// HealthStatus
	// Health status
	HealthStatus *string `json:"healthStatus,omitempty"`

	// LogFileNames
	// List of log file name
	LogFileNames []string `json:"logFileNames,omitempty"`

	// LogLevel
	// Log level
	LogLevel *string `json:"logLevel,omitempty"`

	// NumOfLogs
	// # of Logs
	NumOfLogs *int `json:"numOfLogs,omitempty"`
}

func NewWSGAdministrationApplicationLogAndStatus() *WSGAdministrationApplicationLogAndStatus {
	m := new(WSGAdministrationApplicationLogAndStatus)
	return m
}

// WSGAdministrationApplicationLogAndStatusList
//
// Definition: administration_applicationLogAndStatusList
type WSGAdministrationApplicationLogAndStatusList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationApplicationLogAndStatus `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationApplicationLogAndStatusListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationApplicationLogAndStatusList
}

func newWSGAdministrationApplicationLogAndStatusListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationApplicationLogAndStatusListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationApplicationLogAndStatusListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationApplicationLogAndStatusList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationApplicationLogAndStatusList() *WSGAdministrationApplicationLogAndStatusList {
	m := new(WSGAdministrationApplicationLogAndStatusList)
	return m
}

// WSGAdministrationAutoExportBackup
//
// Definition: administration_autoExportBackup
type WSGAdministrationAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// FTP server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

type WSGAdministrationAutoExportBackupAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationAutoExportBackup
}

func newWSGAdministrationAutoExportBackupAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationAutoExportBackupAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationAutoExportBackupAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationAutoExportBackup)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationAutoExportBackup() *WSGAdministrationAutoExportBackup {
	m := new(WSGAdministrationAutoExportBackup)
	return m
}

// WSGAdministrationBackupFile
//
// Definition: administration_backupFile
type WSGAdministrationBackupFile struct {
	// BackupElapsed
	// backup elapsed of the configuration backup file
	BackupElapsed *int `json:"backupElapsed,omitempty"`

	// ControlPlaneSoftwareVersion
	// control plane software version of the configuration backup file
	ControlPlaneSoftwareVersion *string `json:"controlPlaneSoftwareVersion,omitempty"`

	// CreatedBy
	// creator of the configuration backup file.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedOn
	// the create time of the configuration backup file.
	CreatedOn *float64 `json:"createdOn,omitempty"`

	// DataPlaneSoftwareVersion
	// data plane software version of the configuration backup file
	DataPlaneSoftwareVersion *string `json:"dataPlaneSoftwareVersion,omitempty"`

	// FileSize
	// file size of the backup file
	FileSize *int `json:"fileSize,omitempty"`

	// Id
	// Identifier of system configuration backup file.
	Id *string `json:"id,omitempty"`

	// Md5
	// file md5 of the backup file
	Md5 *string `json:"md5,omitempty"`

	// ScgVersion
	// SCG version of the configuration backup file.
	ScgVersion *string `json:"scgVersion,omitempty"`

	// Type
	// type of the configuration backup file
	Type *string `json:"type,omitempty"`
}

func NewWSGAdministrationBackupFile() *WSGAdministrationBackupFile {
	m := new(WSGAdministrationBackupFile)
	return m
}

// WSGAdministrationClusterBackupList
//
// Definition: administration_clusterBackupList
type WSGAdministrationClusterBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationClusterBackupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationClusterBackupListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationClusterBackupList
}

func newWSGAdministrationClusterBackupListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationClusterBackupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationClusterBackupListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationClusterBackupList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationClusterBackupList() *WSGAdministrationClusterBackupList {
	m := new(WSGAdministrationClusterBackupList)
	return m
}

// WSGAdministrationClusterBackupSummary
//
// Definition: administration_clusterBackupSummary
type WSGAdministrationClusterBackupSummary struct {
	// CreatedOn
	// Created date and time of the cluster backup file
	CreatedOn *string `json:"createdOn,omitempty"`

	// Filesize
	// filesize of the cluster backup file.
	Filesize *float64 `json:"filesize,omitempty"`

	// Id
	// Identifier of cluster backup file.
	Id *string `json:"id,omitempty"`

	// Version
	// the patch version of the cluster backup file.
	Version *string `json:"version,omitempty"`
}

func NewWSGAdministrationClusterBackupSummary() *WSGAdministrationClusterBackupSummary {
	m := new(WSGAdministrationClusterBackupSummary)
	return m
}

// WSGAdministrationConfigurationBackupList
//
// Definition: administration_configurationBackupList
type WSGAdministrationConfigurationBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationBackupFile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationConfigurationBackupListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationConfigurationBackupList
}

func newWSGAdministrationConfigurationBackupListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationConfigurationBackupListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationConfigurationBackupListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationConfigurationBackupList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationConfigurationBackupList() *WSGAdministrationConfigurationBackupList {
	m := new(WSGAdministrationConfigurationBackupList)
	return m
}

// WSGAdministrationConnectZD
//
// Definition: administration_connectZD
type WSGAdministrationConnectZD struct {
	Ip *string `json:"ip,omitempty"`

	Password *string `json:"password,omitempty"`

	User *string `json:"user,omitempty"`
}

func NewWSGAdministrationConnectZD() *WSGAdministrationConnectZD {
	m := new(WSGAdministrationConnectZD)
	return m
}

// WSGAdministrationCreateAdminAAAServer
//
// Definition: administration_createAdminAAAServer
type WSGAdministrationCreateAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type"`
}

func NewWSGAdministrationCreateAdminAAAServer() *WSGAdministrationCreateAdminAAAServer {
	m := new(WSGAdministrationCreateAdminAAAServer)
	return m
}

// WSGAdministrationDefaultRoleMapping
//
// Definition: administration_defaultRoleMapping
type WSGAdministrationDefaultRoleMapping struct {
	// DefaultAdmin
	// DefaultAdmin of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultAdmin *string `json:"defaultAdmin"`

	// DefaultUserGroup
	// DefaultUserGroup of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultUserGroup *string `json:"defaultUserGroup"`
}

func NewWSGAdministrationDefaultRoleMapping() *WSGAdministrationDefaultRoleMapping {
	m := new(WSGAdministrationDefaultRoleMapping)
	return m
}

// WSGAdministrationLdapServer
//
// Definition: administration_ldapServer
type WSGAdministrationLdapServer struct {
	// AdminDomainName
	// Admin Domain Name of LDAP Server object
	// Constraints:
	//    - required
	AdminDomainName *string `json:"adminDomainName"`

	// AdminPassword
	// Admin password of LDAP Server object
	// Constraints:
	//    - required
	AdminPassword *string `json:"adminPassword"`

	// BaseDomainName
	// Base Domain Name of LDAP Server object
	// Constraints:
	//    - required
	BaseDomainName *string `json:"baseDomainName"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// KeyAttribute
	// Key attribute of LDAP Server object
	// Constraints:
	//    - required
	KeyAttribute *string `json:"keyAttribute"`

	// Port
	// Port number of LDAP Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	// SearchFilter
	// Search filter of LDAP Server object
	// Constraints:
	//    - required
	SearchFilter *string `json:"searchFilter"`
}

func NewWSGAdministrationLdapServer() *WSGAdministrationLdapServer {
	m := new(WSGAdministrationLdapServer)
	return m
}

// WSGAdministrationLicenses
//
// Definition: administration_licenses
type WSGAdministrationLicenses struct {
	// Count
	// number of licenses
	Count *int `json:"count,omitempty"`

	// CreateTime
	// license effective date
	CreateTime *string `json:"createTime,omitempty"`

	// Description
	// license description
	Description *string `json:"description,omitempty"`

	// ExpireDate
	// license expiry date
	ExpireDate *string `json:"expireDate,omitempty"`

	// Name
	// license name
	Name *string `json:"name,omitempty"`
}

func NewWSGAdministrationLicenses() *WSGAdministrationLicenses {
	m := new(WSGAdministrationLicenses)
	return m
}

// WSGAdministrationLicenseServer
//
// Definition: administration_licenseServer
type WSGAdministrationLicenseServer struct {
	// IpAddress
	// local license server IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// local license server port
	// Constraints:
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// UseCloud
	// use cloud license server
	UseCloud *bool `json:"useCloud,omitempty"`
}

type WSGAdministrationLicenseServerAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicenseServer
}

func newWSGAdministrationLicenseServerAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicenseServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicenseServerAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationLicenseServer)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationLicenseServer() *WSGAdministrationLicenseServer {
	m := new(WSGAdministrationLicenseServer)
	return m
}

// WSGAdministrationLicensesList
//
// Definition: administration_licensesList
type WSGAdministrationLicensesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicenses `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicensesList
}

func newWSGAdministrationLicensesListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicensesListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicensesListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationLicensesList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationLicensesList() *WSGAdministrationLicensesList {
	m := new(WSGAdministrationLicensesList)
	return m
}

// WSGAdministrationLicensesSummary
//
// Definition: administration_licensesSummary
type WSGAdministrationLicensesSummary struct {
	CapacityControlLicenseCount *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`

	// LicenseTypeDescription
	// license type description
	LicenseTypeDescription *string `json:"licenseTypeDescription,omitempty"`
}

func NewWSGAdministrationLicensesSummary() *WSGAdministrationLicensesSummary {
	m := new(WSGAdministrationLicensesSummary)
	return m
}

// WSGAdministrationLicensesSummaryCapacityControlLicenseCountType
//
// Definition: administration_licensesSummaryCapacityControlLicenseCountType
type WSGAdministrationLicensesSummaryCapacityControlLicenseCountType struct {
	// TotalCount
	// total count of licenses
	TotalCount *int `json:"totalCount,omitempty"`

	// UsedCount
	// consumed count of licenses
	UsedCount *int `json:"usedCount,omitempty"`
}

func NewWSGAdministrationLicensesSummaryCapacityControlLicenseCountType() *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType {
	m := new(WSGAdministrationLicensesSummaryCapacityControlLicenseCountType)
	return m
}

// WSGAdministrationLicensesSummaryList
//
// Definition: administration_licensesSummaryList
type WSGAdministrationLicensesSummaryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicensesSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesSummaryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicensesSummaryList
}

func newWSGAdministrationLicensesSummaryListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicensesSummaryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicensesSummaryListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationLicensesSummaryList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationLicensesSummaryList() *WSGAdministrationLicensesSummaryList {
	m := new(WSGAdministrationLicensesSummaryList)
	return m
}

// WSGAdministrationLicensesSyncLogs
//
// Definition: administration_licensesSyncLogs
type WSGAdministrationLicensesSyncLogs struct {
	// CreateDateTime
	// license sync log's create time
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// SyncResult
	// sync license result
	// Constraints:
	//    - oneof:[SUCCESS,FAILURE]
	SyncResult *string `json:"syncResult,omitempty"`
}

func NewWSGAdministrationLicensesSyncLogs() *WSGAdministrationLicensesSyncLogs {
	m := new(WSGAdministrationLicensesSyncLogs)
	return m
}

// WSGAdministrationLicensesSyncLogsList
//
// Definition: administration_licensesSyncLogsList
type WSGAdministrationLicensesSyncLogsList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicensesSyncLogs `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesSyncLogsListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationLicensesSyncLogsList
}

func newWSGAdministrationLicensesSyncLogsListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationLicensesSyncLogsListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationLicensesSyncLogsListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationLicensesSyncLogsList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationLicensesSyncLogsList() *WSGAdministrationLicensesSyncLogsList {
	m := new(WSGAdministrationLicensesSyncLogsList)
	return m
}

// WSGAdministrationModfiyLicenseServer
//
// Definition: administration_modfiyLicenseServer
type WSGAdministrationModfiyLicenseServer struct {
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// Constraints:
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty"`

	// UseCloud
	// Constraints:
	//    - required
	UseCloud *bool `json:"useCloud"`
}

func NewWSGAdministrationModfiyLicenseServer() *WSGAdministrationModfiyLicenseServer {
	m := new(WSGAdministrationModfiyLicenseServer)
	return m
}

// WSGAdministrationModifyAdminAAAServer
//
// Definition: administration_modifyAdminAAAServer
type WSGAdministrationModifyAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type"`
}

func NewWSGAdministrationModifyAdminAAAServer() *WSGAdministrationModifyAdminAAAServer {
	m := new(WSGAdministrationModifyAdminAAAServer)
	return m
}

// WSGAdministrationModifyAutoExportBackup
//
// Definition: administration_modifyAutoExportBackup
type WSGAdministrationModifyAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// ftp server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

func NewWSGAdministrationModifyAutoExportBackup() *WSGAdministrationModifyAutoExportBackup {
	m := new(WSGAdministrationModifyAutoExportBackup)
	return m
}

// WSGAdministrationModifyLogLevel
//
// Definition: administration_modifyLogLevel
type WSGAdministrationModifyLogLevel struct {
	// ApplicationName
	// Application name.
	ApplicationName *string `json:"applicationName,omitempty"`

	// LogLevel
	// Log level.
	// Constraints:
	//    - oneof:[DEBUG,INFO,WARN,ERROR]
	LogLevel *string `json:"logLevel,omitempty"`
}

func NewWSGAdministrationModifyLogLevel() *WSGAdministrationModifyLogLevel {
	m := new(WSGAdministrationModifyLogLevel)
	return m
}

// WSGAdministrationModifyScheduleBackup
//
// Definition: administration_modifyScheduleBackup
type WSGAdministrationModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

func NewWSGAdministrationModifyScheduleBackup() *WSGAdministrationModifyScheduleBackup {
	m := new(WSGAdministrationModifyScheduleBackup)
	return m
}

// WSGAdministrationRadiusServer
//
// Definition: administration_radiusServer
type WSGAdministrationRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Port number of RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Protocol
	// Constraints:
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP,PEAP]
	Protocol *string `json:"protocol,omitempty"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm"`

	SecondaryRadiusServer *WSGAdministrationSecondaryRadiusServer `json:"secondaryRadiusServer,omitempty"`

	// SharedSecret
	// Shared secret of RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGAdministrationRadiusServer() *WSGAdministrationRadiusServer {
	m := new(WSGAdministrationRadiusServer)
	return m
}

// WSGAdministrationRetrieveAdminAAAServer
//
// Definition: administration_retrieveAdminAAAServer
type WSGAdministrationRetrieveAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	// Id
	// ID of this Admin AAA Server
	Id *string `json:"id,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty"`
}

type WSGAdministrationRetrieveAdminAAAServerAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationRetrieveAdminAAAServer
}

func newWSGAdministrationRetrieveAdminAAAServerAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationRetrieveAdminAAAServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationRetrieveAdminAAAServerAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationRetrieveAdminAAAServer)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationRetrieveAdminAAAServer() *WSGAdministrationRetrieveAdminAAAServer {
	m := new(WSGAdministrationRetrieveAdminAAAServer)
	return m
}

// WSGAdministrationRetrieveAdminAAAServerList
//
// Definition: administration_retrieveAdminAAAServerList
type WSGAdministrationRetrieveAdminAAAServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationRetrieveAdminAAAServerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationRetrieveAdminAAAServerListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationRetrieveAdminAAAServerList
}

func newWSGAdministrationRetrieveAdminAAAServerListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationRetrieveAdminAAAServerListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationRetrieveAdminAAAServerListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationRetrieveAdminAAAServerList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationRetrieveAdminAAAServerList() *WSGAdministrationRetrieveAdminAAAServerList {
	m := new(WSGAdministrationRetrieveAdminAAAServerList)
	return m
}

// WSGAdministrationRetrieveAdminAAAServerListType
//
// Definition: administration_retrieveAdminAAAServerListType
type WSGAdministrationRetrieveAdminAAAServerListType struct {
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty"`
}

func NewWSGAdministrationRetrieveAdminAAAServerListType() *WSGAdministrationRetrieveAdminAAAServerListType {
	m := new(WSGAdministrationRetrieveAdminAAAServerListType)
	return m
}

// WSGAdministrationScheduleBackup
//
// Definition: administration_scheduleBackup
type WSGAdministrationScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty"`

	// EnableScheduleBackup
	// enable schedule backup
	EnableScheduleBackup *bool `json:"enableScheduleBackup,omitempty"`

	// Hour
	// hour
	Hour *int `json:"hour,omitempty"`

	// Interval
	// schedule interval
	// Constraints:
	//    - oneof:[MONTHLY,WEEKLY,DAILY]
	Interval *string `json:"interval,omitempty"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type WSGAdministrationScheduleBackupAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationScheduleBackup
}

func newWSGAdministrationScheduleBackupAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationScheduleBackupAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationScheduleBackupAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationScheduleBackup)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationScheduleBackup() *WSGAdministrationScheduleBackup {
	m := new(WSGAdministrationScheduleBackup)
	return m
}

// WSGAdministrationSecondaryRadiusServer
//
// Definition: administration_secondaryRadiusServer
type WSGAdministrationSecondaryRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// MaxRetries
	// Max number(how many times) of retries for re-connection to primary
	// Constraints:
	//    - required
	//    - default:2
	//    - min:2
	//    - max:10
	MaxRetries *int `json:"maxRetries"`

	// Port
	// Port number of Secondary RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Protocol
	// Constraints:
	//    - default:'PAP'
	//    - oneof:[PAP,CHAP,PEAP]
	Protocol *string `json:"protocol,omitempty"`

	// RequestTimeOut
	// Request timeout(seconds) value of re-connection to primary
	// Constraints:
	//    - required
	//    - default:3
	//    - min:2
	//    - max:20
	RequestTimeOut *int `json:"requestTimeOut"`

	// RetryPriInvl
	// Interval of re-connection to primary(1-60 minute)
	// Constraints:
	//    - required
	//    - default:5
	//    - min:1
	//    - max:60
	RetryPriInvl *int `json:"retryPriInvl"`

	// SharedSecret
	// Shared secret of Secondary RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGAdministrationSecondaryRadiusServer() *WSGAdministrationSecondaryRadiusServer {
	m := new(WSGAdministrationSecondaryRadiusServer)
	return m
}

// WSGAdministrationTacacsServer
//
// Definition: administration_tacacsServer
type WSGAdministrationTacacsServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip"`

	// Port
	// Port number of TACACS+ Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port"`

	// Service
	// Constraints:
	//    - required
	Service *WSGCommonRealm `json:"service"`

	// SharedSecret
	// Shared secret of TACACS+ Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret"`
}

func NewWSGAdministrationTacacsServer() *WSGAdministrationTacacsServer {
	m := new(WSGAdministrationTacacsServer)
	return m
}

// WSGAdministrationUpgradeHistoryList
//
// Definition: administration_upgradeHistoryList
type WSGAdministrationUpgradeHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationUpgradeHistorySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationUpgradeHistoryListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationUpgradeHistoryList
}

func newWSGAdministrationUpgradeHistoryListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationUpgradeHistoryListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationUpgradeHistoryListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationUpgradeHistoryList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationUpgradeHistoryList() *WSGAdministrationUpgradeHistoryList {
	m := new(WSGAdministrationUpgradeHistoryList)
	return m
}

// WSGAdministrationUpgradeHistorySummary
//
// Definition: administration_upgradeHistorySummary
type WSGAdministrationUpgradeHistorySummary struct {
	// ApFwVersion
	// apFwVersion of the upgrade history
	ApFwVersion *string `json:"apFwVersion,omitempty"`

	// CbVersion
	// cbVersion of the upgrade history
	CbVersion *string `json:"cbVersion,omitempty"`

	// DpVersion
	// dpVersion of the upgrade history
	DpVersion *string `json:"dpVersion,omitempty"`

	// ElapsedSeconds
	// elapsedSeconds of the upgrade history
	ElapsedSeconds *int `json:"elapsedSeconds,omitempty"`

	// FileName
	// fileName of the upgrade history
	FileName *string `json:"fileName,omitempty"`

	// OldApFwVersion
	// oldApFwVersion of the upgrade history
	OldApFwVersion *string `json:"oldApFwVersion,omitempty"`

	// OldCbVersion
	// oldCbVersion of the upgrade history
	OldCbVersion *string `json:"oldCbVersion,omitempty"`

	// OldDpVersion
	// oldDpVersion of the upgrade history
	OldDpVersion *string `json:"oldDpVersion,omitempty"`

	// OldVersion
	// oldVersion of the upgrade history
	OldVersion *string `json:"oldVersion,omitempty"`

	// StartTime
	// startTime of the upgrade history
	StartTime *string `json:"startTime,omitempty"`

	// Version
	// version of the upgrade history
	Version *string `json:"version,omitempty"`
}

func NewWSGAdministrationUpgradeHistorySummary() *WSGAdministrationUpgradeHistorySummary {
	m := new(WSGAdministrationUpgradeHistorySummary)
	return m
}

// WSGAdministrationUpgradePatchInfo
//
// Definition: administration_upgradePatchInfo
type WSGAdministrationUpgradePatchInfo struct {
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`

	UploadPatchInfo *WSGClusterBladeUploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

type WSGAdministrationUpgradePatchInfoAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationUpgradePatchInfo
}

func newWSGAdministrationUpgradePatchInfoAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationUpgradePatchInfoAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationUpgradePatchInfoAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationUpgradePatchInfo)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationUpgradePatchInfo() *WSGAdministrationUpgradePatchInfo {
	m := new(WSGAdministrationUpgradePatchInfo)
	return m
}

// WSGAdministrationUpgradeStatus
//
// Definition: administration_upgradeStatus
type WSGAdministrationUpgradeStatus struct {
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`
}

type WSGAdministrationUpgradeStatusAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationUpgradeStatus
}

func newWSGAdministrationUpgradeStatusAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationUpgradeStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationUpgradeStatusAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationUpgradeStatus)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationUpgradeStatus() *WSGAdministrationUpgradeStatus {
	m := new(WSGAdministrationUpgradeStatus)
	return m
}

// WSGAdministrationZdAP
//
// Definition: administration_zdAP
type WSGAdministrationZdAP struct {
	// Connected
	// AP Conntected
	Connected *string `json:"connected,omitempty"`

	// Mac
	// AP MAC
	Mac *string `json:"mac,omitempty"`
}

func NewWSGAdministrationZdAP() *WSGAdministrationZdAP {
	m := new(WSGAdministrationZdAP)
	return m
}

// WSGAdministrationZdAPList
//
// Definition: administration_zdAPList
type WSGAdministrationZdAPList struct {
	Extra interface{} `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationZdAP `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationZdAPListAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationZdAPList
}

func newWSGAdministrationZdAPListAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationZdAPListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationZdAPListAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationZdAPList)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationZdAPList() *WSGAdministrationZdAPList {
	m := new(WSGAdministrationZdAPList)
	return m
}

// WSGAdministrationZdImport
//
// Definition: administration_zdImport
type WSGAdministrationZdImport struct {
	// ApMacList
	// List of AP MAC
	ApMacList []string `json:"apMacList,omitempty"`

	// Ip
	// ZD IP address
	Ip *string `json:"ip,omitempty"`

	// Password
	// ZD password
	Password *string `json:"password,omitempty"`

	// User
	// ZD user name
	User *string `json:"user,omitempty"`
}

func NewWSGAdministrationZdImport() *WSGAdministrationZdImport {
	m := new(WSGAdministrationZdImport)
	return m
}

// WSGAdministrationZdImportStatus
//
// Definition: administration_zdImportStatus
type WSGAdministrationZdImportStatus struct {
	// Details
	// Details
	Details *string `json:"details,omitempty"`

	// Message
	// Message
	Message *string `json:"message,omitempty"`

	// Progress
	// Progress
	Progress *int `json:"progress,omitempty"`

	// State
	// State
	State *string `json:"state,omitempty"`
}

type WSGAdministrationZdImportStatusAPIResponse struct {
	*RawAPIResponse
	Data *WSGAdministrationZdImportStatus
}

func newWSGAdministrationZdImportStatusAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGAdministrationZdImportStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGAdministrationZdImportStatusAPIResponse) Hydrate() error {
	r.Data = new(WSGAdministrationZdImportStatus)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGAdministrationZdImportStatus() *WSGAdministrationZdImportStatus {
	m := new(WSGAdministrationZdImportStatus)
	return m
}

// AddAdminaaa
//
// Operation ID: addAdminaaa
//
// Use this API command to create a new Admin AAA server
//
// Request Body:
//	 - body *WSGAdministrationCreateAdminAAAServer
func (s *WSGAdministrationService) AddAdminaaa(ctx context.Context, body *WSGAdministrationCreateAdminAAAServer, mutators ...RequestMutator) (*WSGCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddAdminaaa, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*WSGCommonCreateResultAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGCommonCreateResultAPIResponse), err
}

// AddRestart
//
// Operation ID: addRestart
//
// Use this API command to restart the controller.
func (s *WSGAdministrationService) AddRestart(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddRestart, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// AddShutdown
//
// Operation ID: addShutdown
//
// Use this API command to shut down the controller.
func (s *WSGAdministrationService) AddShutdown(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPost, RouteWSGAddShutdown, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// DeleteAdminaaaById
//
// Operation ID: deleteAdminaaaById
//
// Use this API command to delete an existing Admin AAA server
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAdministrationService) DeleteAdminaaaById(ctx context.Context, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteWSGDeleteAdminaaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindAdminaaa
//
// Operation ID: findAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
//
// Required Parameters:
// - type_ string
//		- required
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, type_ string, mutators ...RequestMutator) (*WSGAdministrationRetrieveAdminAAAServerListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationRetrieveAdminAAAServerListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationRetrieveAdminAAAServerListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAdminaaa, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.QueryParams.Set("type", type_)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationRetrieveAdminAAAServerListAPIResponse), err
}

// FindAdminaaaById
//
// Operation ID: findAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, id string, mutators ...RequestMutator) (*WSGAdministrationRetrieveAdminAAAServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationRetrieveAdminAAAServerAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationRetrieveAdminAAAServerAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindAdminaaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationRetrieveAdminAAAServerAPIResponse), err
}

// FindLicenses
//
// Operation ID: findLicenses
//
// Use this API command to get all licenses currently assign in SCG.
func (s *WSGAdministrationService) FindLicenses(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicensesListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicensesListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationLicensesListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindLicenses, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationLicensesListAPIResponse), err
}

// FindLicenseServer
//
// Operation ID: findLicenseServer
//
// Use this API command to get license server configuration.
func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicenseServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicenseServerAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationLicenseServerAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindLicenseServer, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationLicenseServerAPIResponse), err
}

// FindLicensesSummary
//
// Operation ID: findLicensesSummary
//
// Use this API command to get licenses summary information.
func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicensesSummaryListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicensesSummaryListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationLicensesSummaryListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindLicensesSummary, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationLicensesSummaryListAPIResponse), err
}

// FindLicensesSyncLogs
//
// Operation ID: findLicensesSyncLogs
//
// Use this API command to get licenses synchronize logs.
func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context, mutators ...RequestMutator) (*WSGAdministrationLicensesSyncLogsListAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGAdministrationLicensesSyncLogsListAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGAdministrationLicensesSyncLogsListAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindLicensesSyncLogs, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGAdministrationLicensesSyncLogsListAPIResponse), err
}

// UpdateAdminaaaById
//
// Operation ID: updateAdminaaaById
//
// Use this API command to modify an existing Admin AAA server
//
// Request Body:
//	 - body *WSGAdministrationModifyAdminAAAServer
//
// Required Parameters:
// - id string
//		- required
func (s *WSGAdministrationService) UpdateAdminaaaById(ctx context.Context, body *WSGAdministrationModifyAdminAAAServer, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateAdminaaaById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateLicenseServer
//
// Operation ID: updateLicenseServer
//
// Use this API command to update license server configuration.
//
// Request Body:
//	 - body *WSGAdministrationModfiyLicenseServer
func (s *WSGAdministrationService) UpdateLicenseServer(ctx context.Context, body *WSGAdministrationModfiyLicenseServer, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateLicenseServer, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// UpdateLicensesSync
//
// Operation ID: updateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateLicensesSync, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
