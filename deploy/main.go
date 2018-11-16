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
	http.ListenAndServe(":8081", nil)
}

func firstPage(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("this is deploy page"))
	exec_shell("/apps/gopath/src/devops/deploy/deploy.sh")

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
