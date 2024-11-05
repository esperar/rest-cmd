package uptime

import (
	"os/exec"
	"strings"
)

func Uptime() UptimeResponse {
	cmd := exec.Command("uptime")
	output, err := cmd.CombinedOutput()
	response := UptimeResponse{}

	if err != nil {
		response.Error = err.Error()
	} else {
		response.Output = strings.TrimSpace(string(output))
	}

	return response
}
