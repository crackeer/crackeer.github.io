func preTraverse(root *Node) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for {
		if len(queue) < 1 {
			break
		}

		last := queue[len(queue)-1]
		fmt.Printf("%d ", last.Value)
		queue = queue[0 : len(queue)-1]
		if last.Right != nil {
			queue = append(queue, last.Right)
		}

		if last.Left != nil {
			queue = append(queue, last.Left)
		}
	}
}