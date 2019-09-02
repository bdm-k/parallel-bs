package bap

import (
	"math/rand"
	"testing"
	"time"
)

func makeIntSlice() (s []int){
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

func TestMain(m *testing.M) {
	for i := 0; i < 10; i++ {
		m.Run()
	}
}

func TestSeparator(t *testing.T) {
	s := makeIntSlice()
	t.Log("length of sequence : ", len(s))

	ds := separator(s, true)
	if check(ds, true) {
		t.Log("\nsorted down successfully")
	} else {
		t.Fatalf("\nsorting up failed\noriginal\n%v\nsorted\n%v", s, ds)
	}

	us := separator(s, false)
	if check(us, false) {
		t.Log("\nsorted up successfully")
	} else {
		t.Fatalf("sorting down failed\noriginal\n%v\nsorted\n%v", s, us)
	}
}
