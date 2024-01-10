package main

func main() {
	head := &TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	postorderTraversal(head)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	answer := []int{}

	if root.Left != nil {
		answer = append(answer, postorderTraversal(root.Left)...)
	}

	if root.Right != nil {
		answer = append(answer, postorderTraversal(root.Right)...)
	}

	answer = append(answer, root.Val)

	return answer
}
