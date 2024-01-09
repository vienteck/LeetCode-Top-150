package main

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	hasCycle(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	tracker := map[*ListNode]struct{}{}

	for head != nil {
		if _, exists := tracker[head]; exists {
			return true
		} else {
			tracker[head] = struct{}{}
		}
		head = head.Next
	}
	return false
}
