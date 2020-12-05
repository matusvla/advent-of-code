package plane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlane(t *testing.T) {
	type args struct {
		firstSeat int
		lastSeat  int
	}
	tests := []struct {
		name      string
		args      args
		occupied  []int
		wantEmpty []int
	}{
		{
			name: "success",
			args: args{
				firstSeat: 5,
				lastSeat:  8,
			},
			occupied:  []int{7, 8},
			wantEmpty: []int{5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewPlane(tt.args.firstSeat, tt.args.lastSeat)
			for _, seat := range tt.occupied {
				p.MarkOccupied(seat)
			}
			assert.Equal(t, tt.wantEmpty, p.EmptySeats())
		})
	}
}
