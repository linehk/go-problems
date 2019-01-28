package other

func isValid(s string) bool {
	l := len(s)
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
			if top > 0 && stack[top-1] == byte(v) {
				top--
			} else {
				return false
			}
		}
	}
	return top == 0
}
