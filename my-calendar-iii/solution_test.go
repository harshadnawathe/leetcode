package mycalendariii

import (
	"reflect"
	"testing"
)

type cmdFunc func(*MyCalendarThree, ...any) any

var cmds = map[string]cmdFunc{
	"MyCalendarThree": func(m *MyCalendarThree, _ ...any) any {
		*m = Constructor()
		return nil
	},
	"book": func(cal *MyCalendarThree, args ...any) any {
		return cal.Book(args[0].(int), args[1].(int))
	},
}

func TestMyCalendarThree(t *testing.T) {
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
				[]string{"MyCalendarThree", "book", "book", "book", "book", "book", "book"},
				[][]any{
					{}, {10, 20}, {50, 60}, {10, 40}, {5, 15}, {5, 10}, {25, 55},
				},
			},
			[]any{nil, 1, 1, 2, 3, 3, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := []any{}
			var cal MyCalendarThree
			for i, cmd := range test.args.cmds {
				got = append(got, cmds[cmd](&cal, test.args.cmdArgs[i]...))
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("got: %v, want: %v", got, test.want)
			}
		})
	}
}
