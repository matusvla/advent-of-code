package seatfiller_wrong

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestFillSeats(t *testing.T) {
	type args struct {
		inb        [][]byte
		iterations int
	}
	var tests = []struct {
		name  string
		args  args
		want  [][]byte
		want1 int
	}{
		//{
		//	name: "success",
		//	args: args{
		//		inb: [][]byte{
		//			[]byte("L.LL.LL.LL"),
		//			[]byte("LLLLLLL.LL"),
		//			[]byte("L.L.L..L.."),
		//			[]byte("LLLL.LL.LL"),
		//			[]byte("L.LL.LL.LL"),
		//			[]byte("L.LLLLL.LL"),
		//			[]byte("..L.L....."),
		//			[]byte("LLLLLLLLLL"),
		//			[]byte("L.LLLLLL.L"),
		//			[]byte("L.LLLLL.LL"),
		//		},
		//		iterations: 1,
		//	},
		//	want: [][]byte{
		//		[]byte("#.##.##.##"),
		//		[]byte("#######.##"),
		//		[]byte("#.#.#..#.."),
		//		[]byte("####.##.##"),
		//		[]byte("#.##.##.##"),
		//		[]byte("#.#####.##"),
		//		[]byte("..#.#....."),
		//		[]byte("##########"),
		//		[]byte("#.######.#"),
		//		[]byte("#.#####.##"),
		//	},
		//	want1: 71,
		//},
		{
			name: "success more iterations",
			args: args{
				inb: [][]byte{
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
				iterations: 20,
			},
			want: [][]byte{
				[]byte("#.#L.L#.##"),
				[]byte("#LLL#LL.L#"),
				[]byte("L.#.L..#.."),
				[]byte("#L##.##.L#"),
				[]byte("#.#L.LL.LL"),
				[]byte("#.#L#L#.##"),
				[]byte("..L.L....."),
				[]byte("#L#L##L#L#"),
				[]byte("#.LLLLLL.L"),
				[]byte("#.#L#L#.##"),
			},
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.inb
			got1 := 0
			for i := 0; i < tt.args.iterations; i++ {
				got, got1 = FillSeats3(got)
				for _, val := range got {
					fmt.Println(string(val))
				}
				fmt.Println("____________________________")
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FillSeats() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FillSeats() got1 = %v, want %v", got1, tt.want1)
			}
			var result int
			for _, val := range got {
				result += bytes.Count(val, []byte{'#'})
			}
			fmt.Println(result)
		})
	}
}
