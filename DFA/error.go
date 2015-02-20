package DFA

import "fmt"

// 词法错误，其值表示出错位置（字节）
type LexicalError int

// 实现error接口
func (this LexicalError) Error() string {
	return fmt.Sprintf("DFA lexical error at byte %d", this)
}

// 语法错误，其值表示出错位置（token）
type SyntaxError int

// 实现error接口
func (this SyntaxError) Error() string {
	return fmt.Sprintf("DFA syntax error at Token %d", this)
}
