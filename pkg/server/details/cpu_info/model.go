package cpu_info

// CPUInfo struct for CPU details
type CPUInfo struct {
	Processor       string `json:"processor"`
	VendorID        string `json:"vendor_id"`
	CPUFamily       string `json:"cpu_family"`
	Model           string `json:"model"`
	ModelName       string `json:"model_name"`
	Stepping        string `json:"stepping"`
	Microcode       string `json:"microcode"`
	CPUMHz          string `json:"cpu_mhz"`
	CacheSize       string `json:"cache_size"`
	PhysicalID      string `json:"physical_id"`
	Siblings        string `json:"siblings"`
	CoreID          string `json:"core_id"`
	CPUCores        string `json:"cpu_cores"`
	APICID          string `json:"apicid"`
	InitialAPICID   string `json:"initial_apicid"`
	FPU             string `json:"fpu"`
	FPUException    string `json:"fpu_exception"`
	CPUIDLevel      string `json:"cpuid_level"`
	WP              string `json:"wp"`
	Flags           string `json:"flags"`
	Bogomips        string `json:"bogomips"`
	ClflushSize     string `json:"clflush_size"`
	CacheAlignment  string `json:"cache_alignment"`
	AddressSizes    string `json:"address_sizes"`
	PowerManagement string `json:"power_management"`
}
