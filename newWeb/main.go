package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", firstPage)
	log.Println("web server is up!")
	http.ListenAndServe(":8080", nil)
}

func firstPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("this is second web page!!!!"))
}
