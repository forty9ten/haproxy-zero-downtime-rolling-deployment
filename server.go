package main

import (
	"flag"
	"fmt"
	"os"

	"net/http"
)

var (
	healthState = "UP"
	bindPort    string
)

func main() {
	var port = flag.String("port", "4910", "port")
	flag.Parse()

	bindPort = *port

	http.HandleFunc("/v1/health", health)
	http.HandleFunc("/v1/ping", ping)

	// toggling endpoints for testing only
	http.HandleFunc("/v1/down", down)
	http.HandleFunc("/v1/up", up)
	http.HandleFunc("/v1/pause", paused)

	// similate work being done
	http.HandleFunc("/v1/dowork", dowork)

	fmt.Printf("Server up and running on port %s\n", bindPort)
	http.ListenAndServe(fmt.Sprintf(":%s", bindPort), nil)
}

func ping(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(200)
}

func health(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(healthState))
	w.WriteHeader(200)
}

func dowork(w http.ResponseWriter, req *http.Request) {
	switch {
	case healthState == "UP":
		w.Write([]byte("DOING SOME WORK.  RUNNING ON PORT: " + bindPort))
		w.WriteHeader(200)
	case healthState == "PAUSE":
		http.Error(w, "PAUSED", 503)
	}
}

func down(w http.ResponseWriter, req *http.Request) {
	os.Exit(-1)
}

func up(w http.ResponseWriter, req *http.Request) {
	healthState = "UP"
}

func paused(w http.ResponseWriter, req *http.Request) {
	healthState = "PAUSE"
}
