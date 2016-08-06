/*
  Zeldaは、昨日PCにパスワードを設定した。

  今日、PCを起動したが、昨日入力したパスワードでは入れない事に気づいた。
  よくよく考えてみると、Caps Lock機能がオンのまま気づかずに入力してしまったようだ。
  *として隠されて表示されるため、入力中には気づかなかったらしい。

  パスワードは、大文字・小文字の半角アルファベット52種類のみ使用する。
  Caps Lockは入力するアルファベットが、小文字の入力なら大文字に、大文字の入力なら小文字として入力される機能である。

  昨日入力したはずのパスワードの文字列が与えられるので
  誤って設定された「現在の」パスワードを求めてください。
*/

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	var question string
	fmt.Scan(&question)

	upperCase := regexp.MustCompile(`[A-Z]`)
	lowerCase := regexp.MustCompile(`[a-z]`)

	answer := make([]byte, 0)
	for _, _s := range question {
		s := string(_s)
		if upperCase.MatchString(s) {
			answer = append(answer, strings.ToLower(s)...)
		}
		if lowerCase.MatchString(s) {
			answer = append(answer, strings.ToUpper(s)...)
		}
	}
	fmt.Println(string(answer))
}
