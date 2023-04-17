package tree

func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}

	for len(queue) > 0 {
		nodeP, nodeQ := queue[0], queue[1]
		queue = queue[2:]

		if nodeP == nil && nodeQ == nil {
			continue
		}

		if nodeP == nil || nodeQ == nil || nodeP.Val != nodeQ.Val {
			return false
		}

		queue = append(queue, nodeP.Left, nodeQ.Left, nodeP.Right, nodeQ.Right)
	}

	return true
}
