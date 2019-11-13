package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAdministrationService struct {
	apiClient *APIClient
}

func NewWSGAdministrationService(c *APIClient) *WSGAdministrationService {
	s := new(WSGAdministrationService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdministrationService() *WSGAdministrationService {
	return NewWSGAdministrationService(ss.apiClient)
}

type WSGAdministrationActiveDirectoryServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of Active Directory Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

	// WindowsDomainName
	// Windows Domain Name of Active Directory Server object
	// Constraints:
	//    - required
	WindowsDomainName *string `json:"windowsDomainName" validate:"required"`
}

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

type WSGAdministrationApPatchHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationApPatchHistory `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

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

type WSGAdministrationApPatchStatus struct {
	ClusterOperationProgress *WSGClusterBladeClusterOperationProgress `json:"clusterOperationProgress,omitempty"`
}

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

type WSGAdministrationApplicationLogAndStatusList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationApplicationLogAndStatus `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// FTP server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

type WSGAdministrationBackupFile struct {
	// BackupElapsed
	// backup elapsed of the configuration backup file
	BackupElapsed *float64 `json:"backupElapsed,omitempty"`

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
	FileSize *float64 `json:"fileSize,omitempty"`

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

type WSGAdministrationClusterBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationClusterBackupSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

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

type WSGAdministrationConfigurationBackupList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationBackupFile `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationCreateAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

type WSGAdministrationDefaultRoleMapping struct {
	// DefaultAdmin
	// DefaultAdmin of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultAdmin *string `json:"defaultAdmin" validate:"required"`

	// DefaultUserGroup
	// DefaultUserGroup of DefaultRoleMapping object
	// Constraints:
	//    - required
	DefaultUserGroup *string `json:"defaultUserGroup" validate:"required"`
}

type WSGAdministrationLdapServer struct {
	// AdminDomainName
	// Admin Domain Name of LDAP Server object
	// Constraints:
	//    - required
	AdminDomainName *string `json:"adminDomainName" validate:"required"`

	// AdminPassword
	// Admin password of LDAP Server object
	// Constraints:
	//    - required
	AdminPassword *string `json:"adminPassword" validate:"required"`

	// BaseDomainName
	// Base Domain Name of LDAP Server object
	// Constraints:
	//    - required
	BaseDomainName *string `json:"baseDomainName" validate:"required"`

	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// KeyAttribute
	// Key attribute of LDAP Server object
	// Constraints:
	//    - required
	KeyAttribute *string `json:"keyAttribute" validate:"required"`

	// Port
	// Port number of LDAP Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

	// SearchFilter
	// Search filter of LDAP Server object
	// Constraints:
	//    - required
	SearchFilter *string `json:"searchFilter" validate:"required"`
}

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

type WSGAdministrationLicenseServer struct {
	// IpAddress
	// local license server IP address
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// local license server port
	// Constraints:
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=0,lte=65535"`

	// UseCloud
	// use cloud license server
	UseCloud *bool `json:"useCloud,omitempty"`
}

type WSGAdministrationLicensesList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicenses `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesSummary struct {
	CapacityControlLicenseCount *WSGAdministrationLicensesSummaryCapacityControlLicenseCountType `json:"capacityControlLicenseCount,omitempty"`

	// LicenseTypeDescription
	// license type description
	LicenseTypeDescription *string `json:"licenseTypeDescription,omitempty"`
}

type WSGAdministrationLicensesSummaryCapacityControlLicenseCountType struct {
	// TotalCount
	// total count of licenses
	TotalCount *int `json:"totalCount,omitempty"`

	// UsedCount
	// consumed count of licenses
	UsedCount *int `json:"usedCount,omitempty"`
}

type WSGAdministrationLicensesSummaryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicensesSummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationLicensesSyncLogs struct {
	// CreateDateTime
	// license sync log's create time
	CreateDateTime *string `json:"createDateTime,omitempty"`

	// SyncResult
	// sync license result
	// Constraints:
	//    - oneof:[SUCCESS,FAILURE]
	SyncResult *string `json:"syncResult,omitempty" validate:"oneof=SUCCESS FAILURE"`
}

type WSGAdministrationLicensesSyncLogsList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationLicensesSyncLogs `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationModfiyLicenseServer struct {
	IpAddress *string `json:"ipAddress,omitempty"`

	// Port
	// Constraints:
	//    - min:0
	//    - max:65535
	Port *int `json:"port,omitempty" validate:"gte=0,lte=65535"`

	// UseCloud
	// Constraints:
	//    - required
	UseCloud *bool `json:"useCloud" validate:"required"`
}

type WSGAdministrationModifyAdminAAAServer struct {
	ActiveDirectoryServer *WSGAdministrationActiveDirectoryServer `json:"activeDirectoryServer,omitempty"`

	DefaultRoleMapping *WSGAdministrationDefaultRoleMapping `json:"defaultRoleMapping,omitempty"`

	LdapServer *WSGAdministrationLdapServer `json:"ldapServer,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	RadiusServer *WSGAdministrationRadiusServer `json:"radiusServer,omitempty"`

	TacacsServer *WSGAdministrationTacacsServer `json:"tacacsServer,omitempty"`

	// Type
	// Specify the type(RADIUS/TACACS/AD/LDAP) of this Admin AAA Server, please be infomed that the type name [TACACS] is for TACACS+ (Terminal Access Controller Access-Control System Plus)
	// Constraints:
	//    - required
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type" validate:"required,oneof=RADIUS TACACS AD LDAP"`
}

type WSGAdministrationModifyAutoExportBackup struct {
	// EnableAutoExportBackup
	// enable auto export backup
	EnableAutoExportBackup *bool `json:"enableAutoExportBackup,omitempty"`

	// FtpServer
	// ftp server name
	FtpServer *string `json:"ftpServer,omitempty"`
}

type WSGAdministrationModifyLogLevel struct {
	// ApplicationName
	// Application name.
	ApplicationName *string `json:"applicationName,omitempty"`

	// LogLevel
	// Log level.
	// Constraints:
	//    - oneof:[DEBUG,INFO,WARN,ERROR]
	LogLevel *string `json:"logLevel,omitempty" validate:"oneof=DEBUG INFO WARN ERROR"`
}

type WSGAdministrationModifyScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

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
	Interval *string `json:"interval,omitempty" validate:"oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type WSGAdministrationRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Realm
	// Constraints:
	//    - required
	Realm *WSGCommonRealm `json:"realm" validate:"required,max=255"`

	SecondaryRadiusServer *WSGAdministrationSecondaryRadiusServer `json:"secondaryRadiusServer,omitempty"`

	// SharedSecret
	// Shared secret of RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

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
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS TACACS AD LDAP"`
}

