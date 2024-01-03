package main

func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	rotateRight(head, 4)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {

	count := -1
	dummy := head
	for dummy != nil {
		count++
		dummy = dummy.Next
	}
	rotations := k % count
	dummy = head
	for i := 0; i <= rotations-1; i++ {
		dummy = dummy.Next
	}
	newHead := dummy.Next
	dummy.Next = nil
	dummy = newHead
	for dummy.Next != nil {
		dummy = dummy.Next
	}
	dummy.Next = head
	return newHead
}
