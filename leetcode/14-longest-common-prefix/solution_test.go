package leetcode

import (
    "testing"
)

func Test_longestCommonPrefix(t *testing.T) {
    type args struct {
        strs []string
    }
    tests := []struct {
        name string
        args args
        want string
    }{
        {
            "test1",
            args{
                []string{
                    "flow",
                    "fl",
                },
            },
            "fl",
        },
        {
            "test2",
            args{
                []string{
                    "flow",
                },
            },
            "flow",
        },
        {
            "test3",
            args{
                []string{
                    "flow",
                    "flowz",
                },
            },
            "flow",
        },
        {
            "test4",
            args{
                []string{
                    "flow",
                    "flzw",
                },
            },
            "fl",
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := longestCommonPrefix(tt.args.strs); got != tt.want {
                t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
            }
        })
    }
}
