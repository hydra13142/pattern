package token

// 限定重复次数
func Limit(data []byte, end bool) (int, interface{}) {
	var (
		i, j, l int
		a, b    int
		p, q    bool
	)
	if data[0] != '{' {
		return -1, nil
	}
	l = len(data)
	i, a = Dec(data, 1, 4)
	if i >= l {
		return 0, nil
	}
	p = (j != 1)
	if data[i] == ',' {
		i++
		j, b = Dec(data, i, 4)
		if j >= l {
			return 0, nil
		}
		q = (i != j)
	} else {
		b = a
	}
	if data[j] == '}' && (p || q) {
		return j + 1, [2]int{a, b}
	}
	return -1, nil
}
