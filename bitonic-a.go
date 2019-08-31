package main

import (
  "fmt"
)

func main()  {

}

func pow2(l int) int {
  i := 2
  for i < l  {
    i *= 2
  }
  return i/2
}

func separator(s []int down bool) []int {
  if len(s) == 1 {
    return s
  } else if len(s)%2 == 0 {
    f := separator(s[:len(s)/2], false)
    s := separator(s[len(s)/2:], true)
    return marger(append(f, s...), down)
  } else if len(s)%2 == 1 {
    f := separator(s[:len(s)/2+1], false)
    s := separator(s[len(s)/2+1:], true)
    return merger(append(f, s...), down)
  }
}

func merger(s []int, down bool) []int {
  if len(s) == 1 {
    return s
  } else {
    comparer(s, down)
    a := pow2(len(s))
    if down {
      f := merger(s[:a], down)
      s := merger(s[a:], down)
      return append(f, s...)
    } else {
      f := merger(s[:len(s)-a])
      s := merger(s[len(s)-a:])
      return append(f, s...)
    }
  }
}

func comparer(s []int, down bool) {
  a := pow2(len(s))
  for i := 0; i < len(s)-a; i++ {
    if b := s[i] > s[i+a]; b == down {
      s[i], s[i+a] = s[i+a], s[i]
    }
  }
}
