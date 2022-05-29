package main

import (
  "strings"
  "fmt"
)

func main() {
  total := 0
  temp := ""
  str := "abcba"
  target := "abc"
  fullMatch := true
  for fullMatch == true {
    for _, letter := range target {
      if strings.Contains(str, string(letter)) {
        temp += string(letter)
        str = strings.Replace(str, string(letter), "", 1)
          if temp == target {
            total++
            temp = ""
          }
      } else {
        fullMatch = false
        break
      }
    }
  }
  fmt.Println(total)
}
