package main

import "fmt"

func main() {
	//head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	removeNthFromEnd(head, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Create a dummy node to handle the case where the head needs to be removed
	dummy := &ListNode{Next: head}
	ahead, behind := dummy, dummy

	// Move ahead n+1 steps
	for i := 0; i <= n; i++ {
		ahead = ahead.Next
	}

	// Move both ahead and behind until ahead reaches the end
	for ahead != nil {
		ahead = ahead.Next
		behind = behind.Next
	}

	// Remove the nth node from the end
	behind.Next = behind.Next.Next

	return dummy.Next
}

func myVersion(head *ListNode, n int) *ListNode {
	getCount := 0
	current := head

	for current != nil {
		current = current.Next
		getCount++
	}

	if getCount == 1 && n == 1 {
		return nil
	}
	if getCount == 2 {
		if n == 2 {
			return head.Next
		} else {
			head.Next = nil
			return head
		}
	}
	if getCount == n {
		head = head.Next
		return head
	}
	target := getCount - n
	current = head
	tracker := 0
	prev := head
	for current != nil {
		if tracker == target {
			if current.Next != nil {
				//current = current.Next
				prev.Next = current.Next
			} else {
				prev.Next = nil
			}
			break
		} else {
			prev = current
			current = current.Next
			tracker++
		}
	}
	fmt.Println(prev)
	return head
}
