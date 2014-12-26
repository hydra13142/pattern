package token

func oct(c byte) int {
	if c >= '0' && c <= '7' {
		return int(c - '0')
	}
	return -1
}

func dec(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	return -1
}

func hex(c byte) int {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0')
	case c >= 'A' && c <= 'F':
		return int(c - 'A' + 10)
	case c >= 'a' && c <= 'f':
		return int(c - 'a' + 10)
	}
	return -1
}

// 八进制整数
func Oct(s []byte, b, l int) (i int, n int) {
	if l, w := l+b, len(s); l > w {
		l = w
	}
	for i = b; i < l; i++ {
		c := oct(s[i])
		if c < 0 {
			break
		}
		n = n*8 + c
	}
	return i, n
}

// 十进制整数
func Dec(s []byte, b, l int) (i int, n int) {
	if l, w := l+b, len(s); l > w {
		l = w
	}
	for i = b; i < l; i++ {
		c := dec(s[i])
		if c < 0 {
			break
		}
		n = n*10 + c
	}
	return i, n
}

// 十六进制整数
func Hex(s []byte, b, l int) (i int, n int) {
	if l, w := l+b, len(s); l > w {
		l = w
	}
	for i = b; i < l; i++ {
		c := hex(s[i])
		if c < 0 {
			break
		}
		n = n*16 + c
	}
	return i, n
}
