package findInPartiallySortedMatrix

import (
	"testing"
)

func TestFindInPartiallySortedMatrix(t *testing.T) {
	tests := []struct {
		input  [][]int
		target int
		want   bool
	}{
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 7, true},
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 5, false},
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 1, true},
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 15, true},
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 0, false},
		{[][]int{
			{1, 2, 8, 9},
			{2, 4, 9, 12},
			{4, 7, 10, 13},
			{6, 8, 11, 15},
		}, 16, false},
		{nil, 16, false},
	}
	for i, tt := range tests {
		got1 := findInPartiallySortedMatrix1(tt.input, tt.target)
		if got1 != tt.want {
			t.Errorf("%v. got1 %v, want %v", i, got1, tt.want)
		}

		got2 := findInPartiallySortedMatrix2(tt.input, tt.target)
		if got2 != tt.want {
			t.Errorf("%v. got2 %v, want %v", i, got2, tt.want)
		}
	}
}
