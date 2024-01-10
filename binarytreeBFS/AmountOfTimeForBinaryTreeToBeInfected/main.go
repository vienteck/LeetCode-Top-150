package main

import (
	"fmt"
)

// TreeNode represents a node in a binary tree.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example usage
	root := initializeTree()
	start := 3
	time := amountOfTime(root, start)
	fmt.Println("Time:", time)
}

// Function to initialize a binary tree for testing
func initializeTree() *TreeNode {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:  3,
				Left: nil,
				Right: &TreeNode{
					Val:  4,
					Left: nil,
					Right: &TreeNode{
						Val:   5,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}

	root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   10,
				Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
				Right: &TreeNode{Val: 9, Left: nil, Right: nil},
			},
		},
	}
	return root
}

// Function to calculate the amount of time for the entire tree to be infected
func amountOfTime(root *TreeNode, start int) int {
	if root == nil {
		return 0
	}

	time := 0
	visited := map[*TreeNode]struct{}{} // Keep track of visited nodes
	queue := []*TreeNode{root}
	found := false

	for len(queue) > 0 {
		n := len(queue)
		for i := 0; i < n; i++ {
			current := queue[0]
			queue = queue[1:]

			// Mark the node as visited
			visited[current] = struct{}{}

			// Check if the current node is the starting node
			if !found && current.Val == start {
				found = true
			}

			// Check and enqueue adjacent uninfected nodes
			if current.Left != nil && current.Left != current {
				if _, exists := visited[current.Left]; !exists {
					queue = append(queue, current.Left)
					visited[current.Left] = struct{}{}
				}
			}

			if current.Right != nil && current.Right != current {
				if _, exists := visited[current.Right]; !exists {
					queue = append(queue, current.Right)
					visited[current.Right] = struct{}{}
				}
			}
		}

		if !found {
			time++
		}
	}

	return time
}
