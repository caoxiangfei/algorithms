package main

import "fmt"

//1.构建单向链表数据结构
type Single struct {
	Value int
	Next  *Single
}

//2.完成insert函数
func (s *Single) Insert(value int, position int) {
	cur := s
	for cur != nil {
		if cur.Value == position {
			temp := cur.Next
			cur.Next = &Single{
				Value: 0,
				Next:  nil,
			}
			cur.Next.Value = value
			cur.Next.Next = temp
		}
		cur = cur.Next
	}
}

//3.完成delete函数
//找到值为value的点，将其前一节点直接连到其下一节点
//如果value为头节点？
//如果value为尾节点？
func (s *Single) Delete(value int) (result *Single) {
	if s.Value == value {
		return s.Next
	}

	cur := s
	for cur.Next != nil {
		if cur.Next.Value == value {
			if cur.Next.Next == nil {
				cur.Next = nil

				return s
			} else {
				cur.Next = cur.Next.Next
				return s
			}
		}
		cur = cur.Next
	}
	return nil
}

func main() {
	//4.根据输入完成单向链表构建
	//输入总数n
	var n int
	//头节点值head
	var head int

	//value节点插入到position节点后面
	var value, position int
	//需删除的节点del
	var del int

	fmt.Scan(&n)
	fmt.Scan(&head)

	var singleList Single
	singleList.Value = head
	s := &singleList
	for i := 0; i < n-1; i++ {
		fmt.Scan(&value)
		fmt.Scan(&position)
		s.Insert(value, position)
	}

	fmt.Scan(&del)
	singleList.Delete(del)

	cur := &singleList
	for cur.Next != nil {
		fmt.Print(cur.Value, " ")
		cur = cur.Next
	}
	fmt.Println(cur.Value)

}
