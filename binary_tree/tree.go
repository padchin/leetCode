package main

import "fmt"

type TreeNode struct {
	Key   int
	Left  *TreeNode
	Right *TreeNode
}

func insert(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return &TreeNode{Key: key, Left: nil, Right: nil}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
	} else if key > root.Key {
		root.Right = insert(root.Right, key)
	}

	return root
}

func buildBST(pairs [][]int) *TreeNode {
	var root *TreeNode
	for _, pair := range pairs {
		parent, child := pair[0], pair[1]
		if root == nil {
			root = &TreeNode{Key: parent, Left: nil, Right: nil}
		}
		root = insert(root, child)
	}
	return root
}

func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Printf("%d ", root.Key)
		inorderTraversal(root.Right)
	}
}

func main() {
	pairs := [][]int{
		{0, 1}, {0, 2}, {1, 3}, {1, 4},
		{2, 5}, {2, 6}, {3, 7}, {3, 8},
		{4, 9}, {4, 10}, {5, 11}, {5, 12},
		{6, 13}, {6, 14},
	}

	root := buildBST(pairs)
	fmt.Println("Inorder traversal of the BST:")
	inorderTraversal(root)
}
