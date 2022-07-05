package main

import "fmt"

type Node struct {
	val   string
	left  *Node
	right *Node
}

func (n *Node) preorder() {
	if n != nil {
		fmt.Println(n.val)
		n.left.preorder()
		n.right.preorder()
	}
}

func (n *Node) inorder() {
	if n != nil {
		n.left.inorder()
		fmt.Println(n.val)
		n.right.inorder()
	}
}

func (n *Node) postorder() {
	if n != nil {

		n.left.preorder()
		n.right.preorder()
		fmt.Println(n.val)
	}
}

func main() {
	plus := Node{"+", nil, nil}
	a := Node{"a", nil, nil}
	minus := Node{"-", nil, nil}
	b := Node{"b", nil, nil}
	c := Node{"c", nil, nil}

	plus.left = &a
	plus.right = &minus
	minus.left = &b
	minus.right = &c

	plus.postorder()
	fmt.Println()
	plus.inorder()
	fmt.Println()
	plus.preorder()

}
