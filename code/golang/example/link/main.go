package main

import "fmt"

type Node struct {
	Next  *Node
	Value int64
}

func main() {
	var data []int64
	for i := 1; i < 11; i++ {
		data = append(data, int64(i))
	}
	head := BuildLinkList(data)
	list := TraverseLinkList(ReverseLinkByStep(head, 3))
	for _, value := range list {
		fmt.Println(value)
	}
}

// BuildLinkList
//
//	@param list
//	@return *Node
func BuildLinkList(list []int64) *Node {
	if len(list) < 1 {
		return nil
	}

	head := &Node{
		Value: list[0],
	}
	var curNode *Node = head
	for _, value := range list[1:] {
		curNode.Next = &Node{
			Value: value,
		}
		curNode = curNode.Next
	}
	return head
}

// TraverseLinkList
//
//	@param head
//	@return []int64
func TraverseLinkList(head *Node) []int64 {
	if head == nil {
		return nil
	}

	var (
		curNode *Node = head
		list    []int64
	)

	for curNode != nil {
		list = append(list, curNode.Value)
		curNode = curNode.Next
	}
	return list
}

// ReverseLinkList1
//
//	@param head
//	@return *Node
func ReverseLinkList1(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseLinkList1(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// ReverseLinkListK
//  @param head
//  @param k
//  @return *Node
func ReverseLinkListK(head *Node, k int64) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	if k == 1 {
		return head
	}
	newHead := ReverseLinkListK(head.Next, k-1)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

func ReverseLinkByStep(head *Node, step int64) *Node {
	if step <= 1 {
		return head
	}

	var (
		curNode *Node = head
		newHead *Node = head
	)
	for i := 0; i < int(step); i++ {
		if curNode != nil {
			curNode = curNode.Next
		} else {
			return head
		}
	}
	newHead = ReverseLinkListK(head, step)
	head.Next = ReverseLinkByStep(curNode, step)
	return newHead
}
