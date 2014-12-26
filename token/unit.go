package token

// 表示ASCII字符（范围0-255）的集合
type Unit [8]uint32

var (
	u_s = Unit{0x00003E00, 0x00000001, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	u_S = Unit{0xFFFFC1FF, 0xFFFFFFFE, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF}

	u_d = Unit{0x00000000, 0x03FF0000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	u_D = Unit{0xFFFFFFFF, 0xFC00FFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF}

	u_c = Unit{0x00000000, 0x00000000, 0x07FFFFFE, 0x07FFFFFE, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	u_C = Unit{0xFFFFFFFF, 0xFFFFFFFF, 0xF8000001, 0xF8000001, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF}

	u_w = Unit{0x00000000, 0x03FF0000, 0x87FFFFFE, 0x07FFFFFE, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	u_W = Unit{0xFFFFFFFF, 0xFC00FFFF, 0x78000001, 0xF8000001, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF, 0xFFFFFFFF}
)

// 设置包含某个字符
func (a *Unit) Set(i byte) {
	(*a)[i>>5] |= 1 << uint32(i&31)
}

// 设置不含某个字符
func (a *Unit) Cls(i byte) {
	(*a)[i>>5] &= ^(1 << uint32(i&31))
}

// 返回是否包含某个字符
func (a *Unit) Get(i byte) bool {
	return (*a)[i>>5]&(1<<uint32(i&31)) != 0
}

// 是否包含全部字符
func (a *Unit) Full() bool {
	for i := 0; i < 8; i++ {
		if (*a)[i] != ^uint32(0) {
			return false
		}
	}
	return true
}

// 是否没有任何字符
func (a *Unit) None() bool {
	for i := 0; i < 8; i++ {
		if (*a)[i] != uint32(0) {
			return false
		}
	}
	return true
}

// 取反集
func (x *Unit) Not(a *Unit) *Unit {
	for i := 0; i < 8; i++ {
		(*x)[i] = ^(*a)[i]
	}
	return x
}

// 交集
func (x *Unit) And(a, b *Unit) *Unit {
	for i := 0; i < 8; i++ {
		(*x)[i] = (*a)[i] & (*b)[i]
	}
	return x
}

// 并集
func (x *Unit) Or(a, b *Unit) *Unit {
	for i := 0; i < 8; i++ {
		(*x)[i] = (*a)[i] | (*b)[i]
	}
	return x
}

// 非共有字符的集合
func (x *Unit) Xor(a, b *Unit) *Unit {
	for i := 0; i < 8; i++ {
		(*x)[i] = (*a)[i] ^ (*b)[i]
	}
	return x
}

// 减去
func (x *Unit) Sub(a, b *Unit) *Unit {
	for i := 0; i < 8; i++ {
		(*x)[i] = (*a)[i] & (^(*b)[i])
	}
	return x
}

// 字符组（预定义、自定义）
func (u *Unit) compile(data []byte) int {
	var (
		i, t, l int
		p, q, r bool
		a, b    byte
	)
	switch data[0] {
	case '\\':
		if len(data) < 2 {
			return 0
		}
		switch data[1] {
		case 's':
			*u = u_s
		case 'S':
			*u = u_S
		case 'b':
			*u = u_d
		case 'D':
			*u = u_D
		case 'a':
			*u = u_c
		case 'C':
			*u = u_C
		case 'w':
			*u = u_w
		case 'W':
			*u = u_W
		default:
			return -1
		}
		return 2
	case '[':
		if l = len(data); l < 2 {
			return 0
		}
		if data[1] == '^' {
			i, r = 2, true
		} else {
			i, r = 1, false
		}
		v := &Unit{}
		for t = 0; i < l && data[i] != ']'; i += t {
			if t = v.compile(data[i:]); t > 0 {
				if q {
					u.Set('-')
				}
				u.Or(u, v)
				p, q = false, false
				continue
			}
			t, b = escape(data[i:]); 
			if t < 0 {
				t, b = 1, data[i]
				if b == '-' {
					q = true
					continue
				}
			}
			if t == 0 {
				return 0
			}
			if q {
				if p {
					for a++; a < b; a++ {
						u.Set(byte(a))
					}
				} else {
					u.Set('-')
				}
			}
			u.Set(byte(b))
			p, q, a = true, false, b
		}
		if i >= l {
			return 0
		}
		if q {
			u.Set('-')
		}
		if r {
			u.Not(u)
		}
		return i + 1
	}
	return -1
}

// 交集是否非空
func And(a, b *Unit) bool {
	for i := 0; i < 8; i++ {
		if (*a)[i]&(*b)[i] != 0 {
			return true
		}
	}
	return false
}

// XOR集合是否非空
func Xor(a, b *Unit) bool {
	for i := 0; i < 8; i++ {
		if (*a)[i]^(*b)[i] != 0 {
			return true
		}
	}
	return false
}

// 减集是否非空
func Sub(a, b *Unit) bool {
	for i := 0; i < 8; i++ {
		if (*a)[i]&^(*b)[i] != 0 {
			return true
		}
	}
	return false
}
