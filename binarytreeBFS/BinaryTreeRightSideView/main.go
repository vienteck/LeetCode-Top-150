package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	queue := []*TreeNode{root}
	answer := []int{}
	for len(queue) > 0 {

		n := len(queue)
		visibleNumber := queue[0].Val
		for i := 0; i < n; i++ {
			visibleNumber = queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		answer = append(answer, visibleNumber)
	}

	return answer
}
