// binary search tree

package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

type BST struct {
	root *Node
}

func (bst *BST) Insert(v int) *BST {
	if bst.root == nil {
		bst.root = &Node{}
	} else {
		bst.root.Insert(v)
	}
	return bst
}

func (node *Node) Insert(v int) *Node {

	if v > node.value { // 삽입 하려는 v의 값이 node 값  보다 큰 경우
		if node.right == nil { // 오른쪽 자식이 비어있다면
			node.right = &Node{value: v} // 노드 생성
			return node.right
		} else {
			return node.right.Insert(v)
		}
	} else { // v가 node 값 보다 작다면
		if node.left == nil {
			node.left = &Node{value: v}
			return node.left
		} else {
			return node.left.Insert(v)
		}

	}
}

func (bst *BST) Search(key int) (bool, int) {
	if bst.root.value == key {
		return true, 0
	}
	return bst.root.Search(key, 1)
}

func (node *Node) Search(key int, cnt int) (bool, int) {
	if node.value == key { // 찾은 경우
		return true, cnt
	}

	if node.value > key { // key 값이 더 큰 경우
		if node.left != nil {
			return node.left.Search(key, cnt+1)
		}
		return false, cnt

	} else { // key 값이 더 작은 경우
		if node.right != nil {
			return node.right.Search(key, cnt+1)
		}
		return false, cnt
	}
}

func main() {

	bst := &BST{}

	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(8)
	bst.Insert(7)
	bst.Insert(6)
	bst.Insert(10)
	bst.Insert(9)

	if success, cnt := bst.Search(3); success {
		fmt.Println(success, cnt)
	} else {
		fmt.Println(success, cnt)
	}
}
