package disk_usage

import "strings"

func ParseDiskUsage(dfOutput string) ([]DiskUsage, error) {
	var diskUsages []DiskUsage
	lines := strings.Split(dfOutput, "\n")

	// Skip the first line (headers)
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		// Split the line by spaces
		fields := strings.Fields(line)
		if len(fields) < 6 {
			continue
		}

		diskUsage := DiskUsage{
			Filesystem:    fields[0],
			Size:          fields[1],
			Used:          fields[2],
			Available:     fields[3],
			UsePercentage: fields[4],
			MountedOn:     fields[5],
		}

		diskUsages = append(diskUsages, diskUsage)
	}

	return diskUsages, nil
}
