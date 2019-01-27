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
	oriLastIndex := len(oriStr) - 1
	newLastIndex := len(newStr) - 1
	for ; oriLastIndex >= 0; oriLastIndex-- {
		if newStr[oriLastIndex] == ' ' {
			newStr[newLastIndex] = '0'
			newLastIndex--
			newStr[newLastIndex] = '2'
			newLastIndex--
			newStr[newLastIndex] = '%'
			newLastIndex--
		} else {
			newStr[newLastIndex] = newStr[oriLastIndex]
			newLastIndex--
		}
	}
	return string(newStr)
}
