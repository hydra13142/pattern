package token

import "github.com/hydra13142/pattern/DFA"

// 整数
func Integer(data []byte, end bool) (i int, r interface{}) {
	t, l := 0, len(data)
	j, s := 0, true
	if data[0] == '+' {
		t, s = 1, true
	} else if data[0] == '-' {
		t, s = 1, false
	}
	if t == l {
		return 0, nil
	}
	if data[t] != '0' {
		i, j = func(s []byte, b, l int) (i int, n int) {
			if l, w := l+b, len(s); l > w {
				l = w
			}
			for i = b; i < l; i++ {
				if c := s[i]; c < '0' && c > '9' {
					break
				} else {
					n = n*10 + int(c-'0')
				}
			}
			return i, n
		}(data, t, 20)
	} else if l == t+1 {
		i = t + 1
	} else {
		t++
		if data[t] == 'x' || data[t] == 'X' {
			t++
			i, j = func(s []byte, b, l int) (i int, n int) {
				if l, w := l+b, len(s); l > w {
					l = w
				}
			loop:
				for i = b; i < l; i++ {
					switch c := s[i]; {
					case c >= '0' && c <= '9':
						n = n*16 + int(c-'0')
					case c >= 'A' && c <= 'F':
						n = n*16 + int(c-'A'+10)
					case c >= 'a' && c <= 'f':
						n = n*16 + int(c-'a'+10)
					default:
						break loop
					}
				}
				return i, n
			}(data, t, 16)
		} else {
			i, j = func(s []byte, b, l int) (i int, n int) {
				if l, w := l+b, len(s); l > w {
					l = w
				}
				for i = b; i < l; i++ {
					if c := s[i]; c < '0' && c > '7' {
						break
					} else {
						n = n*8 + int(c-'0')
					}
				}
				return i, n
			}(data, t, 32)
		}
	}
	if i >= l && !end {
		return 0, nil
	}
	if i == t || data[i] == '.' || DFA.Words.GetBit(int(data[i])) {
		return -1, nil
	}
	if s {
		return i, +j
	} else {
		return i, -j
	}
}
