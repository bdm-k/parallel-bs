package main

import (
  "fmt"
)

func main() {
  var s []int = []int{1, 5, 13, 3, 14, 16, 12, 15, 4, 8, 7, 9, 10, 11, 6, 2}
  s = separator(s, true)
  fmt.Println(s)
}

func separator(s []int, donw bool) []int {
  if len(s) == 1 {
    return s
  } else {
    f := separator(s[:len(s)/2], true)
    s := separator(s[len(s)/2:], false)
    return merger(append(f, s...), down)
  }
}

func merger(s []int, down bool) []int {
  if len(s) == 1 {
    return s
  } else {
    comparer(s, down)
    f := merger(s[:len(s)/2], down)
    s := merger(s[len(s)/2:], down)
    return append(f, s...)
  }
}

func comparer(s []int, down bool) {
  a := len(s) / 2
  for i := 0; i < a; i++ {
    if b := s[i] > s[i+a]; b == down {
      s[i], s[i+a] = s[i+a], s[i]
    }
  }
}
