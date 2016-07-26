/*
  A+Bを計算します。

  しかし、'+'が嫌いなので'+'を使ってはいけません。
  ソースコード上に'+'の文字があると不正解になります。

  ※テンプレートで'+'を使っている方は気をつけてください。
  (特に入力やC++のbits/stdc++.h)
*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Print(-(-a - b))
}
