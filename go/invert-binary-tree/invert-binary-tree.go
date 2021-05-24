package invert_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var tmp *TreeNode
	if root.Left != nil {
		root.Left = invertTreeWithTemp(root.Left, tmp)
	}
	if root.Right != nil {
		root.Right = invertTreeWithTemp(root.Right, tmp)
	}
	tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}

func invertTreeWithTemp(root *TreeNode, tmp *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = invertTreeWithTemp(root.Left, tmp)
	}
	if root.Right != nil {
		root.Right = invertTreeWithTemp(root.Right, tmp)
	}
	tmp = root.Left
	root.Left = root.Right
	root.Right = tmp
	return root
}
