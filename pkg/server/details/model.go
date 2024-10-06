package details

// ServerDetails struct for structured server details (optional)
type ServerDetails struct {
	Hostname      string `json:"hostname"`       // Hostname of the server
	OS            string `json:"os"`             // Operating system information
	KernelVersion string `json:"kernel_version"` // Linux kernel version
	//CPUModel          string `json:"cpu_model"`          // CPU model and details
	//CPULoad           string `json:"cpu_load"`           // Current CPU load
	//TotalMemory       string `json:"total_memory"`       // Total memory (RAM)
	//UsedMemory        string `json:"used_memory"`        // Used memory (RAM)
	//FreeMemory        string `json:"free_memory"`        // Free memory (RAM)
	//DiskUsage         string `json:"disk_usage"`         // Disk usage details
	//IPAddresses       string `json:"ip_addresses"`       // Server's IP addresses
	//NetworkInterfaces string `json:"network_interfaces"` // Network interfaces information
	//OpenPorts         string `json:"open_ports"`         // Open ports on the server
	Uptime string `json:"uptime"` // Server uptime
	//LoggedInUsers     string `json:"logged_in_users"`    // Currently logged in users
	//RunningProcesses  string `json:"running_processes"`  // Running processes on the server
	//FirewallRules     string `json:"firewall_rules"`     // Firewall rules
}
