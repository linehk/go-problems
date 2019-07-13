package _1108_defanging_an_ip_address

import (
	"testing"
)

func Test_defangIPaddr1(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}
	for i, tt := range tests {
		if got := defangIPaddr1(tt.input); got != tt.want {
			t.Fatalf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

// Benchmark_defangIPaddr1-8   	 5000000	       235 ns/op
func Benchmark_defangIPaddr1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defangIPaddr1("1.1.1.1")
		defangIPaddr1("255.100.50.0")
	}
}

func Test_defangIPaddr2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"255.100.50.0", "255[.]100[.]50[.]0"},
	}
	for i, tt := range tests {
		if got := defangIPaddr2(tt.input); got != tt.want {
			t.Fatalf("%v. got %v, want %v", i, got, tt.want)
		}
	}
}

// Benchmark_defangIPaddr2-8   	 2000000	       709 ns/op
func Benchmark_defangIPaddr2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defangIPaddr2("1.1.1.1")
		defangIPaddr2("255.100.50.0")
	}
}
