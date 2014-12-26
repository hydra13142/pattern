package token

// 整数
func Integer(data []byte, end bool) (i int, r interface{}) {
	t, l := 0, len(data)
	j, s := 0, true
	if data[0] == '+' {
		t, s = 1, true
	} else if data[0] == '-' {
		t, s = 1, false
	}
	if l == t {
		return 0, nil
	}
	if data[t] != '0' {
		i, j = Dec(data, t, 20)
	} else if l == t+1 {
		i = t + 1
	} else {
		t++
		if data[t] == 'x' || data[t] == 'X' {
			t++
			i, j = Hex(data, t, 16)
		} else {
			i, j = Oct(data, t, 32)
		}
	}
	if i >= l && !end {
		return 0, nil
	}
	if i == t || data[i] == '.' || u_w.Get(data[i]) {
		return -1, nil
	}
	if s {
		return i, +j
	} else {
		return i, -j
	}
}
