package bap

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
		chf := make(chan []int)
		chs := make(chan []int)
		go func() {
			defer close(chf)
			chf <- separator(s[:len(s)/2], false)
		}()
		go func() {
			defer close(chs)
			chs <- separator(s[len(s)/2:], true)
		}()
		return merger(append(<-chf, <-chs...), down)
	} else {
		chf := make(chan []int)
		chs := make(chan []int)
		go func() {
			defer close(chf)
			chf <- separator(s[:len(s)/2+1], false)
		}()
		go func() {
			defer close(chs)
			chs <- separator(s[len(s)/2+1:], true)
		}()
		return merger(append(<-chf, <-chs...), down)
	}
}

func merger(s []int, down bool) []int {
	if len(s) == 1 {
		return s
	} else {
		comparer(s, down)
		a := pow2(len(s))
		chf := make(chan []int)
		chs := make(chan []int)
		if down {
			go func() {
				defer close(chf)
				chf <- merger(s[:a], down)
			}()
			go func() {
				defer close(chs)
				chs <- merger(s[a:], down)
			}()
		} else {
			go func() {
				defer close(chf)
				chf <- merger(s[:len(s)-a], down)
			}()
			go func() {
				defer close(chs)
				chs <- merger(s[len(s)-a:], down)
			}()
		}
		return append(<-chf, <-chs...)
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
