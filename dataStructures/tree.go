package dataStructures

import "fmt"

//1. 树的实现。采用链表的方式
//树的实现网上有很多种方式，有数组的。但是我参考了csdn一位博主的，采用了链表的node节点方式
//csdn：https://blog.csdn.net/mrbone9?type=blog
//感觉这种更通俗易懂一点
//2. 具体方法
//2.1 创建节点结构体
//2.2 定义节点初始化函数
//2.3 定义构造树函数，将节点从root节点开始，以广度搜索的顺序排开，然后建立树
//         0
//     1        2
//   3   4    5   6
//  7 8 9 10
//[0 1 2 3 4 5 6 7 8 9 10]

type TreeNode struct {
	Value interface{}
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root   *TreeNode
	Height int
}

type TreeFactory interface {
	CreateNode(value interface{}) *TreeNode
	initTree(node *TreeNode) *Tree
	calculateHeight(root *TreeNode) int
	CreateTree([]*TreeNode) *Tree
	PreOrder(node *TreeNode)
	InOrder(node *TreeNode)
	PostOrder(node *TreeNode)
	Bfs(node *TreeNode)
}

func NewTreeFactory() TreeFactory {
	return &Tree{}
}

// CreateNode 构建节点
func (t *Tree) CreateNode(value interface{}) *TreeNode {
	return &TreeNode{value, nil, nil}
}

//初始化树
func (t *Tree) initTree(node *TreeNode) *Tree {
	t.Root = node
	t.Height = 0
	return t
}

//计算树高
func (t *Tree) calculateHeight(root *TreeNode) (h int) {

	if root == nil {
		return 0
	} else {
		left := t.calculateHeight(root.Left)
		right := t.calculateHeight(root.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

// CreateTree  初始化树
func (t *Tree) CreateTree(nodes []*TreeNode) *Tree {
	for i := 0; i < len(nodes); i++ {
		if 2*i+1 < len(nodes) {
			nodes[i].Left = nodes[2*i+1]
			nodes[i].Right = nodes[2*i+2]
		}
	}

	t.initTree(nodes[0])
	t.Height = t.calculateHeight(nodes[0])

	return t
}

// PreOrder 前序遍历
func (t *Tree) PreOrder(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Print(node.Value)
	t.PreOrder(node.Left)
	t.PreOrder(node.Right)
}

// InOrder 中序遍历
func (t *Tree) InOrder(node *TreeNode) {
	if node == nil {
		return
	}

	t.InOrder(node.Left)
	fmt.Print(node.Value)
	t.InOrder(node.Right)
}

// PostOrder 后序遍历
func (t *Tree) PostOrder(node *TreeNode) {
	if node == nil {
		return
	}

	t.PostOrder(node.Left)
	t.PostOrder(node.Right)
	fmt.Print(node.Value)
}

// Bfs bfs
func (t *Tree) Bfs(node *TreeNode) {
	if node == nil {
		return
	}
	var q []*TreeNode
	q = append(q, node)
	for len(q) > 0 {
		temp := q[0]
		q = append(q[:0], q[1:]...)
		fmt.Print(temp.Value, " ")
		if temp.Left != nil {
			q = append(q, temp.Left)
		}
		if temp.Right != nil {
			q = append(q, temp.Right)
		}
	}
}
