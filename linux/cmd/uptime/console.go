package uptime

type UptimeResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitemtpy"`
}
