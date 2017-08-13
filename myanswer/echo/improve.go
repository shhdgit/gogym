package main

import (
  "os"
  "fmt"
)

func main() {
  argLen := len(os.Args)

  for i, arg := range os.Args {
    if i == 0 {
      continue
    }

    fmt.Print(arg)
    if i < argLen - 1 {
      fmt.Print(" ")
    }
  }

  fmt.Print("\n")
}
