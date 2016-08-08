package main

import (
	"fmt"
	"strings"
)

func from(n int) []int {
	slice := make([]int, 0)
	for i := range make([]int, n) {
		slice = append(slice, i+1)
	}
	return slice
}

func main() {
	for _, v := range from(5) {
		fmt.Println(strings.Repeat("*", v))
	}
}
