package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy

	// Move to the left-th node
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	// Reverse the sublist from left to right
	current := prev.Next
	var next *ListNode

	for i := 0; i < right-left+1; i++ {
		temp := current.Next
		current.Next = next
		next = current
		current = temp
	}

	// Connect the reversed sublist back to the original list
	prev.Next.Next = current
	prev.Next = next

	return dummy.Next
}

func main() {
	// Example usage:
	// Original linked list: 1 -> 2 -> 3 -> 4 -> 5
	// Reverse sublist from position 2 to 4
	// Resulting linked list: 1 -> 4 -> 3 -> 2 -> 5

	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	left, right := 2, 4
	result := reverseBetween(head, left, right)

	// Print the result
	for result != nil {
		println(result.Val)
		result = result.Next
	}
}
