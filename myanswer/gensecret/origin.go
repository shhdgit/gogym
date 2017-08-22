package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: go run gensecret.go [byte].")
	}

	nbytes, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
}
