package pattern

import "github.com/hydra13142/pattern/DFA"

func TokenDFA(s string) (func([]byte, bool) (int, interface{}), error) {
	dfa, err := DFA.NewDFA(s)
	if err != nil {
		return nil, err
	}
	return func(data []byte, end bool) (int, interface{}) {
		i, j, k, l := 0, 0, -1, len(data)
		for j < l {
			c := data[j]
			k := dfa.Char[c]
			j++
			if i = dfa.Move[i][k]; i < 0 {
				break
			}
			if dfa.Over[i] {
				k = j
			}
		}
		if j >= l && !end {
			return 0, nil
		}
		if k > 0 {
			return k, data[:k]
		} else {
			return -1, nil
		}
	}, nil
}

func TokenLiteral(s string) func([]byte, bool) (int, interface{}) {
	b := []byte(s)
	return func(data []byte, end bool) (int, interface{}) {
		if len(data) < len(b) {
			return 0, nil
		}
		for i := len(b) - 1; i >= 0; i-- {
			if data[i] != b[i] {
				return -1, nil
			}
		}
		return len(b), data[:len(b)]
	}
}
