package DFA

// 表示ASCII字符（范围0-255）的集合
type Unit []uint32

// 预定义字符组
var (
	Space = Unit{0x00003E00, 0x00000001, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	Digit = Unit{0x00000000, 0x03FF0000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	Alpha = Unit{0x00000000, 0x00000000, 0x07FFFFFE, 0x07FFFFFE, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
	Words = Unit{0x00000000, 0x03FF0000, 0x87FFFFFE, 0x07FFFFFE, 0x00000000, 0x00000000, 0x00000000, 0x00000000}
)

// 返回是否包含某个字符
func (this Unit) GetBit(i int) bool {
	return this[i>>5]&(1<<uint32(i&31)) != 0
}

// 设置包含某个字符
func (this Unit) SetBit(i int) Unit {
	this[i>>5] |= 1 << uint32(i&31)
	return this
}

// 设置不含某个字符
func (this Unit) ClsBit(i int) Unit {
	this[i>>5] &= ^(1 << uint32(i&31))
	return this
}

// 作拷贝
func (this Unit) Set(a Unit) Unit {
	for i, l := 0, len(this); i < l; i++ {
		this[i] = a[i]
	}
	return this
}

// 取反集
func (this Unit) Not(a Unit) Unit {
	for i, l := 0, len(this); i < l; i++ {
		this[i] = ^a[i]
	}
	return this
}

// 取并集
func (this Unit) Or(a, b Unit) Unit {
	for i, l := 0, len(this); i < l; i++ {
		this[i] = a[i] | b[i]
	}
	return this
}

// 取交集
func (this Unit) And(a, b Unit) Unit {
	for i, l := 0, len(this); i < l; i++ {
		this[i] = a[i] & b[i]
	}
	return this
}

// 取减集
func (this Unit) Sub(a, b Unit) Unit {
	for i, l := 0, len(this); i < l; i++ {
		this[i] = a[i] & (^b[i])
	}
	return this
}

// 取异或
func (this Unit) Xor(a, b Unit) Unit {
	for i, l := 0, len(this); i < l; i++ {
		this[i] = a[i] ^ b[i]
	}
	return this
}

// 是否包含全部字符
func Full(a Unit) bool {
	for i, l := 0, len(a); i < l; i++ {
		if a[i] != 0xffffffff {
			return false
		}
	}
	return true
}

// 是否没有任何字符
func None(a Unit) bool {
	for i, l := 0, len(a); i < l; i++ {
		if a[i] != 0x00000000 {
			return false
		}
	}
	return true
}

// 交集是否非空
func Same(a, b Unit) bool {
	for i, l := 0, len(a); i < l; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// 交集是否非空
func Both(a, b Unit) bool {
	for i, l := 0, len(a); i < l; i++ {
		if a[i]&b[i] != 0 {
			return true
		}
	}
	return false
}
