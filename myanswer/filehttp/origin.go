package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var basedir string

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: go run server.go [basedir]")
	}

	basedir = os.Args[1]
	fmt.Printf("Serving content: %s\n", basedir)

	http.HandleFunc("/", serverContent)
	http.ListenAndServe(":7777", nil)
}

func serverContent(w http.ResponseWriter, r *http.Request) {
	reqpath := r.URL.Path

	// fmt.Fprintf(w, "URL.Path = %q\n", path)
	if reqpath == "/" {
		reqpath = "index.html"
	}

	filepath := path.Join(basedir, reqpath)

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Fprintf(w, "%q\n", err)
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Fprintf(w, "%q\n", err)
	}
}
