func printTree(root *Node) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	level := 1
	for {
		if len(queue) < 1 {
			break
		}
		fmt.Printf("level->%d: ", level)
		tmpQueue := []*Node{}
		for _, node := range queue {
			fmt.Printf("%d ", node.Value)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}

			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		queue = tmpQueue
		level++
		fmt.Println("")
	}
}