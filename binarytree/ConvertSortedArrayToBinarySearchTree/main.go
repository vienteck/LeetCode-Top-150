package main

import (
	"fmt"
)

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}

	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

func printTreeInOrder(root *TreeNode) {
	if root != nil {
		printTreeInOrder(root.Left)
		fmt.Print(root.Val, " ")
		printTreeInOrder(root.Right)
	}
}

func main() {
	// Example usage:
	sortedArray := []int{-10, -3, 0, 5, 9}
	resultTree := sortedArrayToBST(sortedArray)

	// Print the resulting tree in order to verify
	printTreeInOrder(resultTree)
}
