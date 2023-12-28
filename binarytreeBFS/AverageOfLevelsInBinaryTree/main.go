package main

func main() {
	root := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 20,
			},
			Right: &TreeNode{
				Val: 1,
			},
		},
		Right: &TreeNode{
			Val: 10,
			Right: &TreeNode{
				Val: -25,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
	}
	averageOfLevels(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{0}
	}

	var results []float64

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		n := len(queue)
		var average float64

		for i := 0; i < n; i++ {
			current := queue[0]
			queue = queue[1:]
			average += float64(current.Val)
			if current.Left != nil {
				queue = append(queue, current.Left)
			}

			if current.Right != nil {
				queue = append(queue, current.Right)
			}
		}
		average = average / float64(n)
		results = append(results, average)
	}

	return results
}
