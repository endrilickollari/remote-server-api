package running_processes

import "strings"

func ParseDiskUsage(dfOutput string) ([]RunningProcesses, error) {
	var runningProcesses []RunningProcesses
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

		diskUsage := RunningProcesses{
			User:  fields[0],
			PID:   fields[1],
			CPU:   fields[2],
			VSZ:   fields[3],
			RSS:   fields[4],
			TTY:   fields[5],
			Stat:  fields[6],
			Start: fields[7],
			Time:  fields[8],
			CMD:   fields[9],
		}

		runningProcesses = append(runningProcesses, diskUsage)
	}

	return runningProcesses, nil
}
