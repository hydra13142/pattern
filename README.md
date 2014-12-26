pattern
=======

PACKAGE DOCUMENTATION

	package token
		import "github.com/hydra13142/pattern/token"

		token包用于方便的进行token的提取，让使用者专注于语法层面。
		本包提供了大量预定义的token提取规则函数，可以提取多种token，也可以自定义token提取函数。
		使用io.Reader和多个token提取函数，即可创建一个token扫描器，可以依次提取出token来。

FUNCTIONS

	func Alternative(takes ...func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{})
		选择其一匹配

	func And(a, b *Unit) bool
		交集是否非空

	func Byte(data []byte, end bool) (int, interface{})
		所有其余未转义的ASCII字符

	func Class(data []byte, end bool) (int, interface{})
		字符组（预定义、自定义）

	func Dec(s []byte, b, l int) (i int, n int)
		十进制整数

	func Edge(data []byte, end bool) (int, interface{})
		边界

	func Escape(data []byte, end bool) (int, interface{})
		预定义的转义字符、十六进制表示的ASCII字符

	func Field(data []byte, end bool) (int, interface{})
		连续可打印ASCII字符

	func Float(data []byte, end bool) (i int, r interface{})
		浮点数

	func Group(data []byte, end bool) (int, interface{})
		\num形式的子匹配

	func Hex(s []byte, b, l int) (i int, n int)
		十六进制整数

	func Integer(data []byte, end bool) (i int, r interface{})
		整数

	func Limit(data []byte, end bool) (int, interface{})
		限定重复次数

	func Oct(s []byte, b, l int) (i int, n int)
		八进制整数

	func QuoteByte(data []byte, end bool) (int, interface{})
		单字符（单引号包围）

	func QuoteString(data []byte, end bool) (i int, q interface{})
		字符串（双引号或反引号包围）

	func Space(data []byte, end bool) (int, interface{})
		空白

	func Sub(a, b *Unit) bool
		减集是否非空

	func Tolerate(head, tail string, take func([]byte, bool) (int, interface{})) func([]byte, bool) (int, interface{})
		顺序匹配

	func Var(data []byte, end bool) (int, interface{})
		$name，${name}形式的变量

	func Word(data []byte, end bool) (int, interface{})
		标识符

	func Xor(a, b *Unit) bool
		XOR集合是否非空

TYPES

	type Scanner struct {
		*State
		// contains filtered or unexported fields
	}
		token扫描器

	func NewScanner(r io.Reader, f ...TokenFunc) *Scanner
		创建token扫描器，r为数据来源，f为规则列表

	func (s *Scanner) Next() bool
		提取下一个token，必须在Token方法（包括第一次）之前调用

	func (s *Scanner) Reset(r io.Reader)
		重设数据源，以便复用扫描器

	type State struct {
		io.Reader
		// contains filtered or unexported fields
	}
		保管扫描器状态

	func (s *State) Err() error
		返回可能的错误（不完整/不匹配），到达文档末尾不返回错误

	func (s *State) Token() (int, interface{})
		返回下一个token，第一个返回值为采用的规则的索引，第二个为规则产生的输出

	type TokenFunc func([]byte, bool) (int, interface{})
		提取token的函数类型，第一个返回值为消费掉的字节数

	type Unit [8]uint32
		表示ASCII字符（范围0-255）的集合

	func (x *Unit) And(a, b *Unit) *Unit
		交集

	func (a *Unit) Cls(i byte)
		设置不含某个字符

	func (a *Unit) Full() bool
		是否包含全部字符

	func (a *Unit) Get(i byte) bool
		返回是否包含某个字符

	func (a *Unit) None() bool
		是否没有任何字符

	func (x *Unit) Not(a *Unit) *Unit
		取反集

	func (x *Unit) Or(a, b *Unit) *Unit
		并集

	func (a *Unit) Set(i byte)
		设置包含某个字符

	func (x *Unit) Sub(a, b *Unit) *Unit
		减去

	func (x *Unit) Xor(a, b *Unit) *Unit
		非共有字符的集合
