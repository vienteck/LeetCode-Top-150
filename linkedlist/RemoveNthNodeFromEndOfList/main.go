package main

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	removeNthFromEnd(head, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil && n == 1 {
		return nil
	}

	getCount := 0
	current := head

	for current != nil {
		current = current.Next
		getCount++
	}

	target := getCount - n - 1
	current = head
	tracker := 0
	for current != nil {
		if tracker == target {
			if current.Next.Next != nil {
				current.Next = current.Next.Next
			} else {
				current.Next = nil
			}
			break
		} else {
			current = current.Next
			tracker++
		}
	}

	return head
}
