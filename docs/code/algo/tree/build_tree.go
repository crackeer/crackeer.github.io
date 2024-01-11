func buildTree(data []int) *Node {
	if len(data) < 1 {
		return nil
	}
	root := &Node{
		Value: data[0],
	}
	queue := []*Node{root}
	from := 1
	for {
		curNode := queue[0]
		if from > len(data)-1 {
			break
		}
		curNode.Left = &Node{
			Value: data[from],
		}
		queue = append(queue, curNode.Left)
		from++
		if from > len(data)-1 {
			break
		}
		curNode.Right = &Node{
			Value: data[from],
		}
		queue = append(queue, curNode.Right)
		from++
		queue = queue[1:]
	}
	return root
}