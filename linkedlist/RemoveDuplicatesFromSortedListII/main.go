package main

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}}}
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}
	deleteDuplicates(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		if current.Next.Val == current.Next.Next.Val {
			val := current.Next.Val
			for current.Next != nil && current.Next.Val == val {
				current.Next = current.Next.Next
			}
		} else {
			current = current.Next
		}
	}

	return dummy.Next
}

func myVersion(head *ListNode) *ListNode {
	answer := &ListNode{}
	dummy := answer
	current := head

	tracker := map[int]int{}
	for current != nil {
		if _, exists := tracker[current.Val]; !exists {
			tracker[current.Val] = 1
		} else {
			tracker[current.Val]++
		}
		current = current.Next
	}

	current = head
	for current != nil {
		if _, exists := tracker[current.Val]; exists {
			count := tracker[current.Val]
			if count > 1 {
				current = current.Next
				continue
			} else {
				dummy.Next = current
				dummy = dummy.Next
			}
		}
		current = current.Next
	}
	dummy.Next = nil
	return answer.Next
}
