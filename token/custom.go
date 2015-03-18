package token

import "github.com/hydra13142/pattern/DFA"

// 匹配DFA正则，返回[]byte
func TokenDFA(s string) (func([]byte, bool) (int, interface{}), error) {
	dfa, err := DFA.NewDFA(s)
	if err != nil {
		return nil, err
	}
	return func(data []byte, end bool) (int, interface{}) {
		i, j, k, l := 0, 0, -1, len(data)
		if dfa.Over[0] {
			k = 0
		}
		for j < l {
			if i = dfa.Move[i][dfa.Char[data[j]]]; i < 0 {
				break
			}
			if j++; dfa.Over[i] {
				k = j
			}
		}
		if j >= l && !end {
			return 0, nil
		}
		if k > 0 {
			return k, data[:k]
		}
		return -1, nil
	}, nil
}

// 匹配字符串字面量，返回[]byte
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
