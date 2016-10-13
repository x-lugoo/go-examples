// Code for https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	word := strings.Fields(s)

	for i := 0; i < len(word); i++ {
		m[word[i]]++
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
