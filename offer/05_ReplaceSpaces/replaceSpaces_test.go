package replaceSpaces

import (
	"testing"
)

func TestReplaceSpaces(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello world", "hello%20world"},
		{" helloworld", "%20helloworld"},
		{"helloworld ", "helloworld%20"},
		{"hello  world", "hello%20%20world"},
		{"", ""},
		{" ", "%20"},
		{"helloworld", "helloworld"},
		{"   ", "%20%20%20"},
	}
	for i, tt := range tests {
		got := replaceSpaces(tt.input)
		if got != tt.want {
			t.Errorf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}
