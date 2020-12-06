package answers

import "testing"

func TestIntersection(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				s: "abcd\nbcnjd\nbiedk",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersection(tt.args.s); got != tt.want {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
