package DFA

// 创建一个DFA正则
func NewDFA(s string) (*DFA, error) {
	tkn, err := translate([]byte(s))
	if err != nil {
		return nil, err
	}
	nfa, err := build(tkn)
	if err != nil {
		return nil, err
	}
	dfa, err := change(nfa, table(tkn))
	if err != nil {
		return nil, err
	}
	return dfa, nil
}

func (this *DFA) test(data []byte) int {
	i, j, k, l := 0, 0, -1, len(data)
	for j < l {
		c := data[j]
		k := this.Char[c]
		j++
		if i = this.Move[i][k]; i < 0 {
			break
		}
		if this.Over[i] {
			k = j
		}
	}
	return k
}

// 如匹配返回真
func (this *DFA) Match(data []byte) bool {
	for i := 0; i < len(data); i++ {
		j := this.test(data[i:])
		if j >= 0 {
			return true
		}
	}
	return false
}

// 如字符串匹配返回真
func (this *DFA) MatchString(data string) bool {
	return this.Match([]byte(data))
}

// 返回可能的第一个匹配
func (this *DFA) Find(data []byte) []byte {
	for i := 0; i < len(data); i++ {
		j := this.test(data[i:])
		if j >= 0 {
			return data[i : i+j]
		}
	}
	return nil
}

// 返回可能的第一个匹配的索引，类似regexp.Regexp.FindIndex
func (this *DFA) FindIndex(data []byte) []int {
	for i := 0; i < len(data); i++ {
		j := this.test(data[i:])
		if j >= 0 {
			return []int{i, i + j}
		}
	}
	return nil
}

// 返回最多n个可能的匹配，n<0时返回所有可能的匹配
func (this *DFA) FindAll(data []byte, n int) [][]byte {
	var s [][]byte
	for i := 0; i < len(data); {
		if len(s) == n {
			break
		}
		j := this.test(data[i:])
		if j >= 0 {
			s = append(s, data[i:i+j])
			i += j
		} else {
			i++
		}
	}
	return s
}

// 返回最多n个可能的匹配的索引，n<0时返回所有可能的匹配的索引
func (this *DFA) FindAllIndex(data []byte, n int) [][]int {
	var s [][]int
	for i := 0; i < len(data); {
		if len(s) == n {
			break
		}
		j := this.test(data[i:])
		if j >= 0 {
			s = append(s, []int{i, i + j})
			i += j
		} else {
			i++
		}
	}
	return s
}

// 返回字符串可能的第一个匹配
func (this *DFA) FindString(data string) string {
	s := this.Find([]byte(data))
	if s != nil {
		return string(s)
	}
	return ""
}

// 返回字符串可能的第一个匹配的索引
func (this *DFA) FindStringIndex(data string) []int {
	return this.FindIndex([]byte(data))
}

// 返回字符串最多n个可能的匹配，n<0时返回所有可能的匹配
func (this *DFA) FindAllString(data string, n int) []string {
	var s []string
	for i := 0; i < len(data); {
		if len(s) == n {
			break
		}
		j := this.test([]byte(data[i:]))
		if j >= 0 {
			s = append(s, data[i:i+j])
			i += j
		} else {
			i++
		}
	}
	return s
}

// 返回字符串最多n个可能的匹配的索引，n<0时返回所有可能的匹配的索引
func (this *DFA) FindAllStringIndex(data string, n int) [][]int {
	return this.FindAllIndex([]byte(data), n)
}
