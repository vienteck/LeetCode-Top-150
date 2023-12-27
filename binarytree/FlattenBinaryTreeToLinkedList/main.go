package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var prev *TreeNode

	var flattenTree func(node *TreeNode)
	flattenTree = func(node *TreeNode) {
		if node == nil {
			return
		}

		// Post-order traversal
		flattenTree(node.Right)
		flattenTree(node.Left)

		// Update pointers
		node.Right = prev
		node.Left = nil
		prev = node
	}

	flattenTree(root)
}

func main() {
	// Example:
	//        1
	//       / \
	//      2   5
	//     / \   \
	//    3   4   6

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	flatten(root)

	// Print the flattened linked list
	printFlattenedList(root)
}

func printFlattenedList(root *TreeNode) {
	for root != nil {
		print(root.Val, " ")
		root = root.Right
	}
}
