package string

import (
	"strconv"
)

func myAtoi(str string) int {
	const INT_MAX = int(^uint32(0) >> 1)
	const INT_MIN = ^INT_MAX
	sOp := ""
	sRes := ""
	for _, ch := range str {
		if '0' <= ch && ch <= '9' {
			sRes += string(ch)
			continue
		}
		if ch == '.' {
			break
		}
		if ch == '-' || ch == '+' {
			if len(sOp) != 0 {
				break
			}
			if len(sRes) == 0 {
				sOp = string(ch)
				continue
			}
		}
		if ch != ' ' {
			if len(sOp) == 0 && len(sRes) == 0 {
				return 0
			}
			break
		} else {
			if len(sRes) != 0 || len(sOp) != 0 {
				break
			}
		}
	}
	if len(sRes) == 0 {
		return 0
	}
	sRes = sOp + sRes
	iRes, _ := strconv.Atoi(sRes)
	if iRes > INT_MAX {
		return INT_MAX
	}
	if iRes < INT_MIN {
		return INT_MIN
	}
	return iRes
}
