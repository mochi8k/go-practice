package main

import (
	"fmt"
)

func from(n int) []int {
	slice := make([]int, 0)
	for i := range make([]int, n) {
		slice = append(slice, i+1)
	}
	return slice
}

func reduce(numbers []int, callback func(prev int, current int) int, dn int) int {
	currentNum := dn
	for _, v := range numbers {
		currentNum = callback(currentNum, v)
	}
	return currentNum
}

func main() {

	numberOfLines := 100
	multiple := 10

	reduce(from(numberOfLines), func(prev int, current int) int {
		for _, v := range from(current) {
			if (prev+v)%multiple == 0 {
				fmt.Print("★")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
		return prev + current
	}, 0)

}
