package dataStructures

import "fmt"

//单向链表
type SLNode struct {
	Value interface{}
	Next  *SLNode
}

func NewSLNode(value interface{}) *SLNode {
	return &SLNode{
		Value: value,
		Next:  nil,
	}
}

type SLLRepository interface {
	NewLinkList() *SingleLinkList
	CreateNode(value interface{}) *SLNode
	PopFront()
	PushFront(value interface{})
	PushBack(value interface{})
	PopBack()
	Delete(value interface{})
	PrintSingleList()
}

type SingleLinkList struct {
	head   *SLNode
	length uint
}

func NewSLLRepository() SLLRepository {
	return &SingleLinkList{}
}

func (sl *SingleLinkList) NewLinkList() *SingleLinkList {

	return &SingleLinkList{
		head:   &SLNode{0, nil},
		length: 0,
	}
}

func (sl *SingleLinkList) CreateNode(value interface{}) *SLNode {
	return NewSLNode(value)
}

// PopFront 头部弹出
func (sl *SingleLinkList) PopFront() {
	//将头节点的指针指向第一个节点的下一个节点
	sl.head.Next = sl.head.Next.Next
	//将第一个节点的next指向nil
	sl.head.Next.Next = nil
	sl.length--
}

// PushFront 头部添加
func (sl *SingleLinkList) PushFront(value interface{}) {
	//1. 创建节点
	node := NewSLNode(value)
	//2. 将头部节点当前指向的内存地址赋给新节点
	node.Next = sl.head.Next
	//3. 将头部节点的指向改为新节点
	sl.head.Next = node
	sl.length++
}

// PushBack 尾部添加 没有tail
func (sl *SingleLinkList) PushBack(value interface{}) {
	//1. 创建节点
	node := NewSLNode(value)
	//2. 定义变量存储头节点
	cur := sl.head
	//3. 遍历链表
	for cur.Next != nil {
		cur = cur.Next
	}
	//4. 此时已经来到链表尾部 将尾部节点指向新节点。
	//复杂度O(n) 如果有tail，O(1)
	cur.Next = node
	sl.length++
}

// PopBack 尾部删除
func (sl *SingleLinkList) PopBack() {

	//1. 定义变量存储头节点
	cur := sl.head
	//2. 遍历链表
	for cur.Next.Next != nil {
		cur = cur.Next
	}
	//3. 此时已经来到链表尾部的前一个元素
	//复杂度O(n) 如果有tail，O(1)
	cur.Next = nil
	sl.length--
}

//删除某一特定值的节点
func (sl *SingleLinkList) Delete(value interface{}) {
	//1. 定义变量存储头节点
	cur := sl.head
	//2. 遍历链表
	for cur.Next != nil {
		if cur.Next.Value == value {
			if cur.Next.Next != nil {
				cur.Next = cur.Next.Next
				break
			} else {
				cur.Next = nil
				break
			}
		}
		cur = cur.Next
	}

	//复杂度O(n) 如果有tail，O(1)
	//cur.Next = nil
	sl.length--
}

func (sl *SingleLinkList) PrintSingleList() {
	fmt.Println("The length of list is ", sl.length)
	cur := sl.head
	for cur.Next != nil {
		fmt.Println(cur.Value)
		cur = cur.Next
	}
}

/*
PushFront(Key) add to front
Key TopFront() return front item
PopFront() remove front item
PushBack(Key) add to back
Key TopBack() return the last item
PopBack() remove back item
Boolean Find(Key) is key in list?
Erase(Key) remove key from list
Boolean Empty() empty list?
AddBefore(Node, Key) adds key before node
AddAfter(Node, Key) adds key after node
*/
