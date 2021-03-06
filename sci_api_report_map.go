package bigdog

// API Version: 1.0.0

type (
	SCIReportSectionMetadata struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	}

	SCIReportMetadata struct {
		ID       int                        `json:"id"`
		Title    string                     `json:"title"`
		Sections []SCIReportSectionMetadata `json:"sections"`
	}
)

// SCIReportMetadata returns some basic information parsed from the SCI about what reports provide which sections
func SCIReportMetadataDefinitions() []SCIReportMetadata {
	return []SCIReportMetadata{
		{
			ID:    1,
			Title: "Wireless - Clients Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 15, Title: "Unique Clients Trend Over Time"},
				{ID: 16, Title: "Unique Clients Trend Over Time"},
				{ID: 17, Title: "Top Clients by Traffic Percentile"},
				{ID: 18, Title: "Top 10 OS by Client Count"},
				{ID: 19, Title: "Top 10 Manufacturers by Client Count"},
				{ID: 13, Title: "Top 10 Unique Clients by Traffic"},
				{ID: 112, Title: "Top 10 Authentication Methods by Client Count"},
				{ID: 12, Title: "Overview"},
				{ID: 14, Title: "Clients Details"},
			},
		},
		{
			ID:    2,
			Title: "Network - Wireless Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 21, Title: "Traffic Distribution"},
				{ID: 23, Title: "Traffic Over Time"},
				{ID: 26, Title: "Top APs by Traffic"},
				{ID: 25, Title: "Top APs by Client Count"},
				{ID: 22, Title: "Traffic Trend"},
				{ID: 20, Title: "Overview"},
				{ID: 24, Title: "Top APs by Traffic"},
				{ID: 27, Title: "Top APs by Client Count"},
			},
		},
		{
			ID:    3,
			Title: "Wireless - Applications Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 8, Title: "Top Applications by Traffic"},
				{ID: 7, Title: "Top Applications by Traffic"},
				{ID: 10, Title: "Overview"},
				{ID: 9, Title: "Top Applications by Client Count"},
				{ID: 11, Title: "Top Applications by Client Count"},
			},
		},
		{
			ID:    4,
			Title: "WLANs Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 35, Title: "Overview"},
				{ID: 40, Title: "Top SSIDs by Traffic"},
				{ID: 39, Title: "SSID Changes Over Time"},
				{ID: 36, Title: "Top SSIDs by Traffic"},
				{ID: 38, Title: "Top SSIDs by Client Count"},
				{ID: 37, Title: "Active SSIDs Trend"},
				{ID: 41, Title: "Top SSIDs by Client Count"},
			},
		},
		{
			ID:    5,
			Title: "Airtime Utilization Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 1, Title: "Overview"},
				{ID: 4, Title: "Top APs by Airtime Utilization for 5 GHz"},
				{ID: 2, Title: "Top 10 APs by Airtime Utilization"},
				{ID: 6, Title: "Airtime Utilization Over Time"},
				{ID: 5, Title: "Airtime Utilization Trend"},
				{ID: 3, Title: "Top APs by Airtime Utilization for 2.4 GHz"},
			},
		},
		{
			ID:    6,
			Title: "Sessions Summary Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 34, Title: "Average Durations"},
				{ID: 42, Title: "Session Duration Percentiles"},
				{ID: 33, Title: "Top Sessions Summary"},
			},
		},
		{
			ID:    8,
			Title: "APs Reboot Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 43, Title: "Total Reboots"},
				{ID: 45, Title: "AP Reboots"},
				{ID: 44, Title: "Top AP Reboots"},
			},
		},
		{
			ID:    9,
			Title: "Inventory - APs Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 46, Title: "Overview"},
				{ID: 57, Title: "Top APs by Offline Duration"},
				{ID: 51, Title: "Top AP Software Versions"},
				{ID: 58, Title: "Top APs by Reboot Count"},
				{ID: 56, Title: "Top AP Software Versions"},
				{ID: 53, Title: "Top APs by Reboot Count"},
				{ID: 47, Title: "Top APs by Offline Duration"},
				{ID: 55, Title: "Top AP Models"},
				{ID: 49, Title: "AP Status Trends"},
				{ID: 48, Title: "AP Count Trend"},
				{ID: 50, Title: "Top AP Models"},
				{ID: 54, Title: "Top 10 AP Alarm Types"},
				{ID: 60, Title: "AP Details for Online/Offline Status"},
				{ID: 61, Title: "AP Details for Other Statuses"},
				{ID: 52, Title: "Top 10 AP Reboot Reasons"},
				{ID: 59, Title: "APs Configured in Multiple Systems"},
			},
		},
		{
			ID:    10,
			Title: "Overview",
			Sections: []SCIReportSectionMetadata{
				{ID: 63, Title: "Controllers"},
				{ID: 64, Title: "Access Points"},
				{ID: 115, Title: "Network Usage Overview"},
				{ID: 62, Title: "Ruckus SmartAnalytics"},
				{ID: 68, Title: "Top APs by Client Count"},
				{ID: 66, Title: "Alarms"},
				{ID: 71, Title: "WLANs"},
				{ID: 69, Title: "Total Wireless Traffic"},
				{ID: 73, Title: "Applications (Wireless)"},
				{ID: 70, Title: "Unique Wireless Sessions"},
				{ID: 72, Title: "Radios"},
				{ID: 67, Title: "Switches"},
				{ID: 74, Title: "Events"},
				{ID: 97, Title: "Did you know?"},
			},
		},
		{
			ID:    11,
			Title: "AP Details Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 77, Title: "Details"},
				{ID: 76, Title: "Performance"},
				{ID: 22, Title: "Traffic Trend"},
				{ID: 78, Title: "Stats"},
				{ID: 75, Title: "Summary"},
				{ID: 79, Title: "Uptime History"},
				{ID: 15, Title: "Unique Clients Trend Over Time"},
				{ID: 80, Title: "Top 10 Clients by Traffic Volume"},
				{ID: 7, Title: "Top Applications by Traffic"},
				{ID: 8, Title: "Top Applications by Traffic"},
				{ID: 40, Title: "Top SSIDs by Traffic"},
				{ID: 82, Title: "RSS Trend"},
				{ID: 81, Title: "Sessions"},
				{ID: 85, Title: "Events"},
				{ID: 95, Title: "Anomalies"},
				{ID: 110, Title: "Anomalies for the Past 30 Days"},
				{ID: 5, Title: "Airtime Utilization Trend"},
				{ID: 83, Title: "SNR Trend"},
				{ID: 14, Title: "Clients Details"},
				{ID: 84, Title: "Alarms"},
			},
		},
		{
			ID:    12,
			Title: "Client Details Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 83, Title: "SNR Trend"},
				{ID: 92, Title: "Sessions"},
				{ID: 82, Title: "RSS Trend"},
				{ID: 86, Title: "Summary"},
				{ID: 8, Title: "Top Applications by Traffic"},
				{ID: 89, Title: "Traffic Trend"},
				{ID: 87, Title: "Stats"},
				{ID: 7, Title: "Top Applications by Traffic"},
			},
		},
		{
			ID:    13,
			Title: "SC Network Traffic Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 103, Title: "SmartCell Average Throughput"},
				{ID: 107, Title: "SmartCell Traffic Volume"},
				{ID: 109, Title: "SmartCell Frequency Sync Loss"},
				{ID: 111, Title: "RSC Traffic Statistics"},
				{ID: 94, Title: "SmartCell Trend Over Time"},
				{ID: 101, Title: "SmartCell Connection Setup Success Rate"},
				{ID: 105, Title: "SmartCell Connection Statistics"},
				{ID: 100, Title: "SmartCell Dropped Call Rate"},
				{ID: 102, Title: "SmartCell Handover Success Rate"},
				{ID: 108, Title: "SmartCell Phase Sync Loss"},
				{ID: 106, Title: "RSC GPS Statistics"},
				{ID: 93, Title: "Summary"},
				{ID: 104, Title: "SmartCell Availability"},
			},
		},
		{
			ID:       14,
			Title:    "Data Explorer",
			Sections: []SCIReportSectionMetadata{},
		},
		{
			ID:    15,
			Title: "Inventory - Controllers Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 114, Title: "Overview"},
				{ID: 99, Title: "License Utilization"},
				{ID: 96, Title: "KRACK Assessment"},
				{ID: 98, Title: "Resource Utilization"},
				{ID: 151, Title: "License Utilization Over Time"},
			},
		},
		{
			ID:    16,
			Title: "Inventory - Switches Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 113, Title: "Overview"},
				{ID: 121, Title: "Top Switch Models"},
				{ID: 117, Title: "Top Switch Software Versions"},
				{ID: 116, Title: "Switch Count Trend"},
				{ID: 118, Title: "Top Switch Software Versions"},
				{ID: 122, Title: "Top Switch Models"},
				{ID: 132, Title: "Port Status Trends"},
			},
		},
		{
			ID:    17,
			Title: "Network - Wired Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 127, Title: "Top Switches by Traffic"},
				{ID: 134, Title: "Overview"},
				{ID: 128, Title: "Top Switches by Traffic"},
				{ID: 142, Title: "Top Switches by Errors"},
				{ID: 136, Title: "Traffic Trend"},
				{ID: 135, Title: "Traffic Distribution by Switch Model and Port Speed"},
				{ID: 143, Title: "Top Switches by Errors"},
				{ID: 123, Title: "Top Switches by PoE Usage"},
				{ID: 124, Title: "Top Switches by PoE Usage"},
				{ID: 141, Title: "Error Trend"},
			},
		},
		{
			ID:    18,
			Title: "Switch Details Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 130, Title: "Top Ports by Traffic"},
				{ID: 138, Title: "Uptime History"},
				{ID: 139, Title: "Details"},
				{ID: 129, Title: "Top Ports by Traffic"},
				{ID: 125, Title: "Summary"},
				{ID: 126, Title: "Resource Utilization"},
				{ID: 137, Title: "LLDP Neighbor List"},
				{ID: 152, Title: "Error Trend"},
				{ID: 131, Title: "Traffic Trend"},
			},
		},
		{
			ID:    19,
			Title: "Comparison Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 140, Title: "Overview"},
				{ID: 146, Title: "Metric 2 Over Time"},
				{ID: 145, Title: "Metric 1 Over Time"},
				{ID: 147, Title: "Table"},
			},
		},
		{
			ID:    20,
			Title: "Client Health Report",
			Sections: []SCIReportSectionMetadata{
				{ID: 144, Title: "Summary"},
				{ID: 150, Title: "Health By Group"},
				{ID: 149, Title: "Health Metric Trends"},
				{ID: 148, Title: "Client Connection Health"},
			},
		},
	}
}
