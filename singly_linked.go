package main

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (root *LinkedList) initialise(_val any) {
	node := Node{}
	node.val = _val
	node.next = nil
	if root.head == nil {
		root.head = &node
		return
	}

	temp := root.head
	for {
		if temp.next == nil {
			temp.next = &node
			return
		}
		temp = temp.next
	}
}

func TraverseHead(list LinkedList) {
	fmt.Print("\nprinting list by passing argument")
	_list := list.head
	for _list != nil {
		fmt.Print("\n", _list.val)
		_list = _list.next
	}
}

func (list *LinkedList) TraverseHead_by_interface() {
	fmt.Print("\nprinting list by interface")
	_list := list.head
	for _list != nil {
		fmt.Print("\n", _list.val)
		_list = _list.next
	}
}

func main() {
	_head := LinkedList{}
	_head.initialise(1)
	_head.initialise(2)
	_head.initialise(3)
	// calling TraverseHead by interface
	_head.TraverseHead_by_interface()
	// if we are to pass Linkedlist as param then below comment works
	TraverseHead(_head)
	fmt.Print("\n")

}
