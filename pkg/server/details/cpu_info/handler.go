package cpu_info

import (
	"strings"
)

func ParseCPUInfo(cpuInfoData string) ([]CPUInfo, error) {
	var cpuInfos []CPUInfo
	var cpuInfo CPUInfo

	lines := strings.Split(cpuInfoData, "\n")
	for _, line := range lines {
		if line == "" {
			// End of one processor info, add to slice and start a new one
			cpuInfos = append(cpuInfos, cpuInfo)
			cpuInfo = CPUInfo{}
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "processor":
			cpuInfo.Processor = value
		case "vendor_id":
			cpuInfo.VendorID = value
		case "cpu family":
			cpuInfo.CPUFamily = value
		case "model":
			cpuInfo.Model = value
		case "model name":
			cpuInfo.ModelName = value
		case "stepping":
			cpuInfo.Stepping = value
		case "microcode":
			cpuInfo.Microcode = value
		case "cpu MHz":
			cpuInfo.CPUMHz = value
		case "cache size":
			cpuInfo.CacheSize = value
		case "physical id":
			cpuInfo.PhysicalID = value
		case "siblings":
			cpuInfo.Siblings = value
		case "core id":
			cpuInfo.CoreID = value
		case "cpu cores":
			cpuInfo.CPUCores = value
		case "apicid":
			cpuInfo.APICID = value
		case "initial apicid":
			cpuInfo.InitialAPICID = value
		case "fpu":
			cpuInfo.FPU = value
		case "fpu_exception":
			cpuInfo.FPUException = value
		case "cpuid level":
			cpuInfo.CPUIDLevel = value
		case "wp":
			cpuInfo.WP = value
		case "flags":
			cpuInfo.Flags = value
		case "bogomips":
			cpuInfo.Bogomips = value
		case "clflush size":
			cpuInfo.ClflushSize = value
		case "cache_alignment":
			cpuInfo.CacheAlignment = value
		case "address sizes":
			cpuInfo.AddressSizes = value
		case "power management":
			cpuInfo.PowerManagement = value
		}
	}
	return cpuInfos, nil
}
