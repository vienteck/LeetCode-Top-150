package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	rootIndex := indexOf(rootVal, inorder)

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])

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
	// Inorder: [9,3,15,20,7]
	// Postorder: [9,15,7,20,3]

	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)

	// Print the constructed tree (in-order traversal)
	inOrderTraversal(root)
}
