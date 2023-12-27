package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64
	_ = maxPathSumHelper(root, &maxSum)
	return maxSum
}

func maxPathSumHelper(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	// Calculate the maximum path sum for the left and right subtrees
	leftSum := max(0, maxPathSumHelper(node.Left, maxSum))
	rightSum := max(0, maxPathSumHelper(node.Right, maxSum))

	// Update the global maxSum with the current path sum
	*maxSum = max(*maxSum, leftSum+rightSum+node.Val)

	// Return the maximum path sum starting from the current node
	return max(leftSum, rightSum) + node.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example:
	//        10
	//       / \
	//      2   10
	//     / \    \
	//    20  1   -25
	//           / \
	//          3   4

	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 20,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: -25,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}

	result := maxPathSum(root)

	// Print the result
	fmt.Println(result) // Output: 42 (20 -> 2 -> 10 -> 10)
}
