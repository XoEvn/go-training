package hash

var tempMap = map[int]int{}

func numWays(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 2 {
		return n
	}

	if value, ok := tempMap[n]; ok {
		return value
	} else {
		tempMap[n] = (numWays(n-1) + numWays(n-2)) % 1000000007
		return tempMap[n]
	}

}

type treeNode struct {
	Val   int
	Left  *treeNode
	Right *treeNode
}

func SetValue(val int, node *treeNode) {
	node.Val = val
}

func invertTree(root *treeNode) *treeNode {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}
