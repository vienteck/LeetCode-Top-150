package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

	invertedTree := isSymmetric(root)

	fmt.Println(invertedTree)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}

	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left == nil || right == nil {
			return false
		}

		if left.Val != right.Val {
			return false
		}

		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}

	return true
}
