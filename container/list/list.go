package main

import (
	"container/list"
	"fmt"
)

//链表的遍历
func print(l *list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func main() {
	l := list.New()

	l.PushBack(1)
	l.PushBack(2) //尾插
	print(l)
	fmt.Println("=========")
	l.PushFront(1) //头插
	print(l)
	fmt.Println("=========")
	for e := l.Front(); e != nil; e = e.Next() {
		if e.Value == 1 {
			l.InsertAfter(1.1, e)
		}
		if e.Value == 2 {
			l.InsertAfter(2.2, e)
		}
	}
	print(l)
	fmt.Println("=========")
	l.MoveToFront(l.Back())
	print(l)
	fmt.Println("=========")
	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
