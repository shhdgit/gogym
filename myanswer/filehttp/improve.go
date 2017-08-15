package main

import (
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
	log.Printf("Serving content: %s\n", basedir)

	http.HandleFunc("/", serverContent)
	http.ListenAndServe(":7777", nil)
}

func serverContent(res http.ResponseWriter, req *http.Request) {
	reqpath := req.URL.Path
	reqext := path.Ext(reqpath)

	if reqext == "" {
		reqpath = path.Join(reqpath, "index.html")
	}

	f, err := os.Open(path.Join(basedir, reqpath))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	defer f.Close()

	_, err = io.Copy(res, f)
	if err != nil {
		log.Println("Serve content: ", err)
	}
}
