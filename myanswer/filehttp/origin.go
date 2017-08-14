package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: go run server.go [basedir]")
	}

	basedir := os.Args[1]
	fmt.Printf("Serving content: %s\n", basedir)

	http.HandleFunc("/", serverContent)
	http.ListenAndServe(":7777", nil)
}

func serverContent(res http.ResponseWriter, req *http.Request) {

}
