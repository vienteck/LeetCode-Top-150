package main

func main() {
	l := &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}
	r := &TreeNode{Val: 3, Right: &TreeNode{Val: 5}}
	root := &TreeNode{Val: 1, Left: l, Right: r}
	zigzagLevelOrder(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	answer := [][]int{}
	row := 0
	dir := -1
	for len(queue) > 0 {
		answer = append(answer, []int{})
		n := len(queue)

		for i := 0; i < n; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		if dir > 0 {
			for i := n - 1; i >= 0; i-- {
				answer[row] = append(answer[row], queue[i].Val)
			}
		} else {
			for i := 0; i < n; i++ {
				answer[row] = append(answer[row], queue[i].Val)
			}
		}

		dir *= -1
		queue = queue[n:]
		row++
	}
	return answer
}
