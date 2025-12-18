// Package depthsum sums the values of the nodes that the
// deepest level of a binary tree
package depthsum

import "github.com/murphybytes/leet/collections"

func DeepestLeavesSum(root *collections.TreeNode) int {
	var rowSum int

	var queue collections.Queue[*collections.TreeNode]
	queue.PushBack(root)

	for !queue.Empty() {
		nodesInRow := queue.Size()
		rowSum = 0

		for range nodesInRow {
			node := queue.Front()
			queue.PopFront()

			rowSum += node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}

		}
	}

	return rowSum
}
