package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrammar_findPossibileCombinations(t *testing.T) {
	type fields struct {
		rules map[string][]string
	}
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]struct{}
	}{
		{
			name: "success",
			fields: fields{
				rules: map[string][]string{
					"A,B": {"S", "C"},
					"B,C": {"S"},
					"B,A": {"A"},
					"C,C": {"B"},
					"a":   {"A,C"},
					"b":   {"B"},
				},
			},
			args: args{
				a: []string{"A", "C"},
				b: []string{"S", "C"},
			},
			want: map[string]struct{}{"B": {}},
		},
		{
			name: "empty",
			fields: fields{
				rules: map[string][]string{
					"A,B": {"S", "C"},
					"B,C": {"S"},
					"B,A": {"A"},
					"C,C": {"B"},
					"a":   {"A,C"},
					"b":   {"B"},
				},
			},
			args: args{
				a: []string{"A", "S"},
				b: []string{"A", "C"},
			},
			want: map[string]struct{}{},
		},
		{
			name: "empty",
			fields: fields{
				rules: map[string][]string{
					"A,B": {"S", "C"},
					"B,C": {"S"},
					"B,A": {"A"},
					"C,C": {"B"},
					"a":   {"A,C"},
					"b":   {"B"},
				},
			},
			args: args{
				a: []string{"A", "B"},
				b: []string{"B", "C"},
			},
			want: map[string]struct{}{"C": {}, "S": {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grammar{
				rules: tt.fields.rules,
			}
			got := g.findPossibileCombinations(tt.args.a, tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getTrinagleIndex(t *testing.T) {
	type args struct {
		i           int
		j           int
		firstRowLen int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				i:           2,
				j:           3,
				firstRowLen: 7,
			},
			want: 20,
		},
		{
			name: "success",
			args: args{
				i:           0,
				j:           2,
				firstRowLen: 6,
			},
			want: 11,
		},
		{
			name: "success first",
			args: args{
				i:           0,
				j:           0,
				firstRowLen: 100,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTrinagleIndex(tt.args.i, tt.args.j, tt.args.firstRowLen); got != tt.want {
				t.Errorf("getTrinagleIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrammar_ContainsWord(t *testing.T) {
	type fields struct {
		rules map[string][]string
	}
	type args struct {
		startingPoint string
		terminalWord  string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "success",
			fields: fields{
				rules: map[string][]string{
					"A,B": {"S", "C"},
					"B,C": {"S"},
					"B,A": {"A"},
					"C,C": {"B"},
					"a":   {"A", "C"},
					"b":   {"B"},
				},
			},
			args: args{
				startingPoint: "S",
				terminalWord:  "baaba",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Grammar{
				rules: tt.fields.rules,
			}
			if got := g.ContainsWord(tt.args.startingPoint, tt.args.terminalWord); got != tt.want {
				t.Errorf("ContainsWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
