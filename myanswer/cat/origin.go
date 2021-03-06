// Mock cat command
package main

import (
  "fmt"
  "os"
  "io"
)

func main() {
  for _, fileName := range os.Args[1:] {
    f, err := os.Open(fileName)
    if err != nil {
      fmt.Fprintln(os.Stderr, err)
      continue
    }

    io.Copy(os.Stdout, f)
    f.Close()
  }
}

