package main

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 2, 9}, []int{1, 3, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{0}, []int{1}},
		{[]int{9}, []int{1, 0}},
		{[]int{8, 9, 9}, []int{9, 0, 0}},
	}
	for _, c := range cases {
		in := append([]int{}, c.in...)
		got := PlusOne(in)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("PlusOne(%v) = %v, want %v", c.in, got, c.want)
		}
		if !reflect.DeepEqual(in, c.in) {
			t.Errorf("input mutated: %v", in)
		}
	}
}
