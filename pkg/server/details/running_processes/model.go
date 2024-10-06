package running_processes

type RunningProcesses struct {
	User  string `json:"user"`
	PID   string `json:"process_id"`
	CPU   string `json:"cpu_consumption"`
	VSZ   string `json:"vsz"`
	RSS   string `json:"rss"`
	TTY   string `json:"tty"`
	Stat  string `json:"stat"`
	Start string `json:"started"`
	Time  string `json:"time"`
	CMD   string `json:"command"`
}
