package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	val  interface{}
	link *Node
}

type List struct {
	head *Node
}

func (list *List) push(value any) {
	node := Node{}
	node.val = value
	node.link = nil
	if list.head == nil {
		list.head = &node
	} else {
		temp := list.head
		for {
			if temp.link != nil {
				temp = temp.link
			} else {
				temp.link = &node
				break
			}
		}
	}
	fmt.Print("\n")
}

func (list *List) display() {
	temp := list.head
	if temp != nil {
		for {
			fmt.Print(temp.val, "->")
			if temp.link != nil {
				temp = temp.link
			} else {
				fmt.Println("\n")
				break
			}
		}
	} else {
		fmt.Println("!! No nodes to print.\n")
	}
}

func (list *List) pull() {
	temp := list.head
	if temp != nil {
		q := temp
		for {
			if temp.link != nil {
				q = temp
				temp = temp.link
			} else {
				fmt.Print("Removing the value", q.link.val)
				q.link = nil
				break
			}
		}
	} else {
		fmt.Println("!! No elements in the list to pull.\n")
	}
}

func (list *List) pop() *List {
	// func pop(list *List) *List {
	if list.head == nil {
		fmt.Println("!! Empty list to pop.\n")
	} else {
		fmt.Println("Popping first element:", list.head.val, "\n")
		list.head = list.head.link
	}
	return list
}

func main() {
	l := &List{}
	for {
		fmt.Println("OPTIONS\n1. push\n2. pop\n3. pull\n4. display")
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Enter choice of operation: ")
		scanner.Scan()
		choice := scanner.Text()
		if choice == "1" {
			fmt.Print("Enter the number:")
			scanner.Scan()
			l.push(scanner.Text())
		} else if choice == "2" {
			l = l.pop()
			// l = pop(l)
		} else if choice == "3" {
			l.pull()
		} else if choice == "4" {
			l.display()
		} else {
			fmt.Println("!! Invalid choice, exiting...")
			break
		}
	}

}
