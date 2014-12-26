package token

// 空白
func Space(data []byte, end bool) (int, interface{}) {
	i, l := 0, len(data)
	for i < l && u_s.Get(data[i]) {
		i++
	}
	if i >= l && !end {
		return 0, nil
	}
	if i == 0 {
		return -1, nil
	}
	return i, nil
}

// 连续可打印ASCII字符
func Field(data []byte, end bool) (int, interface{}) {
	i, l := 0, len(data)
	for i < l && data[i] > 32 && data[i] < 127 {
		i++
	}
	if i >= l && !end {
		return 0, nil
	}
	if i == 0 {
		return -1, nil
	}
	return i, string(data[:i])
}

// 标识符
func Word(data []byte, end bool) (int, interface{}) {
	if u_C.Get(data[0]) && data[0] != '_' {
		return -1, nil
	}
	i, l := 1, len(data)
	for i < l && u_w.Get(data[i]) {
		i++
	}
	if i >= l && !end {
		return 0, nil
	}
	return i, string(data[:i])
}
