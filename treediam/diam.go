package treediam

import "github.com/murphybytes/leet/collections"

func TreeDiam(root *collections.TreeNode) int {
	var diam int
	traverse(root, &diam)
	return diam
}

func traverse(root *collections.TreeNode, diam *int) int {
	if root == nil {
		return 0
	}

	left := traverse(root.Left, diam)
	right := traverse(root.Right, diam)
	*diam = max(*diam, left+right)
	return max(left, right) + 1
}
