package DFA

func check(tkn []Token) (int, int, error) {
	j, q, p, b := -1, 0, 1, 0
	s := make([]int, 0, 128)
	for i, t := range tkn {
		switch t.K {
		case 'o':
			switch t.V.(byte) {
			case '(':
				s = append(s, 0)
				j++
				q, p = q+1, 1
				if q > b {
					b = q
				}
			case ')':
				if q--; q < 0 || p == 1 {
					return 0, 0, SyntaxError(i)
				}
				for j--; s[j] != 0; j-- {
					s[j] += s[j+1]
				}
				s[j] = s[j+1] + 2
				s = s[:j+1]
				p = 0
			case '|':
				if p == 1 {
					return 0, 0, SyntaxError(i)
				}
				s = append(s, 2)
				j++
				p = 1
			default:
				if p != 0 {
					return 0, 0, SyntaxError(i)
				}
				p = 2
			}
		case 'n':
			if p != 0 {
				return 0, 0, SyntaxError(i)
			}
			x := t.V.([2]int)
			if x[1] > 0 {
				s[j] *= x[1]
			} else if x[0] > 0 {
				s[j] *= x[0]
			}
			p = 2
		default:
			s = append(s, 2)
			j++
			p = 0
		}
	}
	if q != 0 {
		return 0, 0, SyntaxError(len(tkn))
	}
	q = 0
	for _, j = range s {
		q += j
	}
	return q, b, nil
}
