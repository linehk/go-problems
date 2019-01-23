package duplicationInArrayNoEdit

import (
	"testing"
)

func TestDuplicationInArrayNoEdit(t *testing.T) {
	tests := []struct {
		input    []int
		want     []int
		wantBool bool
	}{
		{[]int{2, 3, 5, 4, 3, 2, 6, 7}, []int{2, 3}, true},
	}
	for i, tt := range tests {
		got1, ok1 := duplicationInArrayNoEdit1(tt.input)
		if ok1 != tt.wantBool {
			t.Errorf("%v. got1 %v, want %v", i, ok1, tt.wantBool)
		}
		f1 := false
		for _, v := range tt.want {
			if got1 == v {
				f1 = true
			}
		}
		if !f1 {
			t.Errorf("%v. got1 %v, want %v", i, got1, tt.want)
		}

		got2, ok2 := duplicationInArrayNoEdit2(tt.input)
		if ok2 != tt.wantBool {
			t.Errorf("%v. got2 %v, want %v", i, ok2, tt.wantBool)
		}
		f2 := false
		for _, v := range tt.want {
			if got2 == v {
				f2 = true
			}
		}
		if !f2 {
			t.Errorf("%v. got2 %v, want %v", i, got2, tt.want)
		}

		got3, ok3 := duplicationInArrayNoEdit3(tt.input)
		if ok3 != tt.wantBool {
			t.Errorf("%v. got3 %v, want %v", i, ok3, tt.wantBool)
		}
		f3 := false
		for _, v := range tt.want {
			if got3 == v {
				f3 = true
			}
		}
		if !f3 {
			t.Errorf("%v. got3 %v, want %v", i, got3, tt.want)
		}
	}
}
