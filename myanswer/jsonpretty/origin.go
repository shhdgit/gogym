package main

import (
	"encoding/json"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: go run jsonpretty.go [file]")
	}

	filepath := os.Args[1]
	extname := path.Ext(filepath)
	if extname != ".json" {
		log.Fatalln("Need a json file")
	}

	f, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	var data interface{}
	dec := json.NewDecoder(f)
	err = dec.Decode(&data)
	if err != nil {
		log.Fatalln(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	err = enc.Encode(&data)
	if err != nil {
		log.Fatalln(err)
	}
}
