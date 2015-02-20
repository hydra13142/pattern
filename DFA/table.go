package DFA

func table(tkn []Token) []Unit {
	x := make(Unit, 8)
	y := make(Unit, 8)
	z := make(Unit, 8)
	u := []Unit{x.Not(x)}
	n := 0
	for _, t := range tkn {
		x = make(Unit, 8)
		switch t.K {
		case 'c':
			x.SetBit(int(t.V.(byte)))
		case 'u':
			x.Set(t.V.(Unit))
		default:
			continue
		}
		for i := n; !None(x); i-- {
			if y.And(u[i], x); !None(y) {
				z.Sub(u[i], x)
				if !None(z) {
					u[i].Set(z)
					u = append(u, make(Unit, 8).Set(y))
					n++
				}
				x.Sub(x, y)
			}
		}
	}
	return append(u, u[0])[1:]
}

func array(u []Unit) []int {
	s := make([]int, 256)
	for i := 0; i < len(u); i++ {
		for j := 0; j < 256; j++ {
			if u[i].GetBit(j) {
				s[j] = i
			}
		}
	}
	return s
}
