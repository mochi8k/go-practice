/*
	6面の普通のダイスのN個のサイコロがあります。
	全部振って出る目の合計の期待値を求めよ。
*/

package main

import (
	"fmt"
)

func main() {
	var numberOfDice int
	fmt.Scan(&numberOfDice)
	fmt.Print(float32(numberOfDice) * 3.5)
}
