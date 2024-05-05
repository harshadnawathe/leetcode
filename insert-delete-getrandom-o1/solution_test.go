package insertdeletegetrandomo1

import (
	"reflect"
	"testing"
)

type cmdFunc func(*RandomizedSet, ...any) any

var cmds = map[string]cmdFunc{
	"RandomizedSet": func(s *RandomizedSet, args ...any) any {
		*s = Constructor()
		return nil
	},
	"insert": func(s *RandomizedSet, args ...any) any {
		return s.Insert(args[0].(int))
	},
	"remove": func(s *RandomizedSet, args ...any) any {
		return s.Remove(args[0].(int))
	},
	"getRandom": func(s *RandomizedSet, args ...any) any {
		return s.GetRandom()
	},
}

func Test_RandomizedSet(t *testing.T) {

	t.Skip("Skipping flakey test")

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
			name: "Example 1",
			args: args{
				cmds:    []string{"RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"},
				cmdArgs: [][]any{{}, {1}, {2}, {2}, {}, {1}, {2}, {}},
			},
			want: []any{nil, true, false, true, 2, true, false, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := []any{}
			var set RandomizedSet
			for i, cmd := range test.args.cmds {
				got = append(got, cmds[cmd](&set, test.args.cmdArgs[i]...))
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
