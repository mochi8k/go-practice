/*
	1からNまでの整数iについて順に、
	iが3の倍数であれば Fizz 、5の倍数であれば Buzz 、
	それ以外であれば i を出力してください。
	ただし、iが3の倍数、かつ 5の倍数であれば FizzBuzz を出力してください。
*/

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
