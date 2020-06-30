package app

//创建栈

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value int) {
	s.element = append(s.element, value)
}

func (s *Stack) IsEmpty() bool {
	if len(s.element) == 0 {
		return true
	}
	return false
}

func (s *Stack) Pop() {
	s.element = s.element[:len(s.element)-1]
}

func (s *Stack) Top() (value int) {
	return s.element[len(s.element)-1]
}

func LongestValidParentheses(s string) int {
	res, start := 0, 0
	st := NewStack()
	for index := range s {
		if string(s[index]) == "(" {
			st.Push(index)
		} else if string(s[index]) == ")" {
			if st.IsEmpty() {
				start = index + 1
			} else {
				st.Pop()
				if st.IsEmpty() {
					res = max(res, index-start+1)
				} else {
					res = max(res, index-st.Top())
				}
			}
		}

	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
