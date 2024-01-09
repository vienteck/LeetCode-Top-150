package main

func main() {
	head := &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	preorderTraversal(head)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	answer := []int{}

	stack := []*TreeNode{root}

	for len(stack) > 0 {
		n := len(stack)
		current := stack[n-1]
		stack = stack[:n-1]
		answer = append(answer, current.Val)
		if current.Right != nil {
			stack = append(stack, current.Right)
		}
		if current.Left != nil {
			stack = append(stack, current.Left)
		}

	}

	return answer
}
