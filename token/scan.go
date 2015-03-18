// token包用于方便的进行token的提取，让使用者专注于语法层面。
// 本包提供了大量预定义的token提取规则函数，可以提取多种token，也可以自定义token提取函数。
// 使用io.Reader和多个token提取函数，即可创建一个token扫描器，可以依次提取出token来。
package token

import (
	"errors"
	"io"
)

// 提取token的函数类型，第一个返回值>0表示成功匹配吃掉的字节数，=0表示字节不够，<0表示匹配失败
type TokenFunc func([]byte, bool) (int, interface{})

// 保管扫描器状态
type State struct {
	idx int
	vlu interface{}
	dat []byte
	err error
	io.Reader
}

// token扫描器
type Scanner struct {
	fac []TokenFunc
	*State
}

// 创建token扫描器，r为数据来源，f为规则列表
func NewScanner(r io.Reader, f ...TokenFunc) *Scanner {
	return &Scanner{f, &State{0, nil, nil, nil, r}}
}

// 重设数据源，以便复用扫描器
func (s *Scanner) Reset(r io.Reader) {
	s.Reader = r
	s.dat = nil
	s.err = nil
}

// 提取下一个token，必须在Token方法（包括第一次）之前调用
func (s *Scanner) Next() bool {
	var (
		d = make([]byte, 1024)
		p = s.err == io.EOF
		n int
	)
	if len(s.dat) == 0 {
		if s.err != nil {
			return false
		}
		n, s.err = s.Read(d)
		p = s.err == io.EOF
		if n == 0 {
			return false
		}
		s.dat = append(s.dat, d[:n]...)
	}
	for i, l := 0, len(s.fac); ; {
		for i = 0; i < l; i++ {
			n, s.vlu = s.fac[i](s.dat, p)
			if n == 0 {
				break
			}
			if n > 0 {
				s.idx = i
				s.dat = s.dat[n:]
				return true
			}
		}
		if i == l {
			break
		}
		if s.err != nil {
			s.err = errors.New("incompleted token")
			return false
		}
		n, s.err = s.Read(d)
		if n == 0 {
			s.err = errors.New("incompleted token")
			return false
		}
		s.dat = append(s.dat, d[:n]...)
		p = (s.err == io.EOF)
	}
	s.err = errors.New("can't identify token")
	return false
}

// 返回下一个token，第一个返回值为采用的规则的索引，第二个为规则产生的输出
func (s *State) Token() (int, interface{}) {
	return s.idx, s.vlu
}

// 返回可能的错误（不完整/不匹配），到达文档末尾不返回错误
func (s *State) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}
