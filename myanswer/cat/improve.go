package main

import (
  "os"
  "io"
  "fmt"
)

func main() {
  for _, fileName := range os.Args[1:] {
    f, err := os.Open(fileName)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      continue
    }

    _, err = io.Copy(os.Stdout, f)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }

    err = f.Close()
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
    }
  }
}
