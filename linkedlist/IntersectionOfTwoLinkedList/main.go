package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tracker := map[*ListNode]struct{}{}
	for headA != nil {
		if _, exists := tracker[headA]; !exists {
			tracker[headA] = struct{}{}
		}
		headA = headA.Next
	}

	for headB != nil {
		if _, exists := tracker[headB]; exists {
			return headB
		}
		headB = headB.Next
	}

	return nil
}
