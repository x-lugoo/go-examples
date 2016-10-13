// Code for https://tour.golang.org/moretypes/26

package main

import "fmt"

func fibonacci() func() int {
	o, p := -1, 1

	return func() int {
		o, p = p, o+p
		return p
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
