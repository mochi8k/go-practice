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
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc.Split(bufio.ScanWords)
	a := nextInt()
	b := nextInt()
	c := nextInt()
	d := nextInt()
	e := nextInt()

	cards := []int{a, b, c, d, e}
	cardMap := map[int]int{}

	for _, v := range cards {
		cardMap[v] = (cardMap[v] + 1)
	}

	allocationNumbers := make([]int, 0)
	for _, value := range cardMap {
		allocationNumbers = append(allocationNumbers, value)
	}

	sort.Ints(allocationNumbers)

	if isFullHouse(allocationNumbers) {
		fmt.Println("FULL HOUSE")
	} else if isThreeCard(allocationNumbers) {
		fmt.Println("THREE CARD")
	} else if isTwoPair(allocationNumbers) {
		fmt.Println("TWO PAIR")
	} else if isOnePair(allocationNumbers) {
		fmt.Println("ONE PAIR")
	} else {
		fmt.Println("NO HAND")
	}

}

func isFullHouse(allocationNumbers []int) bool {
	return reflect.DeepEqual(allocationNumbers, []int{2, 3})
}

func isThreeCard(allocationNumbers []int) bool {
	return reflect.DeepEqual(allocationNumbers, []int{1, 1, 3})
}

func isTwoPair(allocationNumbers []int) bool {
	return reflect.DeepEqual(allocationNumbers, []int{1, 2, 2})
}

func isOnePair(allocationNumbers []int) bool {
	return reflect.DeepEqual(allocationNumbers, []int{1, 1, 1, 2})
}
