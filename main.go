package main

import (
	"log"
	"net/http"

	"github.com/haydenmuhl/hambot/chat"
	"github.com/haydenmuhl/hambot/database"
)

func main() {
	var err error

	err = database.Init()
	if err != nil {
		log.Fatalln(err)
	}

	mux := http.NewServeMux()

	rootHandler := http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		if req.URL.String() != "/" {
			r.WriteHeader(http.StatusNotFound)
			r.Write([]byte("404 Not Found"))
			return
		}
		r.WriteHeader(http.StatusOK)
		r.Write([]byte(`<a href="/on">Turn on</a>`))
		return
	})
	mux.Handle("/", rootHandler)

	onHandler := http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		go chat.Client().Run()
		r.WriteHeader(http.StatusOK)
		r.Write([]byte("<p>Chat bot started</p>"))
		return
	})
	mux.Handle("/on", onHandler)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
	}
}
