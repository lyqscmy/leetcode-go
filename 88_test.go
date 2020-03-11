package leetcode

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3}, []int{1, 2, 2, 3, 5, 6}},
		{"", args{nums1: []int{4, 5, 6, 0, 0, 0}, m: 3, nums2: []int{1, 2, 3}, n: 3}, []int{1, 2, 3, 4, 5, 6}},
		{"", args{nums1: []int{0, 0, 0, 0, 0}, m: 0, nums2: []int{1, 2, 3, 4, 5}, n: 5}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("got %v, want %v", tt.args.nums1, tt.want)
			}
		})
	}
}
