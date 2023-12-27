package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	head := root
	for head != nil {
		dummyNode := new(Node)
		temp := dummyNode

		for head != nil {
			if head.Left != nil {
				temp.Next = head.Left
				temp = temp.Next
			}
			if head.Right != nil {
				temp.Next = head.Right
				temp = temp.Next
			}
			head = head.Next
		}
		head = dummyNode.Next
	}
	return root
}

func main() {
	// Example:
	//        1
	//       / \
	//      2   3
	//     / \   \
	//    4   5   7

	root := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Right: &Node{
				Val: 7,
			},
		},
	}

	connect(root)

	// Print the connected tree
}
