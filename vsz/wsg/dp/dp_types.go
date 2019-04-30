package dp

// API Version: v8_0

type EmptyResult struct {
	DpEmptyResult *string `json:"dp_emptyResult,omitempty"`
}

type SwitchoverDp struct {
	// ClusterName
	// Name of destination cluster, Notice: System will take ipOrFqdn as 1st priority, and clusterName as 2nd.
	ClusterName *string `json:"clusterName,omitempty"`

	// DeleteRecord
	// Flag to delete DP record after switchover cluster. Default value is false.
	DeleteRecord *bool `json:"deleteRecord,omitempty"`

	// DpIDList
	// DP ID list
	DpIDList []string `json:"dpIdList,omitempty"`

	// IPOrFqdn
	// IP or FQDN address of destination cluster, Notice: System will take ipOrFqdn as 1st priority, and
	// clusterName as 2nd.
	IPOrFqdn *string `json:"ipOrFqdn,omitempty"`
}
