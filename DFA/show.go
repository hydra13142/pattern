package DFA

import "fmt"

// 显示DFA内部的情况到标准输出
func (this *DFA) Show(t int) {
	var l int
	t %= 3
	switch t {
	case 0:
		l = 0
	case 1:
		l = len(this.Move[0]) - 1
	case 2:
		l = len(this.Move[0])
	}
	fmt.Printf("Char:\n")
	for t := 0; t < l; t++ {
		fmt.Printf("    [%02d] ", t)
		for i := 0; i < 256; i++ {
			if this.Char[i] != t {
				continue
			}
			if i > 32 && i < 127 {
				fmt.Printf("%c", i)
				continue
			}
			fmt.Printf("\\%02x", i)
		}
		fmt.Println()
	}
	fmt.Printf("table:\n    [ｘ]")
	for t := 0; t < len(this.Move[0]); t++ {
		fmt.Printf("[%02d]", t)
	}
	fmt.Println()
	for i, t := range this.Move {
		fmt.Printf("    [%02d] ", i)
		for _, c := range t {
			if c >= 0 {
				fmt.Printf("%02d  ", c)
			} else {
				fmt.Printf("--  ")
			}
		}
		fmt.Println()
	}
	fmt.Print("exit:\n    ")
	for t := 0; t < len(this.Over); t++ {
		if this.Over[t] {
			fmt.Printf("%d  ", t)
		}
	}
	fmt.Println()
}
