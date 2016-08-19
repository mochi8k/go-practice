package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var tlds = []string{"com", "net"}

const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {

	if len(os.Args) > 1 {
		tlds = strings.Split(os.Args[1], ",")
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune
		for _, r := range text {
			if unicode.IsSpace(r) {
				r = '-'
			}

			if strings.ContainsRune(allowedChars, r) {
				newText = append(newText, r)
			}
		}
		fmt.Println(string(newText) + "." + tlds[rand.Intn(len(tlds))])
	}
}
