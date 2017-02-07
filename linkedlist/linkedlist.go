package linkedlist

import (
	"fmt"
)
//Node represent node in a Linked List
type lNode struct {
	value int
	next  *lNode
}

type LinkedList struct {
	head *lNode
	count int
}

func (l *LinkedList) Add(value int) {
	if (l.head == nil) {
		l.head = l.newNode(value)
		l.count += 1
		return
	}

	//if head is not nil
	temp := l.head
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}
	temp.next = l.newNode(value)
	l.count += 1
}

func (l *LinkedList) newNode(value int) (node *lNode) {
	node = &lNode{value, nil}
	return
}

func (l *LinkedList) Count() int {
	return l.count
}

func (l *LinkedList) get(i int) (value int, err error) {
	temp := l.head
	if temp == nil {
		err = fmt.Errorf("Out of bound")
		return
	}

	for {
		if i == 0 || temp.next == nil {
			break
		}
		temp = temp.next
		i -= 1
	}

	if i > 0 {
		err = fmt.Errorf("Out of bound")
		return
	}

	value = temp.value
	return
}
