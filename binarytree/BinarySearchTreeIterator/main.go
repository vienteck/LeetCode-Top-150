package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	return BSTIterator{
		stack: inOrderTraversal(root),
	}
}

func inOrderTraversal(node *TreeNode) []*TreeNode {
	stack := []*TreeNode{}
	current := node

	for current != nil {
		stack = append(stack, current)
		current = current.Left
	}

	return stack
}

func (it *BSTIterator) Next() int {
	top := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]

	// Perform in-order traversal on the right subtree of the popped node
	if top.Right != nil {
		it.stack = append(it.stack, inOrderTraversal(top.Right)...)
	}

	return top.Val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

func main() {
	// Example:
	//        7
	//       / \
	//      3   15
	//         / \
	//        9   20

	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}

	iterator := Constructor(root)

	// Output: [3, 7, 9, 15, 20]
	for iterator.HasNext() {
		println(iterator.Next())
	}
}
