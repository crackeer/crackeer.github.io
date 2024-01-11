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