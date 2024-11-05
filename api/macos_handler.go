package api

import (
	"encoding/json"
	"github.com/esperer/rest-lic/macos/cmd/vm_stat"
	"net/http"
)

func VmStatMacosHandler(w http.ResponseWriter, r *http.Request) {
	response, err := vm_stat.Vmstat()
	if err != nil {
		http.Error(w, "Failed to retrieve memory info vmstats: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(response)
}
