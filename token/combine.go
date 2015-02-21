package token

// 顺序依次匹配
func Series(takes ...func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{}) {
	return func(data []byte, end bool) (int, interface{}) {
		i := 0
		for _, take := range takes {
			t, _ := take(data[i:], end)
			switch {
			case t > 0:
				i += t
			case t < 0:
				return -1, nil
			default:
				return 0, nil
			}
		}
		return i, data[:i]
	}
}

// 选择其一匹配
func Choice(takes ...func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{}) {
	return func(data []byte, end bool) (int, interface{}) {
		p := false
		for _, take := range takes {
			i, r := take(data, end)
			if i > 0 {
				return i, r
			}
			if i == 0 {
				p = true
			}
		}
		if p {
			return 0, nil
		}
		return -1, nil
	}
}

// 重复至少一次匹配
func Repeat(take func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{}) {
	return func(data []byte, end bool) (int, interface{}) {
		i, t := 0, 0
		for {
			t, _ = take(data[i:], end)
			if t == 0 {
				return 0, nil
			}
			if t < 0 {
				break
			}
			i += t
		}
		if i > 0 {
			return i, data[:i]
		}
		return -1, nil
	}
}
