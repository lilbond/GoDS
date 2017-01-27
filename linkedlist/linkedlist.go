package main

import "fmt"

//Node represent node in a Linked List
type Node struct {
	value int
	next  *Node
}

func main() {
	head := new(Node)

	i := 1
	temp := head
	for i <= 10 {
		next := Node{value: i}
		next.value = i
		temp.next = &next
		temp = temp.next
		i = i + 1
	}

	temp = head
	for {
		if temp == nil {
			break
		}
		fmt.Println(temp.value)
		temp = temp.next
	}
}
