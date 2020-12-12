package seatfiller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_newFloorplan(t *testing.T) {
	type args struct {
		b [][]byte
	}
	tests := []struct {
		name  string
		args  args
		wantF *floorplan
	}{
		{
			name: "success",
			args: args{
				b: [][]byte{
					[]byte("LL.L.L"),
					[]byte(".LL.L."),
					[]byte("LLLLLL"),
				},
			},
			wantF: &floorplan{
				seats: []point{
					{neighbours: 2},
					{neighbours: 3},
					{neighbours: 4, used: true},
					{neighbours: 2},
					{neighbours: 3, used: true},
					{neighbours: 1},

					{neighbours: 5, used: true},
					{neighbours: 6},
					{neighbours: 6},
					{neighbours: 6, used: true},
					{neighbours: 5},
					{neighbours: 4, used: true},

					{neighbours: 2},
					{neighbours: 4},
					{neighbours: 4},
					{neighbours: 4},
					{neighbours: 3},
					{neighbours: 2},
				},
				colLength:  6,
				nextToFill: []int{0, 1, 3, 4, 5, 12, 16, 17},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotF := newFloorplan(tt.args.b)
			assert.Equal(t, tt.wantF, gotF)
		})
	}
}

func TestProcessFloor(t *testing.T) {
	type args struct {
		b [][]byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "success",
		//	args: args{
		//		b: [][]byte{
		//			[]byte("LL.L.L"),
		//			[]byte(".LL.L."),
		//			[]byte("LLLLLL"),
		//		},
		//	},
		//	want: 37,
		//},
		{
			name: "success",
			args: args{
				b: [][]byte{
					[]byte("L.LL.LL.LL"),
					[]byte("LLLLLLL.LL"),
					[]byte("L.L.L..L.."),
					[]byte("LLLL.LL.LL"),
					[]byte("L.LL.LL.LL"),
					[]byte("L.LLLLL.LL"),
					[]byte("..L.L....."),
					[]byte("LLLLLLLLLL"),
					[]byte("L.LLLLLL.L"),
					[]byte("L.LLLLL.LL"),
				},
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessFloor(tt.args.b)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_floorplan_getCoordinates(t *testing.T) {
	type fields struct {
		seats      []point
		colLength  int
		nextToFill []int
	}
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "success",
			fields: fields{
				seats:     make([]point, 500),
				colLength: 15,
			},
			args: args{
				x: 12,
				y: 4,
			},
			want: 72,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &floorplan{
				seats:      tt.fields.seats,
				colLength:  tt.fields.colLength,
				nextToFill: tt.fields.nextToFill,
			}
			if got, _ := f.getCoordinates(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("getCoordinates() = %v, want %v", got, tt.want)
			}
		})
	}
}
