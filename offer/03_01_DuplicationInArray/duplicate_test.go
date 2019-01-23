package duplicate

import (
	"testing"
)

func TestDuplicate1(t *testing.T) {
	tests := []struct {
		input    []int
		want     int
		wantBool bool
	}{
		{[]int{2, 3, 1, 0, 2, 5, 3}, 2, true},
	}
	for i, tt := range tests {
		got1, ok1 := duplicate1(tt.input)
		if !ok1 {
			t.Errorf("%v. got1 %v, want %v", i, ok1, tt.wantBool)
		}
		if got1 != tt.want {
			t.Errorf("%v. got1 %v, want %v", i, got1, tt.want)
		}

		got2, ok2 := duplicate2(tt.input)
		if !ok2 {
			t.Errorf("%v. got2 %v, want %v", i, ok2, tt.wantBool)
		}
		if got2 != tt.want {
			t.Errorf("%v. got2 %v, want %v", i, got2, tt.want)
		}

		got3, ok3 := duplicate3(tt.input)
		if !ok3 {
			t.Errorf("%v. got3 %v, want %v", i, ok3, tt.wantBool)
		}
		if got3 != tt.want {
			t.Errorf("%v. got3 %v, want %v", i, got3, tt.want)
		}

		got4, ok4 := duplicate4(tt.input)
		if !ok4 {
			t.Errorf("%v. got4 %v, want %v", i, ok4, tt.wantBool)
		}
		if got4 != tt.want {
			t.Errorf("%v. got4 %v, want %v", i, got4, tt.want)
		}
	}
}
