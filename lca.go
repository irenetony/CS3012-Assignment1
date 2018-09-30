package lca

import "errors"

//Contruct a Node

type Node struct {
	Left  *Node
	Right *Node
	Key   int
	Value int
}

func EmptyBinaryTree() *Node {
	return &Node{}
}

func (n *Node) Insert(key int, data int) error {

	if n == nil {
		return errors.New("Cannot insert a value into a nil tree")
	}

	switch {
	//If the data is already in the tree, return.
	case key == n.Key:
		return nil
		//If the data value is less than the current node’s value, and if the left child node is nil, insert a new left child node. Else call Insert on the left subtree.
	case key < n.Key:
		if n.Left == nil {
			n.Left = &Node{Key: key, Value: data}
			return nil
		}
		return n.Left.Insert(key, data)
		//If the data value is greater than the current node’s value, do the same but for the right subtree.
	case key > n.Key:
		if n.Right == nil {
			n.Right = &Node{Key: key, Value: data}
			return nil
		}
		return n.Right.Insert(key, data)
	}
	return nil
}

func (n *Node) Get(x int) int {
	if n == nil {
		return -1
	}
	switch {
	//If the current node contains the value, return the node.
	case x == n.Key:
		return n.Value
		//If the data value is less than the current node’s value, call Find for the left child node,
	case x < n.Key:
		return n.Left.Get(x)
		//else call Find for the right child node.
	default:
		return n.Right.Get(x)
	}
}
