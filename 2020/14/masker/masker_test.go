package masker

import (
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMasker_SetMask(t *testing.T) {
	type fields struct {
		mask         int64
		maskedFields int64
	}
	type args struct {
		s string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantFields fields
	}{
		{
			name: "success",
			fields: fields{
				mask:         50,
				maskedFields: 12345,
			},
			args: args{
				s: "10X01",
			},
			wantFields: fields{
				mask:         17,
				maskedFields: 27,
			},
		},
		{
			name: "success bigger",
			fields: fields{
				mask:         50,
				maskedFields: 12345,
			},
			args: args{
				s: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
			},
			wantFields: fields{
				mask:         64,
				maskedFields: 66,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masker{
				mask:         tt.fields.mask,
				maskedFields: tt.fields.maskedFields,
			}
			m.SetMask(tt.args.s)
			assert.Equal(t, tt.wantFields.mask, m.mask)
			assert.Equal(t, tt.wantFields.maskedFields, m.maskedFields)
		})
	}
}

func TestMasker_ApplyMask(t *testing.T) {
	type fields struct {
		mask         int64
		maskedFields int64
	}
	type args struct {
		i int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "success",
			fields: fields{
				mask:         64,
				maskedFields: 66,
			},
			args: args{
				i: 11,
			},
			want: 73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Masker{
				mask:         tt.fields.mask,
				maskedFields: tt.fields.maskedFields,
			}
			if got := m.ApplyMask(tt.args.i); got != tt.want {
				t.Errorf("ApplyMask() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMasker_CountRelativeAddresses(t *testing.T) {
	type fields struct {
		mask              int64
		maskedFields      int64
		relativeAddresses []int64
	}
	type args struct {
		addressBinary string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		//{
		//	name: "success",
		//	args: args{
		//		addressBinary: "11010",
		//	},
		//	want: []int64{5, 4, 1, 0},
		//},
		{
			name: "success",
			args: args{
				addressBinary: "010111111111111111111111101110110111",
			},
			want: []int64{
				34359738368 + 8589934592 + 1024 + 64 + 8,

				34359738368 + 8589934592 + 1024 + 64,
				34359738368 + 8589934592 + 1024 + 8,
				34359738368 + 8589934592 + 64 + 8,
				34359738368 + 1024 + 64 + 8,
				8589934592 + 1024 + 64 + 8,

				34359738368 + 8589934592 + 1024,
				34359738368 + 8589934592 + 64,
				34359738368 + 1024 + 64,
				8589934592 + 1024 + 64,
				34359738368 + 8589934592 + 8,
				34359738368 + 1024 + 8,
				8589934592 + 1024 + 8,
				34359738368 + 64 + 8,
				8589934592 + 64 + 8,
				1024 + 64 + 8,

				34359738368 + 8589934592,
				34359738368 + 1024,
				8589934592 + 1024,
				34359738368 + 64,
				8589934592 + 64,
				1024 + 64,
				34359738368 + 8,
				8589934592 + 8,
				1024 + 8,
				64 + 8,

				34359738368,
				8589934592,
				1024,
				64,
				8,
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countRelativeAddresses(tt.args.addressBinary)
			sort.Slice(got, func(i, j int) bool { return got[i] < got[j] })
			sort.Slice(tt.want, func(i, j int) bool { return tt.want[i] < tt.want[j] })
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMasker(t *testing.T) {
	type fields struct {
		mask              int64
		maskedFields      int64
		relativeAddresses []int64
	}
	type args struct {
		memMap  map[int64]int64
		address int64
		value   int64
	}
	tests := []struct {
		name string
		mask string
		args args
		want map[int64]int64
	}{
		{
			name: "success",
			mask: "1X0X",
			args: args{
				memMap:  make(map[int64]int64),
				address: 5,
				value:   12,
			},
			want: map[int64]int64{
				8:  12,
				9:  12,
				12: 12,
				13: 12,
			},
		},
		{
			name: "success",
			mask: "000000000000000000000000000000X1001X",
			args: args{
				memMap:  make(map[int64]int64),
				address: 42,
				value:   100,
			},
			want: map[int64]int64{
				26: 100,
				27: 100,
				58: 100,
				59: 100,
			},
		},
		{
			name: "success 2nd step",
			mask: "00000000000000000000000000000000X0XX",
			args: args{
				memMap: map[int64]int64{
					26: 100,
					27: 100,
					58: 100,
					59: 100,
				},
				address: 26,
				value:   1,
			},
			want: map[int64]int64{
				16: 1,
				17: 1,
				18: 1,
				19: 1,
				24: 1,
				25: 1,
				26: 1,
				27: 1,
				58: 100,
				59: 100,
			},
		},
		{
			name: "success real data",
			mask: "X0X0111111000000100101100X000X10X001",
			args: args{
				memMap:  map[int64]int64{},
				address: 56856,
				value:   474071,
			},
			want: map[int64]int64{
				56856 + 34359738368 + 8589934592 + 1024 + 64 + 8: 474071,
				56856 + 34359738368 + 8589934592 + 1024 + 64:     474071,
				56856 + 34359738368 + 8589934592 + 1024 + 8:      474071,
				56856 + 34359738368 + 8589934592 + 64 + 8:        474071,
				56856 + 34359738368 + 1024 + 64 + 8:              474071,
				56856 + 8589934592 + 1024 + 64 + 8:               474071,
				56856 + 34359738368 + 8589934592 + 1024:          474071,
				56856 + 34359738368 + 8589934592 + 64:            474071,
				56856 + 34359738368 + 1024 + 64:                  474071,
				56856 + 8589934592 + 1024 + 64:                   474071,
				56856 + 34359738368 + 8589934592 + 8:             474071,
				56856 + 34359738368 + 1024 + 8:                   474071,
				56856 + 8589934592 + 1024 + 8:                    474071,
				56856 + 34359738368 + 64 + 8:                     474071,
				56856 + 8589934592 + 64 + 8:                      474071,
				56856 + 1024 + 64 + 8:                            474071,
				56856 + 34359738368 + 8589934592:                 474071,
				56856 + 34359738368 + 1024:                       474071,
				56856 + 8589934592 + 1024:                        474071,
				56856 + 34359738368 + 64:                         474071,
				56856 + 8589934592 + 64:                          474071,
				56856 + 1024 + 64:                                474071,
				56856 + 34359738368 + 8:                          474071,
				56856 + 8589934592 + 8:                           474071,
				56856 + 1024 + 8:                                 474071,
				56856 + 64 + 8:                                   474071,
				56856 + 34359738368:                              474071,
				56856 + 8589934592:                               474071,
				56856 + 1024:                                     474071,
				56856 + 64:                                       474071,
				56856 + 8:                                        474071,
				56856 + 0:                                        474071,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Masker{}
			m.SetMask(tt.mask)

			m.AddAddresses(tt.args.memMap, tt.args.address, tt.args.value)
			if !reflect.DeepEqual(tt.want, tt.args.memMap) {
				t.Fatal("not equal")
			}
		})
	}
}
