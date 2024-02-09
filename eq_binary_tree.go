package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	_walk(t, ch)
	close(ch)
}

func _walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t.Left, ch)
		ch <- t.Value
		_walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		n1, ok1 := <-ch1
		n2, ok2 := <-ch2

		if ok1 != ok2 || n1 != n2 {
			return false
		}

		if !ok1 {
			return true
		}
	}
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(2), tree.New(1)))
}
