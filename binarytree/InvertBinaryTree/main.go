package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// Swap left and right subtrees
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

	return root
}

func main() {
	// Example tree:
	//          4
	//         / \
	//        2   7
	//       / \ / \
	//      1  3 6  9

	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	invertedTree := invertTree(root)

	// Print the inverted tree (in-order traversal)
	inOrderTraversal(invertedTree)
}

func inOrderTraversal(root *TreeNode) {
	if root != nil {
		inOrderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inOrderTraversal(root.Right)
	}
}
