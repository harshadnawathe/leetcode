package removeoutermostparentheses

import "testing"

func Test_removeOuterParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				s: "(()())(())",
			},
			want: "()()()",
		},
		{
			name: "Example 2",
			args: args{
				s: "(()())(())(()(()))",
			},
			want: "()()()()(())",
		},
		{
			name: "Example 3",
			args: args{
				s: "()()",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeOuterParentheses(tt.args.s); got != tt.want {
				t.Errorf("removeOuterParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
