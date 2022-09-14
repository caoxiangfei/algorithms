package dataStructures

//AVL树里面的旋转感觉挺复杂 AVL树也是树在树的基础上进行改进就行了
//我的算法还是基于coursera上面的data structures  and algorithms 的ppt完成算法编写
//课程地址：https://www.coursera.org/specializations/data-structures-algorithms
//为了理解avl树的旋转 我参考了csdn上一位博主的文章，他写的很清晰
//csdn：https://blog.csdn.net/weixin_45818891?type=blog
//名称：二叉树进阶 ---AVL树的介绍 AVL树插入的4种情况 (详细图解+代码演示)

type AVLTreeNode struct {
	Parent *AVLTreeNode
	Key    int
	Left   *AVLTreeNode
	Right  *AVLTreeNode
	Height int
}

type AVLTree struct {
	Root *AVLTreeNode
}

type AVLFactory interface {
	AVLNode(key int) *AVLTreeNode
	InitAVL([]*AVLTreeNode) *AVLTree

	//用于计算height
	adjustHeight(node *AVLTreeNode)

	// Find 找到值为key的节点 如果没有节点的值为key 返回距离最近的一点
	Find(key int, root *AVLTreeNode) (node *AVLTreeNode)

	//找到一个节点，该节点的值刚刚大于node节点的值。例如node节点的值为3，树中大于3的值有5 7 9，那么刚刚大于该节点的值就是5
	//leftDescendant和rightAncestor服务Next函数
	Next(node *AVLTreeNode) *AVLTreeNode
	leftDescendant(leftChild *AVLTreeNode) *AVLTreeNode
	rightAncestor(rightParent *AVLTreeNode) *AVLTreeNode

	//找到一组节点，该组中的节点的值都处于x和y之间。
	Search(x, y int, root *AVLTreeNode) []*AVLTreeNode

	Insert(key int, root *AVLTreeNode)
	Delete(node *AVLTreeNode)
}

func NewAVLFactory() AVLFactory {
	return &AVLTree{}
}

func (avl *AVLTree) AVLNode(key int) *AVLTreeNode {
	return &AVLTreeNode{nil, key, nil, nil, 0}
}

func (avl *AVLTree) InitAVL(nodes []*AVLTreeNode) *AVLTree {
	for i := 0; i < len(nodes); i++ {
		if 2*i+1 < len(nodes) {
			nodes[i].Left = nodes[2*i+1]
			nodes[2*i+1].Parent = nodes[i]
			nodes[i].Right = nodes[2*i+2]
			nodes[2*i+2].Parent = nodes[i]
		}
	}
	for _, cur := range nodes {
		cur.Height = avl.adjustHeight(cur)
	}

	avl.Root = nodes[0]
	return avl
}

func (avl *AVLTree) adjustHeight(node *AVLTreeNode) (h int) {
	if node == nil {
		return 0
	} else {
		left := avl.adjustHeight(node.Left)
		right := avl.adjustHeight(node.Right)
		if left > right {
			return left + 1
		} else {
			return right + 1
		}
	}
}

func (avl *AVLTree) Find(key int, root *AVLTreeNode) (node *AVLTreeNode) {
	if root.Key == key {
		return root
	} else if root.Key > key {
		if root.Left != nil {
			return avl.Find(key, root.Left)
		}
		return root
	} else if root.Key < key {
		if root.Right != nil {
			return avl.Find(key, root.Right)
		}
		return root
	}
	return nil
}

func (avl *AVLTree) Next(node *AVLTreeNode) *AVLTreeNode {
	if node.Right != nil {
		return avl.leftDescendant(node.Right)
	} else {
		return avl.rightAncestor(node)
	}
}

func (avl *AVLTree) leftDescendant(node *AVLTreeNode) *AVLTreeNode {
	if node.Left == nil {
		return node
	} else {
		return avl.leftDescendant(node.Left)
	}
}

func (avl *AVLTree) rightAncestor(node *AVLTreeNode) *AVLTreeNode {

	if node.Key < node.Parent.Key {
		return node.Parent
	} else {
		return avl.rightAncestor(node.Parent)
	}
}

func (avl *AVLTree) Search(x, y int, root *AVLTreeNode) (re []*AVLTreeNode) {
	node := avl.Find(x, root)
	for node.Key <= y {
		if node.Key >= x {
			re = append(re, node)
		}
		node = avl.Next(node)
	}
	return re
}

func (avl *AVLTree) Insert(key int, root *AVLTreeNode) {
	node := avl.Find(key, avl.Root)
	child := &AVLTreeNode{
		Parent: node,
		Key:    key,
		Left:   nil,
		Right:  nil,
		Height: 1,
	}
	if key >= node.Key {
		node.Right = child
	} else {
		node.Right = child
	}
}
