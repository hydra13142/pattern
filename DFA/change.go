package DFA

// 经典DFA正则，没有分组等NFA正则的功能，速度快
type DFA struct {
	Move [][]int
	Char []int
	Over []bool
}

func change(nfa []NFA, unt []Unit) (*DFA, error) {
	L := len(nfa)
	N := L >> 5
	if L&31 != 0 {
		N++
	}
	i, j, l := 0, 0, len(unt)
	extend := func(u []int) Unit {
		x := make(Unit, N)
		add := func(i int) {
			if !x.GetBit(i) {
				x.SetBit(i)
				u = append(u, i)
			}
		}
		for j := 0; j < len(u); j++ {
			i := u[j]
			add(i + nfa[i].next)
			add(i + nfa[i].skip)
			add(i + nfa[i].back)
		}
		return x
	}
	jump := func(x, y Unit) Unit {
		u := []int{}
		for i := 0; i < L; i++ {
			if x.GetBit(i) && nfa[i].cost != 0 {
				if Both(nfa[i].unit, y) {
					u = append(u, i+nfa[i].cost)
				}
			}
		}
		return extend(u)
	}
	med := make([]Unit, L)
	med[0] = extend([]int{0})
	dfa := new(DFA)
	dfa.Move = make([][]int, L)
	for ; i <= j; i++ {
		dfa.Move[i] = make([]int, l)
		for k := 0; k < l; k++ {
			u := jump(med[i], unt[k])
			if None(u) {
				dfa.Move[i][k] = -1
			} else {
				t := 0
				for ; t <= j; t++ {
					if Same(med[t], u) {
						break
					}
				}
				if t > j {
					j++
					med[j] = u
				}
				dfa.Move[i][k] = t
			}
		}
	}
	dfa.Char = array(unt)
	dfa.Over = make([]bool, i)
	for j = 0; j < i; j++ {
		dfa.Over[j] = med[j].GetBit(L - 1)
	}
	dfa.Move = dfa.Move[:i]
	return dfa, nil
}
