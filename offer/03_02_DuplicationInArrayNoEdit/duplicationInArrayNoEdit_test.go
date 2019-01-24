package duplicationInArrayNoEdit

import (
	"testing"
)

func TestDuplicationInArrayNoEdit(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
		ok    bool
	}{
		// 多个重复数字
		{[]int{2, 3, 5, 4, 3, 2, 6, 7}, []int{2, 3}, true},
		// 一个
		{[]int{3, 2, 1, 4, 4, 5, 6, 7}, []int{4}, true},
		// 重复数字是数组中最小
		{[]int{1, 2, 3, 4, 5, 6, 7, 1, 8}, []int{1}, true},
		// 重复数字数数组中最大
		{[]int{1, 7, 3, 4, 5, 6, 8, 2, 8}, []int{8}, true},
		// 两个数字的数组
		{[]int{1, 1}, []int{1}, true},
		// 重复数字在数组中间
		{[]int{3, 2, 1, 3, 4, 5, 6, 7}, []int{3}, true},
		// 多个
		{[]int{1, 2, 2, 6, 4, 5, 6}, []int{2, 6}, true},
		// 一个数字重复三次
		{[]int{1, 2, 2, 6, 4, 5, 2}, []int{2}, true},
		{nil, []int{0}, false},
	}
	for i, tt := range tests {
		got1, ok1 := duplicationInArrayNoEdit1(tt.input)
		if ok1 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok1, tt.ok)
		}
		if !contains(tt.want, got1) {

			t.Errorf("%v. got %v, want %v", i, got1, tt.want)
		}

		got2, ok2 := duplicationInArrayNoEdit2(tt.input)
		if ok2 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok2, tt.ok)
		}
		if !contains(tt.want, got2) {
			t.Errorf("%v. got %v, want %v", i, got2, tt.want)
		}

		got3, ok3 := duplicationInArrayNoEdit3(tt.input)
		if ok3 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok3, tt.ok)
		}
		if !contains(tt.want, got3) {
			t.Errorf("%v. got %v, want %v", i, got3, tt.want)
		}
	}
}

func contains(want []int, got int) bool {
	f := false
	for _, v := range want {
		if got == v {
			f = true
		}
	}
	return f
}
