package api

import (
	"encoding/json"
	"net/http"

	"github.com/esperer/rest-lic/linux/cmd/df"
	"github.com/esperer/rest-lic/linux/cmd/free"
	"github.com/esperer/rest-lic/linux/cmd/uptime"
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

func DfLinuxHandler(w http.ResponseWriter, r *http.Request) {
	response, err := df.Df()
	if err != nil {
		http.Error(w, "Failed to retrieve disk info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
