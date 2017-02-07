package tree

import (
  "testing"
)

func TestInOrder(t *testing.T) {
  tree := &BTree{}

  tree.Add(10)
  tree.Add(2)
  tree.Add(11)

  tree.InOrder()
}
