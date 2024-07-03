package main

import (
	"fmt"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func (t *Tree) String() string {
	queue := []*Tree{t}
	output := ""
	level := 1
	currentCount := 0
	var sep string
	for len(queue) != 0 {
		sep = " "
		el := queue[0]
		currentCount++

		if currentCount == level {
			onlyNil := true
			for _, tree := range queue {
				if tree != nil {
					onlyNil = false
				}
			}
			if onlyNil {
				output += fmt.Sprintf("NULL%s", sep)
				break
			}

			level *= 2
			currentCount = 0
			sep = "\n"
		}
		if el == nil {
			queue = append(queue[1:], nil, nil)
			output += fmt.Sprintf("NULL%s", sep)
			continue
		}
		queue = append(queue[1:], el.Left, el.Right)

		output += fmt.Sprintf("%d%s", el.Val, sep)
	}

	return output
}

func InvertTree(t *Tree) *Tree {
	if t == nil {
		return nil
	}

	newTree := &Tree{
		Val:   t.Val,
		Left:  InvertTree(t.Right),
		Right: InvertTree(t.Left),
	}

	return newTree
}

func main() {
	example := &Tree{
		Val: 5,
		Right: &Tree{
			Val: 2,
			Right: &Tree{
				Val:   3,
				Right: nil,
				Left: &Tree{
					Val:   0,
					Right: nil,
					Left:  nil,
				},
			},
			Left: &Tree{
				Val: -1,
				Right: &Tree{
					Val:   8,
					Right: nil,
					Left:  nil,
				},
			},
		},
		Left: &Tree{
			Val:   0,
			Right: nil,
			Left:  nil,
		},
	}

	fmt.Println(example)
	fmt.Println("-----------")

	newTree := InvertTree(example)

	fmt.Println(newTree)
}
