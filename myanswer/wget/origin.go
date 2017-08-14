package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "At least need 3 parameters.")
		os.Exit(1)
	}

	url := os.Args[1]
	fileName := os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	defer file.Close()

	io.Copy(file, resp.Body)
}
