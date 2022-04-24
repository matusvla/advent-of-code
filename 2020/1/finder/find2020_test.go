package finder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindAndMultiply2020(t *testing.T) {
	type args struct {
		expReport []int
	}
	type want struct {
		result int
		found  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				expReport: []int{2, 10, 54, 67, 2010, 456},
			},
			want: want{
				result: 20100,
				found:  true,
			},
		},
		{
			name: "fail",
			args: args{
				expReport: []int{2, 10, 54, 67, 2011, 456},
			},
			want: want{
				result: 0,
				found:  false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := findAndMultiplyTwo(tt.args.expReport, 2020)
			assert.Equal(t, tt.want.result, result)
			assert.Equal(t, tt.want.found, found)
		})
	}
}

func TestFindAndMultiply(t *testing.T) {
	type args struct {
		expReport   []int
		sumResult   int
		numberCount int
	}
	type want struct {
		result int
		found  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success 3",
			args: args{
				expReport:   []int{15, 2, 19, 1, 17},
				sumResult:   20,
				numberCount: 3,
			},
			want: want{
				result: 34,
				found:  true,
			},
		},
		{
			name: "success 4",
			args: args{
				expReport:   []int{1, 2, 3, 14},
				sumResult:   20,
				numberCount: 4,
			},
			want: want{
				result: 84,
				found:  true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := FindAndMultiply(tt.args.expReport, tt.args.sumResult, tt.args.numberCount)
			assert.Equal(t, tt.want.result, result)
			assert.Equal(t, tt.want.found, found)
		})
	}
}
