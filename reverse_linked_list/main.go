package main

// Node Определяем структуру элемента списка
type Node struct {
	Data int
	Next *Node
}

func reverseLinkedList(head *Node) *Node {
	var nodesData []int

	next := head
	nodesData = append(nodesData, next.Data)

	for {
		if next.Next == nil {
			break
		}
		next = next.Next
		nodesData = append(nodesData, next.Data)

	}

	node := &Node{
		Data: nodesData[len(nodesData)-1],
		Next: nil,
	}

	prevNode := node

	for i := len(nodesData) - 2; i >= 0; i-- {
		next := &Node{}
		next.Data = nodesData[i]
		prevNode.Next = next
		prevNode = next
	}

	return node
}
