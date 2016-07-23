package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	for i := 1; i <= num; i++ {
		var st string

		if i%3 == 0 {
			st += "Fizz"
		}
		if i%5 == 0 {
			st += "Buzz"
		}

		if st == "" {
			fmt.Println(i)
		} else {
			fmt.Println(st)
		}

	}
}
