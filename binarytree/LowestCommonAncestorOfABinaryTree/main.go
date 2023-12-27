package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// If the current node is either p or q, it is the LCA
	if root == p || root == q {
		return root
	}

	// Recursively search in the left and right subtrees
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// If both p and q are found in different subtrees, the current node is the LCA
	if left != nil && right != nil {
		return root
	}

	// If one of p or q is found, return that node as a potential LCA
	if left != nil {
		return left
	} else {
		return right
	}
}

func main() {
	// Example:
	//        3
	//       / \
	//      5   1
	//     / \ / \
	//    6  2 0  8
	//      / \
	//     7   4

	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	p := &TreeNode{Val: 5}
	q := &TreeNode{Val: 1}

	result := lowestCommonAncestor(root, p, q)

	// Print the result
	println(result.Val) // Output: 3
}
