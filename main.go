package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walker(t, ch)
	close(ch)
}

func Walker(t *tree.Tree, ch chan int) {
	// if the node is empty go back
	if t == nil {
		return
	}
	Walker(t.Left, ch)
	ch <- t.Value
	Walker(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

  //gorutines walks through the channel and adds the value
	go Walk(t1, ch1)
	go Walk(t2, ch2)

  // here we are checking each value and the ok value to both channels
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		fmt.Println("ops", v1, ok1, v2, ok2)
		if (ok1 && (v1 != v2)) || ok1 != ok2 {
			return false
		} else if !ok1 {
			break
		}
	}
	return true

}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(2)))
}
