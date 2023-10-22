package main

import "fmt"

type treeNode struct {
	data  int
	left  *treeNode
	right *treeNode
}

var pairs = [][]int{
	{0, 1}, {0, 2}, {1, 3}, {1, 4},
	{2, 5}, {2, 28}, {3, 7}, {3, 8},
	{4, 9}, {4, 10}, {5, 17}, {5, 12},
	{6, 33}, {6, 14},
}

func insert2(root *treeNode, data int) *treeNode {
	if root == nil {
		return &treeNode{data: data}
	}

	if data < root.data {
		root.left = insert2(root.left, data)
	} else if data > root.data {
		root.right = insert2(root.right, data)
	}

	return root
}

func buildBST1(in [][]int) *treeNode {
	var tree *treeNode

	for _, pair := range in {
		parent, child := pair[0], pair[1]
		if tree == nil {
			tree = insert2(tree, parent)
		}

		tree = insert2(tree, child)
	}

	return tree
}

func traverse(tree *treeNode) {
	if tree == nil {
		return
	}
	traverse(tree.left)
	fmt.Println(tree.data)
	traverse(tree.right)
}

func main() {
	traverse(buildBST1(pairs))
}
