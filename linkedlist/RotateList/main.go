package main

func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	rotateRight(head, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	if head.Next == nil {
		return head
	}

	current := head
	holder := head
	var answer *ListNode
	count := 1
	for current.Next != nil {
		count++
		current = current.Next
	}
	current = head
	k = k % count
	if k == 0 {
		return head
	}
	for i := 1; i < k; i++ {
		for current.Next.Next != nil {
			current = current.Next
		}
		answer = current.Next
		current.Next = nil
		answer.Next = holder
		current = answer
		holder = answer
	}

	return answer
}
