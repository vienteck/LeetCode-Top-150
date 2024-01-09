package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	level := 1
	for len(queue) > 0 {
		n := len(queue)

		for i := 0; i < n; i++ {
			if queue[i].Left == nil && queue[i].Right == nil {
				return level
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		level++
	}
	return 0
}
