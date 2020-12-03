package tobogan

import "testing"

func TestCheckTree(t *testing.T) {
	type args struct {
		line   string
		lineNo int
		step   float32
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no tree no bound",
			args: args{
				line:   "...#.",
				lineNo: 0,
				step:   3,
			},
			want: false,
		},
		{
			name: "tree no bound",
			args: args{
				line:   "...#.",
				lineNo: 1,
				step:   3,
			},
			want: true,
		},
		{
			name: "no tree bound",
			args: args{
				line:   "...#..",
				lineNo: 14,
				step:   3,
			},
			want: false,
		},
		{
			name: "tree bound",
			args: args{
				line:   "...#..",
				lineNo: 15,
				step:   3,
			},
			want: true,
		},
		{
			name: "not natural bound between",
			args: args{
				line:   "######",
				lineNo: 1,
				step:   0.5,
			},
			want: false,
		},
		{
			name: "not natural bound hit",
			args: args{
				line:   "######",
				lineNo: 2,
				step:   0.5,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckTree(tt.args.line, tt.args.lineNo, tt.args.step); got != tt.want {
				t.Errorf("CheckTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
