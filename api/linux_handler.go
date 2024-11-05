package api

import (
	"encoding/json"
	"github.com/esperer/rest-lic/linux/cmd/uptime"
	"net/http"
)

func UptimeLinuxHandler(w http.ResponseWriter, r *http.Request) {
	response := uptime.Uptime()
	json.NewEncoder(w).Encode(response)
}
