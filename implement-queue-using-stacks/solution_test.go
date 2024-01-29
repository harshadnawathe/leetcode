package implementqueueusingstacks

import (
	"reflect"
	"testing"
)

func TestMyQueue(t *testing.T) {
	type args struct {
		cmds    []string
		cmdArgs [][]any
	}

	tests := []struct {
		name string
		args args
		want []any
	}{
		{
			"Example 1",
			args{
				[]string{"MyQueue", "push", "push", "peek", "pop", "empty"},
				[][]any{{}, {1}, {2}, {}, {}, {}},
			},
			[]any{nil, nil, nil, 1, 1, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := []any{}
			var q MyQueue
			for i := range tt.args.cmds {
				got = append(got, cmdFuncs[tt.args.cmds[i]](&q, tt.args.cmdArgs[i]...))
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

type cmdFunc func(*MyQueue, ...any) any

var cmdFuncs = map[string]cmdFunc{
	"MyQueue": func(q *MyQueue, arg ...any) any {
		*q = Constructor()
		return nil
	},
	"push": func(q *MyQueue, arg ...any) any {
		q.Push(arg[0].(int))
		return nil
	},
	"pop":   func(q *MyQueue, arg ...any) any { return q.Pop() },
	"peek":  func(q *MyQueue, arg ...any) any { return q.Peek() },
	"empty": func(q *MyQueue, arg ...any) any { return q.Empty() },
}
