package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example:
	//        1
	//       / \
	//      2   3

	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	result := sumNumbers(root)

	// Print the result
	println(result) // Output: 25 (12 + 13)
}

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = currentSum*10 + node.Val

	// Check if it's a leaf node
	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	// Recursively calculate the sum for left and right subtrees
	return sumNumbersHelper(node.Left, currentSum) + sumNumbersHelper(node.Right, currentSum)
}
