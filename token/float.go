package token

import "math"

// 浮点数
func Float(data []byte, end bool) (i int, r interface{}) {
	var (
		t, n, e int
		p, f    float64
	)
	Int := func(data []byte) (i, j int) {
		for l := len(data); i < l; i++ {
			c := dec(data[i])
			if c < 0 {
				break
			}
			j = j*10 + c
		}
		if i >= len(data) {
			return 0, 0
		}
		if i == 0 {
			return -1, 0
		}
		return i, j
	}
	Frac := func(data []byte) (int, float64) {
		if len(data) < 2 || data[0] != '.' {
			return -1, 0
		}
		i, j := Int(data[1:])
		if i <= 0 {
			return i, 0
		}
		i, k, l := i+1, 1, 1
		for ; k < i; k++ {
			l *= 10
		}
		return i, float64(j) / float64(l)
	}
	Exp := func(data []byte) (i, j int) {
		if len(data) < 2 || (data[0] != 'e' && data[0] != 'E') {
			return -1, 0
		}
		t, p := 1, true
		if data[1] == '+' {
			t, p = 2, true
		} else if data[1] == '-' {
			t, p = 2, false
		}
		i, j = Int(data[t:])
		if i <= 0 {
			return i, 0
		}
		if p {
			return i + t, +j
		} else {
			return i + t, -j
		}
	}
	l, s := len(data), true
	if data[0] == '+' {
		i, s = 1, true
	} else if data[0] == '-' {
		i, s = 1, false
	}
	t, n = Int(data[i:])
	if t <= 0 {
		return t, nil
	} else {
		i, f = i+t, float64(n)
		if i == l {
			goto exit
		}
	}
	t, p = Frac(data[i:])
	if t > 0 {
		i, f = i+t, f+p
		if i == l {
			goto exit
		}
	}
	t, e = Exp(data[i:])
	if t > 0 {
		i, f = i+t, f*math.Pow10(e)
		if i == l {
			goto exit
		}
	}
	if data[i] == '.' || u_w.Get(data[i]) {
		return -1, nil
	}
	goto done
exit:
	if !end {
		return 0, nil
	}
done:
	if s {
		return i, +f
	} else {
		return i, -f
	}
}
