package api

import (
	"encoding/json"
	"github.com/esperer/rest-lic/linux/cmd/free"
	"github.com/esperer/rest-lic/linux/cmd/uptime"
	"net/http"
)

func UptimeLinuxHandler(w http.ResponseWriter, r *http.Request) {
	response := uptime.Uptime()
	json.NewEncoder(w).Encode(response)
}

func FreeLinuxHandler(w http.ResponseWriter, r *http.Request) {
	response, err := free.Free()
	if err != nil {
		http.Error(w, "Failed to retrieve memory info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
