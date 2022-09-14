package dataStructures

import (
	"fmt"
)

//双向链表
//我在这里实现了链表的部分方法，完整的方法也可以参考如下
//这是用于学习的编码，不完美，不严谨，仅供交流学习
/*
PushFront(Key) add to front
Key TopFront() return front item
PopFront() remove front item
PushBack(Key) add to back
Key TopBack() return back item
PopBack() remove back item
Boolean Find(Key) is key in list?
Erase(Key) remove key from list
Boolean Empty() empty list?
AddBefore(Node, Key) adds key before node
AddAfter(Node, Key) adds key after node
*/

type DLNode struct {
	Value interface{}
	Next  *DLNode
	Prev  *DLNode
}

// NewDLNode 定义一个函数用于创建节点
func newDLNode(value interface{}) *DLNode {
	return &DLNode{
		Value: value,
		Next:  nil,
		Prev:  nil,
	}
}

type DLLRepository interface {
	CreateNode(value interface{}) *DLNode
	NewDoubleLinkList() *DoubleLinkList
	IsEmpty() bool
	PushFront(node *DLNode)
	PopFront() error
	PushBack(node *DLNode)
	PopBack() error
	PrintList() error
}

type DoubleLinkList struct {
	head   *DLNode
	tail   *DLNode
	length uint
}

func NewDLLRepository() DLLRepository {
	return &DoubleLinkList{}
}

func (dll *DoubleLinkList) CreateNode(value interface{}) *DLNode {
	return newDLNode(value)
}

// NewDoubleLinkList 定义了一个空链表，只有头部节点head和尾部节点tail。
//此时没有中间节点所以head.Next指向tail, tail.Prev 指向head, 长度为0
//链表操作都是基于此链表
func (dll *DoubleLinkList) NewDoubleLinkList() *DoubleLinkList {
	//1.初始化头部节点和尾部节点
	theHead := dll.CreateNode(0)
	theTail := dll.CreateNode(0)
	//2.初始化链表
	newDoubleLinkList := DoubleLinkList{
		head:   theHead,
		tail:   theTail,
		length: 0,
	}
	//3.头尾指针操作 head.Next指向tail, tail.Prev 指向head, 长度为0
	newDoubleLinkList.head.Next = newDoubleLinkList.tail
	newDoubleLinkList.tail.Prev = newDoubleLinkList.head
	newDoubleLinkList.length = 0

	return &newDoubleLinkList
}

//基于此链表我们如何进行元素操作？
//1. 判断链表是否为空
//2. 链表为空的话，向头部添加节点node就是把head.Next指向node;
//   然后修改tail.Prev指向node
//   链表为空时的头部添加、尾部添加、元素插入三种方法都是同一种操作 定义一个函数叫firstNodeOperation
//3. 此时我们应该将length加1
//4. 如果链表不为空，可以利用head，tail和length进行添加删除操作

//判断链表是否为空
func (dll *DoubleLinkList) IsEmpty() bool {
	if dll.length == 0 {
		return true
	}
	return false
}

//头部添加节点
func (dll *DoubleLinkList) PushFront(node *DLNode) {
	if dll.IsEmpty() {
		dll.firstNodeOperation(node)
	} else {
		dll.head.Next.Prev = node
		node.Next = dll.head.Next
		dll.head.Next = node
		dll.length++
	}
}

//头部删除节点
func (dll *DoubleLinkList) PopFront() error {
	if dll.IsEmpty() {
		return fmt.Errorf("no nodes could pop")
	} else {
		dll.head.Next = dll.head.Next.Next
		dll.head.Next.Prev = nil
		dll.length--
		//TODO 那个被删除的节点还在吗？是否还占用着内存？
	}
	return nil
}

//尾部添加
func (dll *DoubleLinkList) PushBack(node *DLNode) {
	if dll.IsEmpty() {
		dll.firstNodeOperation(node)
	} else {
		node.Prev = dll.tail.Prev
		dll.tail.Prev.Next = node
		dll.tail.Prev = node
		dll.length++
	}
}

//尾部删除
func (dll *DoubleLinkList) PopBack() error {
	if dll.IsEmpty() {
		return fmt.Errorf("no nodes could pop")
	} else {
		dll.tail.Prev.Prev.Next = nil
		dll.tail.Prev = dll.tail.Prev.Prev
		dll.length--
		//TODO 那个被删除的节点还在吗？是否还占用着内存？
	}
	return nil
}

func (dll *DoubleLinkList) PrintList() error {
	if dll.IsEmpty() {
		return fmt.Errorf("no data")
	} else {
		cur := dll.head
		fmt.Println("The length of list is ", dll.length)
		for cur.Next != nil {
			fmt.Println(cur.Next.Value)
			cur = cur.Next
		}
	}
	return nil
}

//firstNodeOperation
func (dll *DoubleLinkList) firstNodeOperation(firstNode *DLNode) *DoubleLinkList {
	dll.head.Next = firstNode
	dll.tail.Prev = firstNode
	dll.length++
	//firstNode.Prev = dll.head
	//firstNode.Next = dll.tail
	return dll
}
