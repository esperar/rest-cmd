package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Run() {
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", OkHandler).Methods("GET")
	// linux
	router.HandleFunc("/linux/uptime", UptimeLinuxHandler).Methods("GET")
	router.HandleFunc("/linux/free", FreeLinuxHandler).Methods("GET")
	//macos
	router.HandleFunc("/macos/vm_stat", VmStatMacosHandler).Methods("GET")
	fmt.Println("Server Start! :3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}
