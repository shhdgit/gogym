// Mock date command
package main

import (
  "fmt"
  "os"
)

func main() {
  cmdText := os.Args[1:]

  for i, argsLen := 0, len(cmdText); i < argsLen; i++ {
    fmt.Print(cmdText[i], " ")
  }

  fmt.Print("\n")
}
