// BST
// binary search tree
package main

import (
	"fmt"
)

type BSTNode struct {
	Key         uint32
	Value       []interface{}
	Left, Right *BSTNode
}

func NewBSTNode(key uint32, value interface{}) *BSTNode {
	n := &BSTNode{
		Key:   key,
		Value: make([]interface{}, 1),
		Left:  nil,
		Right: nil,
	}
	n.Value[0] = value
	return n
}

type BST struct {
	Root *BSTNode
}

func Tree() *BST {
	return &BST{
		Root: nil,
	}
}

/*
	return the exist node or parent node.
*/
func (n *BSTNode) Find(parent *BSTNode, key uint32) (p, cur *BSTNode) {
	if key < n.Key {
		if n.Left == nil {
			return parent, n
		} else {
			return n.Left.Find(n, key)
		}
	} else if key > n.Key {
		if n.Right == nil {
			return parent, n
		} else {
			return n.Right.Find(n, key)
		}
	} else {
		return parent, n
	}
}

func (t *BST) Find(key uint32) (p, cur *BSTNode) {
	return t.Root.Find(t.Root, key)
}

func (t *BST) Add(key uint32, value interface{}) {
	if t.Root == nil {
		t.Root = NewBSTNode(key, value)
		return
	}
	_, n := t.Find(key)
	if key < n.Key {
		n.Left = NewBSTNode(key, value)
	} else if key > n.Key {
		n.Right = NewBSTNode(key, value)
	} else {
		n.Value = append(n.Value, value)
	}
}

func (n *BSTNode) LeftMax() (p, cur *BSTNode) {
	nLeft := n.Left
	p = nLeft
	for nLeft.Right != nil {
		p = nLeft
		nLeft = nLeft.Right
	}
	cur = nLeft
	return
}

func (t *BST) Del(key uint32) {
	if t.Root == nil {
		return
	}
	np, n := t.Find(key)
	if n.Key == key {
		if n.Left != nil && n.Right != nil {
			p, cur := n.LeftMax()
			fmt.Println(p, cur)
			n.Key = cur.Key
			n.Value = cur.Value
			if p.Key == cur.Key {
				n.Left = nil
			} else {
				p.Right = cur.Left
			}
		} else {
			if n.Key < np.Key {
				if n.Left == nil {
					np.Left = n.Right
				} else {
					np.Left = n.Left
				}
			} else if n.Key > np.Key {
				if n.Left == nil {
					np.Right = n.Right
				} else {
					np.Right = n.Left
				}
			} else {
				if n.Left == nil {
					t.Root = n.Right
				} else {
					t.Root = n.Left
				}
			}
		}
	}
}

func (t *BST) PrintN() {
	Print(t.Root)
}

func Print(n *BSTNode) {
	if n == nil {
		return
	}
	Print(n.Left)
	fmt.Printf("<%d,%v>", n.Key, n.Value)
	Print(n.Right)
}

func main() {
	t := Tree()
	t.Add(50, 50)
	t.Add(10, 10)
	t.Add(90, 90)
	t.Add(20, 20)
	t.Add(80, 80)
	t.Add(30, 30)
	t.Add(70, 70)
	t.Add(40, 40)
	t.Add(60, 60)
	t.PrintN()
	fmt.Println()
	t.Del(50)
	t.PrintN()
	fmt.Println()
	t.Del(10)
	t.PrintN()
	fmt.Println()
	t.Del(20)
	t.PrintN()
	fmt.Println()
	t.Del(30)
	t.PrintN()
	fmt.Println()
	t.Del(40)
	t.PrintN()
	fmt.Println()
	t.Del(50)
	t.PrintN()
	fmt.Println()
	t.Del(60)
	t.PrintN()
	fmt.Println()
	t.Del(70)
	t.PrintN()
	fmt.Println()
	t.Del(80)
	t.PrintN()
	fmt.Println()
	t.Del(90)
	t.PrintN()
	fmt.Println("::")
}
