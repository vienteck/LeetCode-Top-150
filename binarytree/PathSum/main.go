package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	// Check if it's a leaf node and the target sum is equal to the node value
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}

	// Recursively check for the target sum in the left and right subtrees
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func main() {
	// Example:
	//        5
	//       / \
	//      4   8
	//     /   / \
	//    11  13  4
	//   /  \      \
	//  7    2      1

	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	targetSum := 22
	result := hasPathSum(root, targetSum)

	// Print the result
	println(result) // Output: true
}
