package replaceSpaces

func replaceSpaces(str string) string {
	if len(str) == 0 {
		return ""
	}

	oriStr := []byte(str)
	count := 0
	for _, v := range oriStr {
		if v == ' ' {
			count++
		}
	}
	newStr := make([]byte, len(oriStr)+2*count)
	copy(newStr, oriStr)
	oriLast := len(oriStr) - 1
	newLast := len(newStr) - 1
	for ; oriLast >= 0; oriLast-- {
		if newStr[oriLast] == ' ' {
			newStr[newLast] = '0'
			newLast--
			newStr[newLast] = '2'
			newLast--
			newStr[newLast] = '%'
			newLast--
		} else {
			newStr[newLast] = newStr[oriLast]
			newLast--
		}
	}
	return string(newStr)
}
