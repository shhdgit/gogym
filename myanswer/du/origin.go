package main

import (
	"log"
	"os"
)

var total int64

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: du [file|folder]")
	}

	filepath := os.Args[1]
	walk(filepath, loginfo)
	log.Printf("disk usage: %dbyte", total)
}

func walk(filename string, cb func(string, os.FileInfo)) {
	finfo, err := os.Stat(filename)
	if err != nil {
		log.Fatalln(err)
	}

	if finfo.IsDir() {
		folder, err := os.Open(filename)
		if err != nil {
			log.Fatalln(err)
		}

		folderinfos, err := folder.Readdir(-1)
		folder.Close()
		if err != nil {
			log.Fatalln(err)
		}

		for _, item := range folderinfos {
			walk(filename+"/"+item.Name(), cb)
		}
	} else {
		cb(filename, finfo)
	}
}

func loginfo(filename string, finfo os.FileInfo) {
	log.Printf("%s : %dbyte", filename, finfo.Size())
	total += finfo.Size()
}
