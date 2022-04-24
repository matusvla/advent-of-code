package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewPacket(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  Packet
		want1 int
	}{
		{
			name: "literal 1",
			args: args{
				s: "110100101111111000101000000000000000",
			},
			want: &LiteralPacket{
				version: 6,
				val:     2021,
			},
			want1: len("110100101111111000101"),
		},
		{
			name: "literal 2",
			args: args{
				s: "110100101111111000101000101000101011111010101010010101",
			},
			want: &LiteralPacket{
				version: 6,
				val:     2021,
			},
			want1: len("110100101111111000101"),
		},
		{
			name: "literal 3",
			args: args{
				s: "11010000001",
			},
			want: &LiteralPacket{
				version: 6,
				val:     1,
			},
			want1: len("11010000001"),
		},
		{
			name: "operator length",
			args: args{
				s: strings.ReplaceAll("001 110 0 000000000011011 110 100 01010 010 100 1000100100 00 00000", " ", ""),
			},
			want: &OperationPacket{
				version:      1,
				typeID:       6,
				lengthTypeID: 0,
				subpackets: []Packet{
					&LiteralPacket{
						version: 6,
						val:     10,
					},
					&LiteralPacket{
						version: 2,
						val:     20,
					},
				},
			},
			want1: len("00111000000000000110111101000101001010010001001000000000") - 7,
		},
		{
			name: "operator length - 3 literals",
			args: args{
				s: "0010100000000000100001110100000011101000000111010000001",
			},
			want: &OperationPacket{
				version:      1,
				typeID:       2,
				lengthTypeID: 0,
				subpackets: []Packet{
					&LiteralPacket{
						version: 6,
						val:     1,
					},
					&LiteralPacket{
						version: 6,
						val:     1,
					},
					&LiteralPacket{
						version: 6,
						val:     1,
					},
				},
			},
			want1: len("0010100000000000100001110100000011101000000111010000001"),
		},
		{
			name: "operator size",
			args: args{
				s: "11101110000000001101010000001100100000100011000001100000",
			},
			want: &OperationPacket{
				version:      7,
				typeID:       3,
				lengthTypeID: 1,
				subpackets: []Packet{
					&LiteralPacket{
						version: 2,
						val:     1,
					},
					&LiteralPacket{
						version: 4,
						val:     2,
					},
					&LiteralPacket{
						version: 1,
						val:     3,
					},
				},
			},
			want1: len("111011100000000011010100000011001000001000110000011"),
		},
		//{
		//	name: "complex 1",
		//	args: args{
		//		s: "01100010000000001000000000000000000101100001000101010110001011001000100000000010000100011000111000110100000000000000000000000000000",
		//	},
		//	want: &OperationPacket{
		//		version:      7,
		//		typeID:       3,
		//		lengthTypeID: 1,
		//		subpackets: []Packet{
		//			&OperationPacket{
		//				version: 2,
		//			},
		//			&OperationPacket{
		//				version: 4,
		//			},
		//		},
		//	},
		//	want1: len("011000100000000010000000000000000001011000010001010101100010110010001000000000100001000110001110001101"),
		//},
		{
			name: "complex 2",
			args: args{
				s: "0110101000000000010010100000000000100001110100000011101000000111010000001",
			},
			want: &OperationPacket{
				version:      3,
				typeID:       2,
				lengthTypeID: 1,
				subpackets: []Packet{
					&OperationPacket{
						version:      1,
						typeID:       2,
						lengthTypeID: 0,
						subpackets: []Packet{
							&LiteralPacket{
								version: 6,
								val:     1,
							},
							&LiteralPacket{
								version: 6,
								val:     1,
							},
							&LiteralPacket{
								version: 6,
								val:     1,
							},
						},
					},
				},
			},
			want1: len("0110101000000000010010100000000000100001110100000011101000000111010000001"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := NewPacket(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPacket() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("NewPacket() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
