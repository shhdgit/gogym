// Mock date command
package main

import (
  "os"
  "fmt"
)

func main() {
  args := os.Args[1:]
  argLen := len(args)

  for i, arg := range args {
    fmt.Print(arg)
    if i < argLen {
      fmt.Print(" ")
    }
  }

  fmt.Print("\n")
}
