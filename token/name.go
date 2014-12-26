package token

// \num形式的子匹配
func Group(data []byte, end bool) (int, interface{}) {
	if data[0] != '\\' {
		return -1, nil
	}
	i, j := Dec(data, 1, 3)
	if i == 1 {
		return -1, nil
	}
	if end || i >= 4 || i < len(data) {
		return i, j
	}
	return 0, nil
}

// $name，${name}形式的变量
func Var(data []byte, end bool) (int, interface{}) {
	if data[0] != '$' {
		return -1, nil
	}
	i, l := 2, len(data)
	if data[1] == '{' {
		for i < l && u_w.Get(data[i]) {
			i++
		}
		if i >= l {
			return 0, nil
		}
		if i == 2 {
			return -1, nil
		}
		if data[i] == '}' {
			return i + 1, string(data[2:i])
		}
		return -1, nil
	}
	if u_w.Get(data[1]) {
		for i < l && u_w.Get(data[i]) {
			i++
		}
		if end || i < l {
			return i, string(data[1:i])
		}
		return 0, nil
	}
	return -1, nil
}
