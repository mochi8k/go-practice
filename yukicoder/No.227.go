/*
  5枚のカードが配られます。それぞれのカードには、1以上13以下のいずれかの整数が書かれています。カードに書かれている整数の組み合わせによって役が決まります。

  配られた5枚のカードが、以下のいずれの役に該当するかを調べてください。複数の役に該当する場合は、以下で先に記述した方の役に該当するものとします。

  FULL HOUSE
  ある数をちょうど3つと、別の数をちょうど2つ含む。
  THREE CARD
  ある数をちょうど3つ含む。
  TWO PAIR
  ある数をちょうど2つと、別の数をちょうど2つ含む。
  ONE PAIR
  ある数をちょうど2つ含む。
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// 1:13
	var a, b, c, d, e int
	// fmt.Scan(&a, &b, &c, &d, &e)
	a, b, c, d, e = 4, 3, 3, 4, 4
	cards := []int{a, b, c, d, e}
	sort.Ints(cards)

	if isFullHouse(cards) {
		fmt.Println("FULL HOUSE")
	// } else if isThreeCard(cards) {
	// 	fmt.Println("THREE CARD")
	// } else if isTwoPair(cards) {
	// 	fmt.Println("TWO PAIR")
	// } else if isOnePair(cards) {
	// 	fmt.Println("ONE PAIR")
	} else {
		fmt.Print("NO HAND")
	}

}

func isConstain(numbers []int, number int) bool {
	for _, n := range numbers {
		if n == number {
			return true
		}
	}
	return false
}

func removeDupicates(numbers []int) []int {
	_numbers := make([]int, 0, len(numbers))
	for _, n := range numbers {
		if !isConstain(_numbers, n) {
			_numbers = append(_numbers, n)
		}
	}
	fmt.Println(_numbers)
	fmt.Println(numbers)
	return _numbers
}

func isFullHouse(cards []int) bool {
	_cards := removeDupicates(cards)
  fmt.Println(_cards)
  return true
}
// func isThreeCard(cards []int) bool {
//
// }
// func isTwoPair(cards []int) bool {
//
// }
// func isOnePair(cards []int) bool {
//
// }
