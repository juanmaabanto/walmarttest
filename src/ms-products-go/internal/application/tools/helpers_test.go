package tools

import "testing"

func TestEsPalindrome(t *testing.T) {
	type args struct {
		valor string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Es palíndromo", args{"somos"}, true},
		{"No es palíndromo", args{"casa"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EsPalindrome(tt.args.valor); got != tt.want {
				t.Errorf("EsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
