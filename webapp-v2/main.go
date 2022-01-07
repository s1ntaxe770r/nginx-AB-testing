package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "[ webapp-v2 ]", log.LstdFlags)
	r := http.NewServeMux()
	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw, r, "./static/index.html")
	})
	r.HandleFunc("/version", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("v2"))
	})
	l.Printf("Starting server on port 9090...")
	http.ListenAndServe(":9090", r)
}
