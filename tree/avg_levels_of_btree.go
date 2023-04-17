package tree

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	result := make([]float64, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			levelSum += node.Val
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, float64(levelSum)/float64(levelSize))
	}

	return result
}
