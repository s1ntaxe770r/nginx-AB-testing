package main

import (
	"log"
	"net/http"
	"os"
)


func main(){
	l := log.New(os.Stdout,"[ webapp-v1 ]",log.LstdFlags)
	r := http.NewServeMux()
	r.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		http.ServeFile(rw,r,"./static/index.html")		
	})
	l.Printf("Starting server on port 8080")	
	http.ListenAndServe(":8080",r)
}
