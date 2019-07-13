package _1108_defanging_an_ip_address

import (
	"strings"
)

func defangIPaddr1(address string) string {
	// 替换 4 次
	return strings.Replace(address, ".", "[.]", 4)
}

func defangIPaddr2(address string) string {
	// 直接新开一个结果字符串重新构造
	var res string
	for _, v := range address {
		if v == '.' {
			res += "[.]"
		} else {
			res += string(v)
		}
	}
	return res
}
