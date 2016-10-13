// Code for https://tour.golang.org/concurrency/8

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	ch <- t.Value
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	k := -1
	s := make([]int, 10)

	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i := 0; i < 10; i++ {
		s[i] = <-c1
	}
	for i := 0; i < 10; i++ {
		k = <-c2
		for j := 0; j < 10; j++ {
			if k == s[j] {
				break
			} else if j == 9 {
				return false
			}
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := tree.New(2)

	c := make(chan int)

	go Walk(t3, c)

	fmt.Println("Dump all of nodes")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	if Same(t1, t2) {
		fmt.Println("The trees are same")
	} else {
		fmt.Println("The trees are NOT same")
	}

	if Same(t1, t3) {
		fmt.Println("The trees are same")
	} else {
		fmt.Println("The trees are NOT same")
	}
}
