package details

// ServerDetails struct for structured server details (optional)
type ServerDetails struct {
	Hostname string `json:"hostname"`
	Status   string `json:"status"`
	Uptime   string `json:"uptime"`
}
