package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	answer := [][]int{}
	row := 0
	for len(queue) > 0 {
		answer = append(answer, []int{})
		n := len(queue)
		currentRow := []int{}
		for i := 0; i < n; i++ {
			currentRow = append(currentRow, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}

			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[n:]
		answer[row] = append(answer[row], currentRow...)
		row++
	}

	return answer
}
