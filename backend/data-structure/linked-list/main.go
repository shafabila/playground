package main

import "fmt"

//inisialisasi node
type Node struct {
	prev *Node
	next *Node
	key interface{}
}
//inisialilasi


//fungsi u add node pada linked list
func (L *List) Insert(Key interface{}) {
	list := &Node{
		next: L.head,
		key: key,
	}
	if L.head != nil {
	L.head.prev = list
	}
	L.head = list

	l := L.head 
	for l.next != nil {
		l = l.nest
	}
	L.tail = l

	//fungsi u show linked list
	func (l *List) Display() {
		list := l.head
		for list != nil {
			fmt.Printf("%+v ->", list.key)
			list = list.next
		}
		fmt.Println()
	}

	func main() {
		link := List{}
		link.Insert(5)
		link.Insert(9)
		link.Insert(18)

		fmt.Printf("Head: %v\n", link.head.key) //result: 18
		fmt.Printf("Tail: %v\n", link.tail.key) //result: 5
		link.Display() //result: 18 -> 9 -> 5
	}
}