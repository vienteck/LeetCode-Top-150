package main

func main() {
	a := &Node{Val: 7}
	b := &Node{Val: 13}
	c := &Node{Val: 11}
	d := &Node{Val: 10}
	e := &Node{Val: 1}
	a.Next = b
	a.Random = nil
	b.Next = c
	b.Random = a
	c.Next = d
	c.Random = e
	d.Next = e
	d.Random = c
	e.Next = nil
	e.Random = a

	copyRandomList(a)
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Create a mapping of original nodes to their corresponding copied nodes
	nodeMap := make(map[*Node]*Node)

	// First pass: Create nodes and link them in the 'Next' direction
	current := head
	for current != nil {
		nodeMap[current] = &Node{Val: current.Val}
		current = current.Next
	}

	// Second pass: Link the 'Random' pointers
	current = head
	for current != nil {
		if current.Next != nil {
			nodeMap[current].Next = nodeMap[current.Next]
		}
		if current.Random != nil {
			nodeMap[current].Random = nodeMap[current.Random]
		}
		current = current.Next
	}

	// The head of the deep copy is the node corresponding to the head of the original list
	return nodeMap[head]
}
