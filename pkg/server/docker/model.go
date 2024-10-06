package docker

type DockerContainer struct {
	ContainerID string `json:"container_id"`
	Image       string `json:"image"`
	Command     string `json:"command"`
	CreatedOn   string `json:"created_on"`
	Status      string `json:"status"`
	Ports       string `json:"ports"`
	Names       string `json:"names"`
}
