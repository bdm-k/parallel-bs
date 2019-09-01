package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("test/1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

  s := make([]int, 0, 100)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
    a, _ := strconv.Atoi(sc.Text())
		s = append(s, a)
	}
	if err = sc.Err(); err != nil {
		log.Fatal(err)
	}

	result := separator(s, true)
	fmt.Print(result)
}

func pow2(l int) int {
	i := 2
	for i < l {
		i *= 2
	}
	return i / 2
}

func separator(s []int, down bool) []int {
	if len(s) == 1 {
		return s
	} else if len(s)%2 == 0 {
		f := separator(s[:len(s)/2], false)
		s := separator(s[len(s)/2:], true)
		return merger(append(f, s...), down)
	} else {
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
			f := merger(s[:len(s)-a], down)
			s := merger(s[len(s)-a:], down)
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
