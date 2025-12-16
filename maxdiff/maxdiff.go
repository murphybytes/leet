package maxdiff

import (
	"math"

	"github.com/murphybytes/leet/collections"
)

func MaxAncestorDiff(root *collections.TreeNode) int {
	currMax := math.MinInt
	currMin := math.MaxInt
	return mad(root, currMin, currMax)
}

func mad(root *collections.TreeNode, currMin, currMax int) int {
	if root == nil {
		diff := currMax - currMin
		if diff < 0 {
			diff *= -1
		}
		return diff
	}

	if root.Val > currMax {
		currMax = root.Val
	}

	if root.Val < currMin {
		currMin = root.Val
	}

	left := mad(root.Left, currMin, currMax)
	right := mad(root.Right, currMin, currMin)

	return max(left, right)
}
