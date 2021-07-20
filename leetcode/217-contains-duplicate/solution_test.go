package leetcode

import (
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test1",
			args{
				[]int{1, 2, 3, 1},
			},
			true,
		},
		{
			"test2",
			args{
				[]int{1, 2, 3, 4},
			},
			false,
		},
		{
			"test3",
			args{
				[]int{},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicate2() = %v, want %v", got, tt.want)
			}
		})
	}
}
