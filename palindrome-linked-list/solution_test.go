package palindromelinkedlist

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				head: listOf(1, 2, 2, 1),
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				head: listOf(1, 2),
			},
			want: false,
		},
		{
			name: "Palindrome with odd length",
			args: args{
				head: listOf(1, 2, 3, 2, 1),
			},
			want: true,
		},
		{
			name: "Non-palindrome with odd length",
			args: args{
				head: listOf(1, 2, 3),
			},
			want: false,
		},
		{
			name: "List with single node",
			args: args{
				head: listOf(1),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
