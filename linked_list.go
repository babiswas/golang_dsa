package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
}

func (node *Node) create_node(val int) *Node {
	node.value = val
	node.next = nil
	return node
}

func (l *List) add_front(val int) {
	node := Node{}
	if l.head == nil {
		l.head = node.create_node(val)
	} else {
		temp := l.head
		l.head = node.create_node(val)
		l.head.next = temp
	}
}

func (l List) display() {
	temp := l.head
	fmt.Println(temp.value)
	for temp.next != nil {
		temp = temp.next
		fmt.Println(temp.value)
	}
}

func (l *List) add_last(val int) {
	node := Node{}
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node.create_node(val)
}

func main() {
	arr := []int{5, 6, 7, 8, 9}
	fmt.Println("Constructing the link list:")
	l := List{}
	fmt.Println(l)
	for i := 0; i < len(arr); i++ {
		l.add_front(arr[i])
	}
	fmt.Println("Traversing the list:")
	fmt.Println("============================")
	l.display()
	l.add_last(50)
	l.add_last(60)
	l.add_last(70)
	l.add_last(80)
	l.add_last(90)
	fmt.Println("============================")
	l.display()
}
