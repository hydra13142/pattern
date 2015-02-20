package glob

import (
	"github.com/hydra13142/pattern"
	"strings"
)

// 采用glob匹配的对象：'?'匹配1个任意字符，'*'匹配零到多个任意字符
type Glob struct {
	Sense bool // 是否大小写敏感
	Whole bool // 是否匹配完整字符串
	rule  []string
}

// 创建一个Glob对象
func Compile(rule string) (*Glob, error) {
	s := pattern.NewScanner(strings.NewReader(rule), pattern.Escape, pattern.Byte)
	a, b := []string{}, []byte{}
	for s.Next() {
		i, r := s.Token()
		c := r.(byte)
		if i == 1 && c == '*' {
			a = append(a, string(b))
			b = []byte{}
		} else {
			b = append(b, c)
		}
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	a = append(a, string(b))
	return &Glob{true, false, a}, nil
}

func (g *Glob) matchByte(a, b byte) bool {
	if a == '?' || a == b {
		return true
	}
	if !g.Sense {
		if a >= 'A' && a <= 'Z' {
			return a+('a'-'A') == b
		}
		if a >= 'a' && a <= 'z' {
			return a-('a'-'A') == b
		}
	}
	return false
}

func (g *Glob) matchString(a, b string) int {
	l := len(a)
	if l == 0 {
		return 0
	}
	if l > len(b) {
		return -1
	}
	for i := 0; i < l; i++ {
		if !g.matchByte(a[i], b[i]) {
			return -1
		}
	}
	return l
}

func (g *Glob) matchSlice(i int, s string) int {
	var x, y, z int
	x = g.matchString(g.rule[i], s)
	if x < 0 {
		return -1
	}
	i++
	if i >= len(g.rule) {
		if !g.Whole || x == len(s) {
			return x
		}
		return -1
	}
	for y = x; y < len(s); y++ {
		z = g.matchSlice(i, s[y:])
		if z >= 0 {
			return y + z
		}
	}
	return -1
}

// 进行匹配
func (g *Glob) Match(s string) bool {
	if g.Whole {
		return g.matchSlice(0, s) >= 0
	}
	for i := 0; i < len(s); i++ {
		if g.matchSlice(0, s[i:]) >= 0 {
			return true
		}
	}
	return false
}
