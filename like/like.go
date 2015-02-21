package like

import (
	"github.com/hydra13142/pattern/token"
	"github.com/hydra13142/pattern/DFA"
	"strings"
)

type element struct {
	kind bool
	char byte
	unit DFA.Unit
}

// 采用like匹配的对象：'_'匹配1个任意字符，'%'匹配零到多个任意字符，支持'\w'格式预定义字符集和'[]'格式自定义字符集
type Like struct {
	Sense bool // 是否大小写敏感
	Whole bool // 是否匹配完整字符串
	rule  [][]*element
}

// 创建一个Like对象
func Compile(rule string) (*Like, error) {
	s := token.NewScanner(strings.NewReader(rule), token.Group, token.Escape, token.Byte)
	a, b := [][]*element{}, []*element{}
	for s.Next() {
		i, r := s.Token()
		if i == 0 {
			b = append(b, &element{kind: true, unit: r.(DFA.Unit)})
			continue
		}
		c := r.(byte)
		if i == 2 && c == '%' {
			a = append(a, b)
			b = []*element{}
			continue
		}
		b = append(b, &element{kind: false, char: c})
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	a = append(a, b)
	return &Like{true, false, a}, nil
}

func (r *Like) matchByte(e *element, b byte) bool {
	if e.kind {
		return e.unit.GetBit(int(b))
	}
	a := e.char
	if a == '_' || a == b {
		return true
	}
	if !r.Sense {
		if a >= 'A' && a <= 'Z' {
			return a+('a'-'A') == b
		}
		if a >= 'a' && a <= 'z' {
			return a-('a'-'A') == b
		}
	}
	return false
}

func (r *Like) matchString(a []*element, b string) int {
	l := len(a)
	if l == 0 {
		return 0
	}
	if l > len(b) {
		return -1
	}
	for i := 0; i < l; i++ {
		if !r.matchByte(a[i], b[i]) {
			return -1
		}
	}
	return l
}

func (r *Like) matchSlice(i int, s string) int {
	var x, y, z int
	x = r.matchString(r.rule[i], s)
	if x < 0 {
		return -1
	}
	i++
	if i >= len(r.rule) {
		if !r.Whole || x == len(s) {
			return x
		}
		return -1
	}
	for y = x; y < len(s); y++ {
		z = r.matchSlice(i, s[y:])
		if z >= 0 {
			return y + z
		}
	}
	return -1
}

// 进行匹配
func (r *Like) Match(s string) bool {
	if r.Whole {
		return r.matchSlice(0, s) >= 0
	}
	for i := 0; i < len(s); i++ {
		if r.matchSlice(0, s[i:]) >= 0 {
			return true
		}
	}
	return false
}
