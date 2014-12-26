package token

// 字符组（预定义、自定义）
func Class(data []byte, end bool) (int, interface{}) {
	var u Unit
	i := u.compile(data)
	if i > 0 {
		return i, u
	} else {
		return i, nil
	}
}
