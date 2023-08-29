func middleTraverse2(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{}
	curNode := root

	for {
		if curNode == nil && len(queue) < 1 {
			break
		}
		for curNode != nil {
			queue = append(queue, curNode)
			curNode = curNode.Left
		}
		lastNode := queue[len(queue)-1]
		fmt.Printf("%d ", lastNode.Value)
		queue = queue[0 : len(queue)-1]
		curNode = lastNode.Right
	}
}