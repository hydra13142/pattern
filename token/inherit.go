package token

import "github.com/hydra13142/pattern/DFA"

func Alone(data []byte, end bool) (int, interface{}) {
	i, c, e := DFA.Alone(0, data)
	if e == nil {
		return i, c
	}
	if int(e.(DFA.LexicalError)) < len(data) || end {
		return -1, nil
	}
	return 0, nil
}

func Group(data []byte, end bool) (int, interface{}) {
	i, u, e := DFA.Group(0, data)
	if e == nil {
		return i, u
	}
	if int(e.(DFA.LexicalError)) < len(data) || end {
		return -1, nil
	}
	return 0, nil
}

func Limit(data []byte, end bool) (int, interface{}) {
	i, p, e := DFA.Limit(0, data)
	if e == nil {
		return i, p
	}
	if int(e.(DFA.LexicalError)) < len(data) || end {
		return -1, nil
	}
	return 0, nil
}

func Escape(data []byte, end bool) (int, interface{}) {
	if data[0] != '\\' {
		return -1, nil
	}
	i, c, e := DFA.Alone(0, data)
	if e == nil {
		return i, c
	}
	if int(e.(DFA.LexicalError)) < len(data) || end {
		return -1, nil
	}
	return 0, nil
}

func Byte(data []byte, end bool) (int, interface{}) {
	return 1, data[0]
}