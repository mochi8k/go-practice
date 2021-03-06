/*
  Y市にはマフィアの構成員が多く住み、犯罪の温床となっている。
  そこで、Y市はマフィア判定器を導入し、構成員と判定された者を
  すべて逮捕することにした。
  マフィア判定器とは、対象者が
  一般市民か構成員か99％の精度で判定できる機器である。

  Y市の全人口1,000,000をすべてに、
  マフィア判定器で判定して「構成員」と判定された者を逮捕する。

  しかし、実はY市には10,000人あたりN人しか構成員がいないということがわかった。
  このとき、逮捕者が出た時、誤認逮捕である確率を求めよ。

  答えの誤差は、絶対誤差、相対誤差±0.01まで許容される。
*/

package main

import (
	"fmt"
)

const population = 1000000

func main() {
	var n float64
	fmt.Scan(&n)

	member := n * 100
	citizen := population - member
	arrestedMember := member * 0.99
	arrestedCitizen := citizen * 0.01
	allArrested := arrestedMember + arrestedCitizen

	fmt.Println(arrestedCitizen / allArrested * 100)
}
