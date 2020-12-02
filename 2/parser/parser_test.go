package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		passLine string
		policy   func(int, int, string, string) bool
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				passLine: "1-3 a: abcde",
				policy:   OldPolicy,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				passLine: "1-3 a: bcde",
				policy:   OldPolicy,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				passLine: "1- a: bcde",
				policy:   OldPolicy,
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "success new",
			args: args{
				passLine: "1-3 a: abcde",
				policy:   NewPolicy,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "success new 2",
			args: args{
				passLine: "1-3 a: cbade",
				policy:   NewPolicy,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "fail new none",
			args: args{
				passLine: "1-3 a: bcde",
				policy:   NewPolicy,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "fail new both",
			args: args{
				passLine: "1-3 a: abacde",
				policy:   NewPolicy,
			},
			want:    false,
			wantErr: false,
		},
		{
			name: "error new",
			args: args{
				passLine: "1- a: bcde",
				policy:   NewPolicy,
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateLine(tt.args.passLine, tt.args.policy)
			assert.Equal(t, tt.wantErr, err != nil, "error")
			assert.Equal(t, tt.want, got, "result")
		})
	}
}