type WSGAdministrationRetrieveAdminAAAServerList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationRetrieveAdminAAAServerListType `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGAdministrationRetrieveAdminAAAServerListType struct {
	Id *string `json:"id,omitempty"`

	Name *WSGCommonNormalName `json:"name,omitempty"`

	// Type
	// Constraints:
	//    - oneof:[RADIUS,TACACS,AD,LDAP]
	Type *string `json:"type,omitempty" validate:"oneof=RADIUS TACACS AD LDAP"`
}

type WSGAdministrationScheduleBackup struct {
	// DateOfMonth
	// date of the month
	DateOfMonth *int `json:"dateOfMonth,omitempty"`

	// DayOfWeek
	// day of the week
	// Constraints:
	//    - oneof:[SUNDAY,MONDAY,TUESDAY,WEDNESDAY,THURSDAY,FRIDAY,SATURDAY]
	DayOfWeek *string `json:"dayOfWeek,omitempty" validate:"oneof=SUNDAY MONDAY TUESDAY WEDNESDAY THURSDAY FRIDAY SATURDAY"`

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
	Interval *string `json:"interval,omitempty" validate:"oneof=MONTHLY WEEKLY DAILY"`

	// Minute
	// minute
	Minute *int `json:"minute,omitempty"`
}

type WSGAdministrationSecondaryRadiusServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// MaxRetries
	// Max number(how many times) of retries for re-connection to primary
	// Constraints:
	//    - required
	//    - default:2
	//    - min:2
	//    - max:10
	MaxRetries *int `json:"maxRetries" validate:"required,gte=2,lte=10"`

	// Port
	// Port number of Secondary RADIUS Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// RequestTimeOut
	// Request timeout(seconds) value of re-connection to primary
	// Constraints:
	//    - required
	//    - default:3
	//    - min:2
	//    - max:20
	RequestTimeOut *int `json:"requestTimeOut" validate:"required,gte=2,lte=20"`

	// RetryPriInvl
	// Interval of re-connection to primary(1-60 minute)
	// Constraints:
	//    - required
	//    - default:5
	//    - min:1
	//    - max:60
	RetryPriInvl *int `json:"retryPriInvl" validate:"required,gte=1,lte=60"`

	// SharedSecret
	// Shared secret of Secondary RADIUS Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

type WSGAdministrationTacacsServer struct {
	// Ip
	// Constraints:
	//    - required
	Ip *WSGCommonIpAddress `json:"ip" validate:"required"`

	// Port
	// Port number of TACACS+ Server object
	// Constraints:
	//    - required
	//    - min:1
	//    - max:65535
	Port *int `json:"port" validate:"required,gte=1,lte=65535"`

	// Service
	// Constraints:
	//    - required
	Service *WSGCommonRealm `json:"service" validate:"required,max=255"`

	// SharedSecret
	// Shared secret of TACACS+ Server object
	// Constraints:
	//    - required
	SharedSecret *string `json:"sharedSecret" validate:"required"`
}

type WSGAdministrationUpgradeHistoryList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationUpgradeHistorySummary `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

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
	ElapsedSeconds *float64 `json:"elapsedSeconds,omitempty"`

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

