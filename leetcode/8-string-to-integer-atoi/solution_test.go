package leetcode

import (
    "testing"
)

func Test_myAtoi(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            "test1",
            args{
                "42",
            },
            42,
        },
        {
            "test2",
            args{
                "   -42",
            },
            -42,
        },
        {
            "test3",
            args{
                "4193 with words",
            },
            4193,
        },
        {
            "test4",
            args{
                "words and 987",
            },
            0,
        },
        {
            "test5",
            args{
                "-91283472332",
            },
            -2147483648,
        },
        {
            "test6",
            args{
                "-2147483648",
            },
            -2147483648,
        },
        {
            "test7",
            args{
                "-2147483649",
            },
            -2147483648,
        },
        {
            "test8",
            args{
                "-2147483647",
            },
            -2147483647,
        },
        {
            "test9",
            args{
                "2147483647",
            },
            2147483647,
        },
        {
            "test10",
            args{
                "2147483648",
            },
            2147483647,
        },
        {
            "test11",
            args{
                "2147483646",
            },
            2147483646,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := myAtoi(tt.args.s); got != tt.want {
                t.Errorf("myAtoi() = %v, want %v", got, tt.want)
            }
        })
    }
}
