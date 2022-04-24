package main

import (
	"reflect"
	"testing"
)

func Test_listPossible(t *testing.T) {
	type args struct {
		x     int
		moves int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "asdf",
			args: args{
				x:     20,
				moves: 1,
			},
			want: []int{20},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listPossible(tt.args.x, tt.args.moves); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
