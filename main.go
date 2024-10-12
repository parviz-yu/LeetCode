package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	q := l.PushBack(5)

	fmt.Println(q)

	for l.Len() > 0 {
		node := l.Remove(l.Back()).(int)
		fmt.Println(node)
	}
}
