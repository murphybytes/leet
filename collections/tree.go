package collections

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(nodes []any) *TreeNode {
	if len(nodes) == 0 || nodes == nil {
		return nil
	}
	root := &TreeNode{Val: nodes[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(nodes) {
		node := queue[0]
		queue = queue[1:]

		if i < len(nodes) && nodes[i] != nil {
			node.Left = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nodes) && nodes[i] != nil {
			node.Right = &TreeNode{Val: nodes[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
