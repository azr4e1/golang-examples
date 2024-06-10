package equivalency

import (
	"golang.org/x/tour/tree"
)

// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values
func Same(t1, t2 *tree.Tree) bool {
	var ch1, ch2 = make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for <-ch1 == <-ch2 {

	}
	return false
}
