package main

import (
	"fmt"
)

func main() {
	var numberOfDice int
	fmt.Scan(&numberOfDice)
	fmt.Print(float32(numberOfDice) * 3.5)
}
