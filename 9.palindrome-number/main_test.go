package __palindrome_number

import "testing"

type args struct {
	x int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{args: args{0}, want: true},
	{args: args{-1}, want: false},
	{args: args{10}, want: false},
	{args: args{101}, want: true},
	{args: args{11}, want: true},
	{args: args{9222999999999992229}, want: true},
}

func Test_isPalindrome_str(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeStr(tt.args.x); got != tt.want {
				t.Errorf("isPalindromeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
