package main

import (
	"math"
)

func main() {
	depth := 10
	print(depth)
	root := generateBinaryTree(1, depth)
	print(maxDepth(root))
}

func generateBinaryTree(val, maxDepth int) *TreeNode {
	if maxDepth == 0 {
		return nil
	}

	node := &TreeNode{Val: val}
	maxChildVal := int(math.Pow(2, float64(maxDepth-1)))

	if val*2 <= maxChildVal {
		node.Left = generateBinaryTree(val*2, maxDepth-1)
	}
	if val*2+1 <= maxChildVal {
		node.Right = generateBinaryTree(val*2+1, maxDepth-1)
	}

	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxDepth := 0
	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			current := queue[0]
			queue = queue[1:]
			if current.Left != nil {
				queue = append(queue, current.Left)
			}
			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}

		maxDepth++
	}

	return maxDepth
}
