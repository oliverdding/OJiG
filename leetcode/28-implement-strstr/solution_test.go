package leetcode

import (
    "reflect"
    "testing"
)

func Test_calPreLength(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name          string
        args          args
        wantPreLength []int
    }{
        {
            "test1",
            args{
                "abababca",
            },
            []int{0, 0, 1, 2, 3, 4, 0, 1},
        },
        {
            "test2",
            args{
                "abcabc",
            },
            []int{0, 0, 0, 1, 2, 3},
        },
        {
            "test3",
            args{
                "bba",
            },
            []int{0, 1, 0},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if gotPreLength := calPreLength(tt.args.s); !reflect.DeepEqual(gotPreLength, tt.wantPreLength) {
                t.Errorf("calPreLength() = %v, want %v", gotPreLength, tt.wantPreLength)
            }
        })
    }
}

func Test_strStr2(t *testing.T) {
    type args struct {
        haystack string
        needle   string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            "test1",
            args{
                "hello",
                "ll",
            },
            2,
        },
        {
            "test2",
            args{
                "aaaaa",
                "bba",
            },
            -1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := strStr2(tt.args.haystack, tt.args.needle); got != tt.want {
                t.Errorf("strStr2() = %v, want %v", got, tt.want)
            }
        })
    }
}
