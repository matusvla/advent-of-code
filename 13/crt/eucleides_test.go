package crt

import (
	"math/big"
	"testing"
)

func TestExtendedEucleides(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	tests := []struct {
		name  string
		args  args
		want  *big.Int
		want1 *big.Int
		want2 *big.Int
	}{
		{
			name: "success",
			args: args{
				a: 7,
				b: 13,
			},
			want:  big.NewInt(1),
			want1: big.NewInt(2),
			want2: big.NewInt(-1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := ExtendedEucleides(big.NewInt(tt.args.a), big.NewInt(tt.args.b))
			if got != tt.want {
				t.Errorf("ExtendedEucleides() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ExtendedEucleides() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("ExtendedEucleides() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
