package crt

import "testing"

func TestChineseRemainderTheorem2Equations(t *testing.T) {
	type args struct {
		a1 int64
		a2 int64
		n1 int64
		n2 int64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				a1: 6,
				a2: 2,
				n1: 10,
				n2: 7,
			},
			want:    16,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChineseRemainderTheorem2Equations(tt.args.a1, tt.args.a2, tt.args.n1, tt.args.n2)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChineseRemainderTheorem2Equations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChineseRemainderTheorem2Equations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChineseRemainderTheorem(t *testing.T) {
	type args struct {
		a []int
		n []int
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				a: []int{0, 2, 659 - 41, 23 - 3, 13 - 2, 19 - 3, 29 - 12, 937 - 72, 17 - 4},
				n: []int{41, 37, 659, 23, 13, 19, 29, 937, 17},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ChineseRemainderTheorem(tt.args.a, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChineseRemainderTheorem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ChineseRemainderTheorem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
