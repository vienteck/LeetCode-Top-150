package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	pqueue := []*TreeNode{p}
	qqueue := []*TreeNode{q}

	for len(pqueue) > 0 {
		pcurrent := pqueue[0]
		pqueue = pqueue[1:]

		qcurrent := qqueue[0]
		qqueue = qqueue[1:]

		if pcurrent == nil && qcurrent == nil {
			continue
		} else if pcurrent == nil || qcurrent == nil {
			return false
		} else if pcurrent.Val != qcurrent.Val {
			return false
		}

		pqueue = append(pqueue, pcurrent.Left)

		qqueue = append(qqueue, qcurrent.Left)

		pqueue = append(pqueue, pcurrent.Right)
		qqueue = append(qqueue, qcurrent.Right)

	}
	return true
}
