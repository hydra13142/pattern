package token

import "github.com/hydra13142/pattern/DFA"

// 匹配任意单字符（接受转义、预定义字符组）
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

// 识别自定义字符组
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

// 匹配形如{i,j}/{i,}/{,j}/{i}的数字范围
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

// 只匹配转义字符和预定义字符组
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

// 匹配任意单字符（无视转义）
func Byte(data []byte, end bool) (int, interface{}) {
	return 1, data[0]
}