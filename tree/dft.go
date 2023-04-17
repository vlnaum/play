package tree

func dft(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	stack := []*TreeNode{root}
	result := make([]int, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		result = append(result, node.Val)
		stack = stack[:len(stack)-1]

		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
