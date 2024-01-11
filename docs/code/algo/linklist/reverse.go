// 递归版本
func ReverseLinkList1(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseLinkList1(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}

// 非递归版本
func ReverseLinkList2(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		preNode  *Node
		curNode  *Node = head
		nextNode *Node
	)

	for curNode != nil {

		nextNode = curNode.Next

		curNode.Next = preNode

		preNode = curNode

		curNode = nextNode

	}

	return preNode
}