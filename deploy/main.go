package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	http.HandleFunc("/", firstPage)
	log.Println("deploy server is up!")
	http.ListenAndServe(":5000", nil)
}

func firstPage(w http.ResponseWriter, req *http.Request) {
	exec_shell("/apps/gopath/src/devops/deploy/deploy.sh")
	w.Write([]byte("this is deploy page"))

}

func exec_shell(s string) {
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())
}
