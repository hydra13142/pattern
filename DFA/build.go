package DFA

// 中间的NFA单个节点
type NFA struct {
	cost int
	next int
	skip int
	back int
	unit Unit
}

func build(tkn []Token) ([]NFA, error) {
	p, q, err := check(tkn)
	if err != nil {
		return nil, err
	}
	N := make([]NFA, p+1)
	Q := make([]byte, q*3+3)
	S := make([]struct{ F, I, O, T int }, q*3+3)
	p, q = 0, 0
	S[0] = struct{ F, I, O, T int }{0, 0, 0, 1}
	Q[0] = '('
	R := false
	pop := func(until func(byte) bool) {
		for ; !until(Q[q]); q-- {
			I, O, T := S[p].I, S[p].O, S[p].T
			switch Q[q] {
			case '&':
				p--
				N[S[p].O].next = I - S[p].O
				N[I].back = S[p].O - I
				S[p].O = O
				S[p].T = T
			case '|':
				p--
				X, Y := S[p].I, S[p].O
				N[I].back = X - I
				N[X].back = T - X
				N[T].next = I - T
				S[p].I = T
				T++
				N[Y].next = O - Y
				N[O].next = T - O
				N[T].skip = Y - T
				S[p].O = T
				T++
				S[p].T = T
			case '(':
				N[T].next = I - T
				N[I].back = T - I
				S[p].I = T
				T++
				N[O].next = T - O
				N[T].skip = O - T
				S[p].O = T
				T++
				S[p].T = T
			}
		}
	}
	for i := 0; i < len(tkn); i++ {
		t := tkn[i]
		if R {
			if t.K == 'u' || t.K == 'c' || (t.K == 'o' && t.V.(byte) == '(') {
				t = Token{'o', byte('&')}
				i--
			}
		}
		switch t.K {
		case 'o':
			switch t.V.(byte) {
			case '(':
				q++
				Q[q] = '('
				R = false
			case ')':
				pop(func() func(byte) bool {
					n := byte(0)
					return func(c byte) bool {
						t := n == '('
						n = c
						return t
					}
				}())
				R = true
			case '?':
				N[S[p].I].skip = S[p].O - S[p].I
			case '+':
				N[S[p].O].back = S[p].I - S[p].O
			case '*':
				N[S[p].I].skip = S[p].O - S[p].I
				N[S[p].O].back = S[p].I - S[p].O
			case '|':
				pop(func(c byte) bool {
					return c == '('
				})
				q++
				Q[q] = '|'
				R = false
			case '&':
				pop(func(c byte) bool {
					return c != '&'
				})
				q++
				Q[q] = '&'
				R = false
			}
		case 'n':
			F, I, O, T := S[p].F, S[p].I, S[p].O, S[p].T
			i, d, l := 1, T-F, t.V.([2]int)
			if l[0] > 0 {
				N[O].next = I + d - O
				for ; i < l[0]; i++ {
					copy(N[T:], N[F:T])
					F, T, I, O = F+d, T+d, I+d, O+d
				}
				if l[1] > l[0] {
					copy(N[T:], N[F:T])
					F, T, I, O = F+d, T+d, I+d, O+d
					N[I].skip = O - I
					for i++; i < l[1]; i++ {
						copy(N[T:], N[F:T])
						F, T, I, O = F+d, T+d, I+d, O+d
					}
				} else if l[1] < 0 {
					N[O].back = I - O
				}
				N[O].next = 0
			} else if l[1] > 0 {
				N[I].skip = O - I
				N[O].next = I + d - O
				for ; i < l[1]; i++ {
					copy(N[T:], N[F:T])
					F, T, I, O = F+d, T+d, I+d, O+d
				}
				N[O].next = 0
			} else {
				N[I].skip = O - I
				N[O].back = I - O
			}
			S[p].O = O
			S[p].T = T
		default:
			T := S[p].T
			N[T].cost = 1
			if t.K == 'u' {
				N[T].unit = t.V.(Unit)
			} else {
				N[T].unit = make(Unit, 8).SetBit(int(t.V.(byte)))
			}
			p++
			S[p] = struct{ F, I, O, T int }{T, T, T + 1, T + 2}
			R = true
		}
	}
	pop(func(c byte) bool {
		return c == '('
	})
	N[0].next = S[p].I
	return N, nil
}
