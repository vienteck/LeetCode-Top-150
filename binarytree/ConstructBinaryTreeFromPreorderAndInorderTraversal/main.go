package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIndex := indexOf(rootVal, inorder)

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])

	return root
}

func indexOf(val int, arr []int) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}

func inOrderTraversal(root *TreeNode) {
	if root != nil {
		inOrderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inOrderTraversal(root.Right)
	}
}

func main() {
	// Example:
	// Preorder: [3,9,20,15,7]
	// Inorder: [9,3,15,20,7]

	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)

	// Print the constructed tree (in-order traversal)
	inOrderTraversal(root)
}
