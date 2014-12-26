package token

// 边界
func Edge(data []byte, end bool) (int, interface{}) {
	if data[0] != '\\' {
		return -1, nil
	}
	if len(data) < 2 {
		return 0, nil
	}
	switch data[1] {
	case 'A', 'Z':
		fallthrough
	case 'b', 'B':
		return 2, data[1]
	}
	return -1, nil
}
