package linkedlist

import (
  "fmt"
  "testing"
)

func TestCount(t *testing.T) {
  linkedList := &LinkedList{}
  var i = 0
  for i < 10 {
    addNode(linkedList, i)
    i += 1
  }

  x, err := linkedList.get(11)

  fmt.Println(err)
  fmt.Println(x)

  if linkedList.Count() != 10 {
    t.Error("Count should be equal to:", 10)
  }
}

func addNode(l *LinkedList, value int) {
  l.Add(value)
}
