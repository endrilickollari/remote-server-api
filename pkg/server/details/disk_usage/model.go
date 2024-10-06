package disk_usage

type DiskUsage struct {
	Filesystem    string `json:"filesystem"`
	Size          string `json:"size"`
	Used          string `json:"used"`
	Available     string `json:"available"`
	UsePercentage string `json:"use_percentage"`
	MountedOn     string `json:"mounted_on"`
}
