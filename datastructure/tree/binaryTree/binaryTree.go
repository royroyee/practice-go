package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {

	result := []int{}

	var inorder func(*TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return result

}

func postorderTraversal(root *TreeNode) []int {

	result := []int{}

	var postorder func(*TreeNode)

	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)
		result = append(result, root.Val)
	}

	postorder(root)

	return result

}
