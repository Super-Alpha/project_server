package data_structure

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历：根->左->右
func PreOrderTraverse(rootNode *TreeNode) {
	if rootNode == nil {
		return
	}
	fmt.Println(rootNode.Val)

	PreOrderTraverse(rootNode.Left)

	PreOrderTraverse(rootNode.Right)
}

// 中序遍历：左->根->右
func MidOrderTraverse(rootNode *TreeNode) {
	if rootNode == nil {
		return
	}
	MidOrderTraverse(rootNode.Left)

	fmt.Println(rootNode.Val)

	MidOrderTraverse(rootNode.Right)
}

// 后序遍历：左->右->根
func PostOrderTraverse(rootNode *TreeNode) {
	if rootNode == nil {
		return
	}
	PostOrderTraverse(rootNode.Left)

	PostOrderTraverse(rootNode.Right)

	fmt.Println(rootNode.Val)
}

func depthOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	depthOrder(root.Left)
	depthOrder(root.Right)
}

// 层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {

		level := make([]int, 0)
		count := len(queue)

		for count > 0 {
			// 获取队首元素
			node := queue[0]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			level = append(level, node.Val)
			// 删除队首元素
			queue = queue[1:]
			count--
		}
		res = append(res, level)
	}
	return res
}

func Traverse() {
	root := &TreeNode{Val: 1}
	rLeft := &TreeNode{Val: 2}
	rRight := &TreeNode{Val: 3}
	root.Left = rLeft
	root.Right = rRight

	//PreOrderTraverse(root)
	//MidOrderTraverse(root)
	//PostOrderTraverse(root)
	//fmt.Println(levelOrder(root))
	depthOrder(root)
}
