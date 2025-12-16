package mindepth

import (
	"fmt"
	"math"

	"github.com/murphybytes/leet/collections"
)

func minDepth(root *collections.TreeNode) int {
	return 0
}

func buildTree(ar []int) *collections.TreeNode {
	return bt(0, ar)
}

func bt(i int, ar []int) *collections.TreeNode {
	fmt.Printf("i=%d\n", i)

	if i >= len(ar) {
		return nil
	}

	if ar[i] == math.MinInt {
		return nil
	}

	root := &collections.TreeNode{Val: ar[i]}

	leftIdx := (2 * i) + 1
	rightIdx := (2 * i) + 2

	root.Left = bt(leftIdx, ar)
	root.Right = bt(rightIdx, ar)
	fmt.Printf("node %v\n", root)
	return root
}
