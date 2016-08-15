/*
  A君は当たりクジ付きのアイスバーが大好きである。
  アイスバーには「ハズレ」と「1本当たり」と「2本当たり」がある。
  ハズレは何ももらえないが、当たりだとその本数をタダでもらえる。

  お店ではアイスバーが箱詰めされて売られている。
  1つの箱にはN本のアイスバーが入っている。
  A君は箱の先頭から順にしかアイスバーを取り出すことはできない。
  買う場合も当たりと引き換えの場合もこの箱からアイスバーを取り出す。
  1つの箱の中のハズレと当たりクジの配置はどの箱も同じである。
  お店には1つのアイスバーの箱があり、売り切れると1つの新しい箱を補充する。
  いまお店には新しい手つかずのアイスバーの箱がある。
  A君は最初は「当たり」クジを持っておらず予算は無限にある。
  K本のアイスバーを食べるためにはA君は最低何本のアイスを買う必要があるか？

  Nは1箱に入っているアイスバーの数。1≤N≤501≤N≤50。NNは正の整数。
  KはA君が食べる予定のアイスバーの数。1≤K≤2000000000=2⋅1091≤K≤2000000000=2⋅109。KKは正の整数。
  Sの長さはNNの値と等しい。
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var numberOfIce, numberOfEat, numberOfEated int
	var s string
	fmt.Scan(&numberOfIce, &numberOfEat, &s)
	cases := strings.Split(s, "")

  for {
    if numberOfEat >= numberOfEated {
      break
    }

  	for _, v := range cases {
      n, _ := strconv.Atoi(v)
  		switch n {
      case 0:

      case 1:

      case 2:
  		}
  	}

  }

}
