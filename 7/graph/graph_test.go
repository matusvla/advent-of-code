package graph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_extractNodeInfo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{
			name: "success",
			args: args{
				s: "5 dark blue bags",
			},
			want:  5,
			want1: "dark blue",
		},
		{
			name: "success no",
			args: args{
				s: "no dark blue bags",
			},
			want:  0,
			want1: "dark blue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := extractNodeInfo(tt.args.s)
			if got != tt.want {
				t.Errorf("extractNodeInfo() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("extractNodeInfo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDirected_AddNode(t *testing.T) {
	type args struct {
		nodeDescription string
	}
	tests := []struct {
		name string
		args args
		want *Directed
	}{
		{
			name: "success",
			args: args{
				nodeDescription: "light red bags contain 1 bright white bag, 2 muted yellow bags.",
			},
			want: &Directed{
				followers: map[string][]nodePath{
					"bright white": {
						{
							nodeID:  "light red",
							edgeVal: 1,
						},
					},
					"muted yellow": {
						{
							nodeID:  "light red",
							edgeVal: 2,
						},
					},
				},
				predecessors: map[string][]nodePath{
					"light red": {
						{
							nodeID:  "bright white",
							edgeVal: 1,
						},
						{
							nodeID:  "muted yellow",
							edgeVal: 2,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := New()
			d.AddNode(tt.args.nodeDescription)
			assert.EqualValues(t, tt.want, d)
		})
	}
}

func TestDirected_FindContainers(t *testing.T) {
	d := New()
	graphDesc := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}
	for _, val := range graphDesc {
		d.AddNode(val)
	}
	fmt.Println(d.FindContainers("shiny gold", nil, 0))
}

func TestDirected_FindContained(t *testing.T) {
	d := New()
	graphDesc := []string{
		"shiny gold bags contain 2 dark red bags.",
		"dark red bags contain 2 dark orange bags.",
		"dark orange bags contain 2 dark yellow bags.",
		"dark yellow bags contain 2 dark green bags.",
		"dark green bags contain 2 dark blue bags.",
		"dark blue bags contain 2 dark violet bags.",
		"dark violet bags contain no other bags.",
	}
	for _, val := range graphDesc {
		d.AddNode(val)
	}
	fmt.Println(d.CountBags("shiny gold"))
}
