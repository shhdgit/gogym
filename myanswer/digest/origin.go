package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: data | go run digest.go [digest type: md5, sha256, sha512].")
	}

	var h hash.Hash
	dtype := os.Args[1]
	switch dtype {
	case "md5":
		h = md5.New()
	case "sha256":
		h = sha256.New()
	case "sha512":
		h = sha512.New()
	default:
		log.Fatalln("Error digest type[md5, sha256, sha512].")
	}

	_, err := io.Copy(h, os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(hex.EncodeToString(h.Sum(nil)))
}
