package convolution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvolver3D_ConvolutionExtendBounds(t *testing.T) {
	type fields struct {
		mask [3][3][3]int
	}
	type args struct {
		in [][][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   [][][]int
	}{
		{
			name: "success",
			fields: fields{
				mask: [3][3][3]int{
					{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
					{
						{1, 1, 1},
						{1, 0, 1},
						{1, 1, 1},
					},
					{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
				},
			},
			args: args{
				in: [][][]int{
					{
						{0, 1, 0},
						{0, 0, 1},
						{1, 1, 1},
					},
				},
			},
			want: [][][]int{
				{
					{0, 1, 1, 1, 0},
					{0, 1, 2, 2, 1},
					{1, 3, 5, 4, 2},
					{1, 2, 4, 3, 2},
					{1, 2, 3, 2, 1},
				},
				{
					{0, 1, 1, 1, 0},
					{0, 1, 1, 2, 1},
					{1, 3, 5, 3, 2},
					{1, 1, 3, 2, 2},
					{1, 2, 3, 2, 1},
				},
				{
					{0, 1, 1, 1, 0},
					{0, 1, 2, 2, 1},
					{1, 3, 5, 4, 2},
					{1, 2, 4, 3, 2},
					{1, 2, 3, 2, 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Convolver3D{
				mask: tt.fields.mask,
			}
			got := c.ConvolutionExtendBoundsWithMod(tt.args.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
