func backTraverse(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	list := []*Node{}

	for {
		if len(queue) < 1 {
			break
		}
		first := queue[len(queue)-1]
		queue = queue[0 : len(queue)-1]
		if first.Left != nil {
			queue = append(queue, first.Left)
		}
		if first.Right != nil {
			queue = append(queue, first.Right)
		}
		list = append(list, first)
	}
	for i := range list {
		fmt.Printf("%d ", list[len(list)-i-1].Value)
	}
}