type WSGAdministrationUpgradePatchInfo struct {
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`

	UploadPatchInfo *WSGClusterBladeUploadPatchInfo `json:"uploadPatchInfo,omitempty"`
}

type WSGAdministrationUpgradeStatus struct {
	ClusterOperationProgress *WSGClusterBladeClusterUpgradeProgress `json:"clusterOperationProgress,omitempty"`
}

type WSGAdministrationZdAP struct {
	// Connected
	// AP Conntected
	Connected *string `json:"connected,omitempty"`

	// Mac
	// AP MAC
	Mac *string `json:"mac,omitempty"`
}

type WSGAdministrationZdAPList struct {
	Extra *WSGCommonRbacMetadata `json:"extra,omitempty"`

	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdministrationZdAP `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

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

// AddAdminaaa
//
// Use this API command to create a new Admin AAA server
//
// Request Body:
//	 - body *WSGAdministrationCreateAdminAAAServer
func (s *WSGAdministrationService) AddAdminaaa(ctx context.Context, body *WSGAdministrationCreateAdminAAAServer) (*WSGCommonCreateResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddRestart
//
// Use this API command to restart the controller.
func (s *WSGAdministrationService) AddRestart(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// AddShutdown
//
// Use this API command to shut down the controller.
func (s *WSGAdministrationService) AddShutdown(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// DeleteAdminaaaById
//
// Use this API command to delete an existing Admin AAA server
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) DeleteAdminaaaById(ctx context.Context, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAdminaaa
//
// Use this API command to retrieve the list of Admin AAA server
//
// Query Parameters:
// - qType string
//		- required
func (s *WSGAdministrationService) FindAdminaaa(ctx context.Context, qType string) (*WSGAdministrationRetrieveAdminAAAServerList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindAdminaaaById
//
// Use this API command to retrieve an existing Admin AAA server
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) FindAdminaaaById(ctx context.Context, pId string) (*WSGAdministrationRetrieveAdminAAAServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicenses
//
// Use this API command to get all licenses currently assign in SCG.
func (s *WSGAdministrationService) FindLicenses(ctx context.Context) (*WSGAdministrationLicensesList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicenseServer
//
// Use this API command to get license server configuration.
func (s *WSGAdministrationService) FindLicenseServer(ctx context.Context) (*WSGAdministrationLicenseServer, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicensesSummary
//
// Use this API command to get licenses summary information.
func (s *WSGAdministrationService) FindLicensesSummary(ctx context.Context) (*WSGAdministrationLicensesSummaryList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindLicensesSyncLogs
//
// Use this API command to get licenses synchronize logs.
func (s *WSGAdministrationService) FindLicensesSyncLogs(ctx context.Context) (*WSGAdministrationLicensesSyncLogsList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateAdminaaaById
//
// Use this API command to modify an existing Admin AAA server
//
// Request Body:
//	 - body *WSGAdministrationModifyAdminAAAServer
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGAdministrationService) UpdateAdminaaaById(ctx context.Context, body *WSGAdministrationModifyAdminAAAServer, pId string) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateLicenseServer
//
// Use this API command to update license server configuration.
//
// Request Body:
//	 - body *WSGAdministrationModfiyLicenseServer
func (s *WSGAdministrationService) UpdateLicenseServer(ctx context.Context, body *WSGAdministrationModfiyLicenseServer) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}

// UpdateLicensesSync
//
// Use this API command to ask all SCG in cluster to sync licenses from license server.
func (s *WSGAdministrationService) UpdateLicensesSync(ctx context.Context) error {
	if ctx == nil {
		return errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("provided context is done: %s", err)
	}
}