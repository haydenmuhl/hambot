package admin

import (
	"net/http"
	"time"

	"github.com/haydenmuhl/hambot/chat"
)

func init() {
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/on", handleOn)
	mux.HandleFunc("/bootstrap/setup", initialSetupForm)
	mux.HandleFunc("/bootstrap/submit", processSetupForm)
}

func handleRoot(r http.ResponseWriter, req *http.Request) {
	if req.URL.String() != "/" {
		r.WriteHeader(http.StatusNotFound)
		r.Write([]byte("404 Not Found"))
		return
	}
	r.Write([]byte(`<a href="/on">Turn on</a>`))
	return
}

func handleOn(r http.ResponseWriter, req *http.Request) {
	if !chat.HasCredentials() {
		r.Header().Set("Location", "/bootstrap/setup")
		r.WriteHeader(http.StatusFound)
		return
	}

	go chat.Client().Run()
	r.WriteHeader(http.StatusOK)
	r.Write([]byte("<p>Chat bot started</p>"))
	return
}

var mux http.ServeMux

var Server = http.Server{
	Addr:              ":8080",
	Handler:           &mux,
	ReadHeaderTimeout: 5 * time.Second,
	WriteTimeout:      5 * time.Second,
}
