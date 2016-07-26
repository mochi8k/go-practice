/*
  1歩でぴったりaセンチメートル歩ける。
  bセンチメートルの区間を歩くのに何歩歩く？
  なお、区間はオーバーしても良い。
*/

package main

import (
  "fmt"
)

func main() {
  var a, b int
  fmt.Scan(&a, &b)
  rem := b%a
  if rem == 0 {
    fmt.Println(b/a)
  } else {
    fmt.Println(b/a+1)
  }
}
