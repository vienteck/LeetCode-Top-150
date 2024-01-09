package main

import (
	"fmt"
	"math"
)

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}

	leftHeight := getHeight(node.Left)
	rightHeight := getHeight(node.Right)

	return 1 + int(math.Max(float64(leftHeight), float64(rightHeight)))
}

func main() {
	// Example usage:
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	fmt.Println(isBalanced(root))
}
