package _0_regular_expression_matching

import "testing"

type args struct {
	s string
	p string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{args: args{s: "", p:""}, want: true},
	{args: args{s: "", p:"*"}, want: false},
	{args: args{s: "a", p:".*"}, want: true},
	{args: args{s: "aaa", p:".*a"}, want: true},
	{args: args{s: "aab", p:"aab"}, want: true},
	{args: args{s: "aaa", p:"aa."}, want: true},
	{args: args{s: "aaa", p:"aa*"}, want: true},
	{args: args{s: "abc", p:"ab*"}, want: false},

}

func Test_isPalindrome_str(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}