package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvaluateExpression(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success simple",
			args: args{
				e: "1 + 4*8+9",
			},
			want: 49,
		},
		{
			name: "success simple brackets",
			args: args{
				e: "1 + (4*8)+9",
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateExpression(tt.args.e); got != tt.want {
				t.Errorf("evaluateExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluate(t *testing.T) {
	type args struct {
		operand  *string
		result   *int
		operator func(int, int) int
	}
	type want struct {
		result  int
		operand string
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "success",
			args: args{
				operand:  func() *string { a := "15"; return &a }(),
				result:   func() *int { a := 0; return &a }(),
				operator: plus,
			},
			want: want{
				result:  15,
				operand: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			evaluate(tt.args.operand, tt.args.result, tt.args.operator)
			assert.Equal(t, tt.want.result, *tt.args.result)
			assert.Equal(t, tt.want.operand, *tt.args.operand)
		})
	}
}

func TestEvaluateExpression2(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				e: "1 + (2 + 3 * 5) * 2 * (1+1)",
			},
			want: 26 * 2 * 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EvaluateExpression2(tt.args.e); got != tt.want {
				t.Errorf("EvaluateExpression2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_evaluateExpression2(t *testing.T) {
	type args struct {
		e string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name:  "success",
		//	args:  args{
		//		e: "1 + ((2 + 3) + 4) + 5",
		//	},
		//	want:  15,
		//},
		//{
		//	name:  "success",
		//	args:  args{
		//		e: "1 + ((2 + 3) + 4) + 5 * 1 + ((2 + 3) + 4) + 5",
		//	},
		//	want:  15*15,
		//},
		{
			name: "success",
			args: args{
				e: "71 + 5 + (9 + (7 + 8) * (5 + 3) + 7)",
			},
			want: 15 * 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvaluateExpression2(tt.args.e)
			if got != tt.want {
				t.Errorf("evaluateExpression2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
