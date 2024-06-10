package main

import (
	"equivalency"
	"fmt"
	"golang.org/x/tour/tree"
)

func main() {
	ch := make(chan int)
	newTree := tree.New(1)
	go equivalency.Walk(newTree, ch)
	k := 10
	for i := 0; i < k; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println(equivalency.Same(tree.New(1), tree.New(1)))
}
