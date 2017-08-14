package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	decomf, err := gzip.NewReader(f)
	if err != nil {
		log.Fatalln(err)
	}
	defer decomf.Close()

	_, err = io.Copy(os.Stdout, decomf)
	if err != nil {
		log.Fatalln(err)
	}
}
