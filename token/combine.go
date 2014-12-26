package token

// 顺序匹配
func Tolerate(head, tail string, take func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{}) {
	return func(data []byte, end bool) (int, interface{}) {
		a, b := len(head), len(tail)
		i, l := 0, len(head)
		if a > 0 {
			if a > l {
				return 0, nil
			}
			if string(data[:a]) != head {
				return -1, nil
			}
			i += a
		}
		j, r := take(data[:a], end)
		if j <= 0 {
			return j, nil
		}
		i += j
		if b > 0 {
			if i+b > l {
				return 0, nil
			}
			if string(data[i:i+b]) != head {
				return -1, nil
			}
			i += b
		}
		return i, r
	}
}

// 选择其一匹配
func Alternative(takes ...func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{}) {
	return func(data []byte, end bool) (int, interface{}) {
		sp := false
		for _, take := range takes {
			i, r := take(data, end)
			if i > 0 {
				return i, r
			}
			if i == 0 {
				sp = true
			}
		}
		if sp {
			return 0, nil
		} else {
			return -1, nil
		}
	}
}
