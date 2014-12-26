package token

func escape(data []byte) (int, byte) {
	if data[0] != '\\' {
		return -1, 0
	}
	l := len(data)
	if l < 2 {
		return 0, 0
	}
	switch data[1] {
	case 'r':
		return 2, byte('\r')
	case 'n':
		return 2, byte('\n')
	case 't':
		return 2, byte('\t')
	case 'v':
		return 2, byte('\v')
	case 'f':
		return 2, byte('\f')
	case 'a':
		return 2, byte('\a')
	case 'x', 'X':
		if l < 3 {
			return 0, 0
		}
		if j := hex(data[2]); j >= 0 {
			if l < 4 {
				return 0, 0
			}
			if k := hex(data[3]); k >= 0 {
				return 4, byte(j*16 + k)
			}
		}
	}
	return 2, data[1]
}

// 预定义的转义字符、十六进制表示的ASCII字符
func Escape(data []byte, end bool) (int, interface{}) {
	i, c := escape(data)
	if i > 0 {
		return i, c
	} else {
		return i, nil
	}
}

// 所有其余未转义的ASCII字符
func Byte(data []byte, end bool) (int, interface{}) {
	return 1, data[1]
}
