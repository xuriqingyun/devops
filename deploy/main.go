package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", firstPage)
	log.Println("deploy server is up!")
	http.ListenAndServe(":5000", nil)
}

func firstPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("this is first  page"))
}
