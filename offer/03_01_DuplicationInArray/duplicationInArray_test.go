package duplicationInArray

import (
	"testing"
)

func TestDuplicationInArray(t *testing.T) {
	tests := []struct {
		input []int
		want  []int
		ok    bool
	}{
		// 示例
		{[]int{2, 3, 1, 0, 2, 5, 3}, []int{2, 3}, true},
		// 重复数字是数组中最小
		{[]int{2, 1, 3, 1, 4}, []int{1}, true},
		// 重复数字是数组中最大
		{[]int{2, 4, 3, 1, 4}, []int{4}, true},
		// 多个重复数字
		{[]int{2, 4, 2, 1, 4}, []int{2, 4}, true},
		// 无重复
		{[]int{2, 1, 3, 0, 4}, []int{0}, false},
		{nil, []int{0}, false},
	}
	for i, tt := range tests {
		got1, ok1 := duplicationInArray1(tt.input)
		if ok1 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok1, tt.ok)
		}
		if !contains(tt.want, got1) {
			t.Errorf("%v. got %v, want %v", i, got1, tt.want)
		}

		got2, ok2 := duplicationInArray2(tt.input)
		if ok2 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok2, tt.ok)
		}
		if !contains(tt.want, got2) {
			t.Errorf("%v. got %v, want %v", i, got2, tt.want)
		}

		got3, ok3 := duplicationInArray3(tt.input)
		if ok3 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok3, tt.ok)
		}
		if !contains(tt.want, got3) {
			t.Errorf("%v. got %v, want %v", i, got3, tt.want)
		}

		got4, ok4 := duplicationInArray4(tt.input)
		if ok4 != tt.ok {
			t.Errorf("%v. got %v, want %v", i, ok4, tt.ok)
		}
		if !contains(tt.want, got4) {
			t.Errorf("%v. got %v, want %v", i, got4, tt.want)
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
