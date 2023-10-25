package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := make([]int, 10001)

	// Заполняем массив минимальными значениями
	minIntValue := -9223372036854775808 // Минимальное значение для 64-битной архитектуры
	for i := 0; i < len(res); i++ {
		res[i] = minIntValue
	}

	var level int
	var maxLevel int
	traverse(root, &res, level, &maxLevel)

	return res[:maxLevel+1]
}

func traverse(root *TreeNode, res *[]int, level int, maxLevel *int) {
	if root == nil {
		return
	}

	traverse(root.Left, res, level+1, maxLevel)
	if (*res)[level] < root.Val {
		(*res)[level] = root.Val
	}
	if *maxLevel < level {
		*maxLevel = level
	}
	traverse(root.Right, res, level+1, maxLevel)
}
