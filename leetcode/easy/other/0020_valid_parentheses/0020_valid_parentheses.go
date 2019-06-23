package other

func isValid(s string) bool {
	l := len(s)
	// 用一个栈来存储对应括号
	stack := make([]byte, l)
	top := 0
	for _, v := range s {
		switch v {
		case '(':
			stack[top] = ')'
			top++
		case '[':
			stack[top] = ']'
			top++
		case '{':
			stack[top] = '}'
			top++
		case ')', ']', '}':
			// 当遇到对应括号时，检测是否与栈顶元素相等
			if top > 0 && stack[top-1] == byte(v) {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}
