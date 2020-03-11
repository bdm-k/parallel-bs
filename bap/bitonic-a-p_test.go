//TODO: create function, TestSpeed

package bap

import (
	"math/rand"
	"testing"
	"time"
)

func makeIntSlice() (s []int) {
	rand.Seed(time.Now().UnixNano())
	length := 0
	for length == 0 {
		length = rand.Intn(10000)
	}
	s = rand.Perm(length)
	return
}

func check(s []int, down bool) (r bool) {
	length := len(s)
	for i := 0; i < length-1; i++ {
		if s[i] == s[i+1] {

		} else if b := s[i] < s[i+1]; b != down {
			r = false
			return
		}
	}
	r = true
	return
}

func TestAccuracy(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := makeIntSlice()
		t.Log("length of sequence : ", len(s))

		ds := separator(s, true)
		if check(ds, true) {
			t.Log("\nsorted down successfully")
		} else {
			t.Fatal("\nsorting down failed")
		}

		us := separator(s, false)
		if check(us, false) {
			t.Log("\nsorted up successfully")
		} else {
			t.Fatal("\nsorting up failed")
		}
	}
}

func TestSpeed(t *testing.T) {
	s := func() []int {
		rand.Seed(621)
		return rand.Perm(10000)
	}()
	var bs, us []int
	beforeloop := time.Now()
	for i := 0; i < 100; i++ {
		bs = separator(s, true)
		us = separator(s, false)
	}
	afterloop := time.Now()
	_, _ = bs, us
	t.Logf("\ntime: %v", afterloop.Sub(beforeloop))
}
