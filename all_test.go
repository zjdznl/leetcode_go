package leetcode

import (
	"testing"
	"zhengjiadi.com/gorepos/tools"
)

func TestSlice(t *testing.T) {
	s := make([]int, 0, 100)
	s1 := make([]int, 100)
	println(len(s), " ", cap(s))
	println(len(s1), " ", cap(s1))
	s = append(s, 1)
	println(len(s), " ", cap(s), " ", tools.Json(s))
	s = s[:0]
	println(len(s), " ", cap(s), " ", tools.Json(s))
	s = append(s, 1)
	println(len(s), " ", cap(s), " ", tools.Json(s))
}

func TestA(t *testing.T) {
	s := make([]int, 0, 100)
	s = append(s, 1, 2, 3)
	println(s[:1])
	println(s[1:])
	println(s[3:])

}
