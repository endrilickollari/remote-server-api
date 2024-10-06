package docker

import "strings"

func ParseDockerContainers(dockerOutput string) ([]DockerContainer, error) {
	var containers []DockerContainer
	lines := strings.Split(dockerOutput, "\n")

	// Skip the header line and empty lines
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		// Split the line by spaces, accounting for variable space in fields
		fields := strings.Fields(line)

		// Check if we have at least 6 fields (the last two fields can be combined)
		if len(fields) < 6 {
			continue
		}

		containerID := fields[0]
		image := fields[1]

		command := fields[2]
		createdOn := fields[3] + " " + fields[4] + " " + fields[5]
		status := fields[6] + " " + fields[7] + " " + fields[8]
		ports := fields[9]
		names := fields[10]

		container := DockerContainer{
			ContainerID: containerID,
			Image:       image,
			Command:     command,
			CreatedOn:   createdOn,
			Status:      status,
			Ports:       ports,
			Names:       names,
		}

		containers = append(containers, container)
	}

	return containers, nil
}
