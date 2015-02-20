package DFA

// 正则语法的token
type Token struct {
	K byte
	V interface{}
}

// 识别单个字符、转义字符、预定义字符组
func Alone(i int, data []byte) (int, Token, error) {
	if data[i] != '\\' {
		return 1, Token{'c', data[i]}, nil
	}
	hex := func(c byte) int {
		switch {
		case c >= '0' && c <= '9':
			return int(c - '0')
		case c >= 'A' && c <= 'F':
			return int(c - ('A' - 10))
		case c >= 'a' && c <= 'f':
			return int(c - ('a' - 10))
		}
		return -1
	}
	j, l := i+1, len(data)
	if j < l {
		switch data[j] {
		case 'r':
			return 2, Token{'c', byte('\r')}, nil
		case 'n':
			return 2, Token{'c', byte('\n')}, nil
		case 't':
			return 2, Token{'c', byte('\t')}, nil
		case 'v':
			return 2, Token{'c', byte('\v')}, nil
		case 'f':
			return 2, Token{'c', byte('\f')}, nil
		case 'a':
			return 2, Token{'c', byte('\a')}, nil
		case 's':
			return 2, Token{'u', make(Unit, 8).Set(Space)}, nil
		case 'S':
			return 2, Token{'u', make(Unit, 8).Not(Space)}, nil
		case 'd':
			return 2, Token{'u', make(Unit, 8).Set(Digit)}, nil
		case 'D':
			return 2, Token{'u', make(Unit, 8).Not(Digit)}, nil
		case 'c':
			return 2, Token{'u', make(Unit, 8).Set(Alpha)}, nil
		case 'C':
			return 2, Token{'u', make(Unit, 8).Not(Alpha)}, nil
		case 'w':
			return 2, Token{'u', make(Unit, 8).Set(Words)}, nil
		case 'W':
			return 2, Token{'u', make(Unit, 8).Not(Words)}, nil
		case 'x':
			if j+2 < l {
				j++
				if t := hex(data[j]); t >= 0 {
					j++
					if k := hex(data[j]); k >= 0 {
						return 4, Token{byte((t << 4) | k), nil}, nil
					}
				}
			}
		default:
			return 2, Token{'c', data[j]}, nil
		}
	}
	return 0, Token{}, LexicalError(j)
}

// 识别自定义字符组
func Group(i int, data []byte) (int, Token, error) {
	var (
		e       error
		x       Token
		p, q, r bool
		j, l, t int
		o       byte
	)
	j, l = i+1, len(data)
	if j >= l {
		return 0, Token{}, LexicalError(j)
	}
	if data[j] == '^' {
		j, r = j+1, true
	}
	brac := func(a, b byte) Unit {
		c := make(Unit, 8)
		if a > b {
			a, b = b, a
		}
		for ; a <= b; a++ {
			c.SetBit(int(a))
		}
		return c
	}
	u := make(Unit, 8)
	for ; j < l && data[j] != ']'; j += t {
		if data[j] == '-' {
			if p {
				if q {
					u.Or(u, brac(o, '-'))
					p, q = false, false
				} else {
					o, q = '-', true
				}
			} else {
				p = true
			}
			t = 1
			continue
		}
		if data[j] == '[' {
			t, x, e = Group(j, data)
		} else {
			t, x, e = Alone(j, data)
		}
		if e != nil {
			return 0, Token{}, e
		}
		if x.K == 'c' {
			c := x.V.(byte)
			if p {
				if q {
					u.Or(u, brac(o, c))
					q = false
				} else {
					u.SetBit('-')
					q = true
				}
			} else {
				if q {
					u.SetBit(int(o))
				}
				q = true
			}
			o = c
		} else {
			if q {
				u.SetBit(int(o))
			}
			if p {
				u.SetBit('-')
			}
			u.Or(u, x.V.(Unit))
			q = false
		}
		p = false
	}
	if j >= l {
		return 0, Token{}, LexicalError(j)
	}
	if q {
		u.SetBit(int(o))
	}
	if p {
		u.SetBit('-')
	}
	if r {
		u.Not(u)
	}
	return j - i + 1, Token{'u', u}, nil
}

// 识别限定次数重复
func Limit(i int, data []byte) (int, Token, error) {
	var (
		j, k, a, b int
		l          = len(data)
	)
	Dec := func(b int, s []byte, l int) (i int, n int) {
		l += b
		if w := len(s); l > w {
			l = w
		}
		for i = b; i < l; i++ {
			if c := s[i]; c < '0' || c > '9' {
				break
			} else {
				n = n*10 + int(c-'0')
			}
		}
		return i, n
	}
	j, a = Dec(i+1, data, 4)
	if j < l {
		switch data[j] {
		case '}':
			if i+1 != j {
				return j - i + 1, Token{'n', [2]int{a, a}}, nil
			}
		case ',':
			k, b = Dec(j+1, data, 4)
			if k < l {
				if j+1 == k {
					b = -1
				}
				if data[k] == '}' && (i+1 != j || j+1 != k) {
					if a > b && b >= 0 {
						a, b = b, a
					}
					if b != 0 {
						return k - i + 1, Token{'n', [2]int{a, b}}, nil
					}
				}
			}
			j = k
		}
	}
	return 0, Token{}, LexicalError(j)
}

// 识别所有token
func translate(data []byte) ([]Token, error) {
	var (
		s    []Token
		e    error
		t    Token
		i, j int
	)
	for l := len(data); i < l; i += j {
		switch data[i] {
		case '(':
			j, t, e = 1, Token{'o', byte('(')}, nil
		case ')':
			j, t, e = 1, Token{'o', byte(')')}, nil
		case '|':
			j, t, e = 1, Token{'o', byte('|')}, nil
		case '?':
			j, t, e = 1, Token{'o', byte('?')}, nil
		case '*':
			j, t, e = 1, Token{'o', byte('*')}, nil
		case '+':
			j, t, e = 1, Token{'o', byte('+')}, nil
		case '{':
			j, t, e = Limit(i, data)
		case '[':
			j, t, e = Group(i, data)
		default:
			j, t, e = Alone(i, data)
		}
		if e != nil {
			return nil, e
		}
		s = append(s, t)
	}
	return s, nil
}
