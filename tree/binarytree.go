package tree

import (
  "fmt"
)

type tNode struct {
  Value int
  Left *tNode
  Right *tNode
}

type BTree struct {
    root *tNode
    count int
}

func (tree *BTree) Add(value int) {
    if (tree.root == nil) {
      tree.root = newNode(value)
      return
    }

    temp := tree.root
    for {
      if value <= tree.root.Value {
        if temp.Left != nil {
          temp = temp.Left
        } else {
          temp.Left = newNode(value)
          break
        }
      } else {
        if temp.Right != nil {
          temp = temp.Right
        } else {
          temp.Right = newNode(value)
          break
        }
      }
    }
}

func (tree *BTree) InOrder() {
  inOrder(tree.root)
}

func newNode(value int) (node *tNode) {
  node = &tNode{value, nil, nil}
  return
}

func inOrder(root *tNode) {
  if (root == nil) {
    return
  }

  inOrder(root.Left)
  fmt.Println(root.Value)
  inOrder(root.Right)
}
