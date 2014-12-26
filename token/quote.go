package token

// 单字符（单引号包围）
func QuoteByte(data []byte, end bool) (int, interface{}) {
	if data[0] != '\'' {
		return -1, nil
	}
	l := len(data)
	if l < 2 {
		return 0, nil
	}
	i, c := escape(data[1:])
	if i < 0 {
		i, c = 1, data[1]
		if c == '\'' {
			return -1, nil
		}
	}
	if i == 0 || l < i+2 {
		return 0, nil
	}
	if data[i+1] != '\'' {
		return -1, nil
	}
	return i + 2, c
}

// 字符串（双引号或反引号包围）
func QuoteString(data []byte, end bool) (i int, q interface{}) {
	l := len(data)
	if data[0] == '"' {
		for i = 1; i < l && data[i] != '"'; i++ {
			if data[i] == '\r' || data[i] == '\r' {
				return -1, nil
			}
			if data[i] == '\\' {
				i++
			}
		}
	} else if data[0] == '`' {
		for i = 1; i < l && data[i] != '`'; i++ {
		}
	} else {
		return -1, nil
	}
	if i >= l {
		return 0, nil
	}
	return i + 1, string(data[1:i])
}
