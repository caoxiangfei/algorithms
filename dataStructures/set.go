package dataStructures

import "fmt"

//set 的实现方式参考了网上的众多资料，其中一位博客的文章我感觉写的非常简单易懂
//就利用这篇文章的内容来实现set了
//csdn：https://blog.csdn.net/HaoDaWang?type=blog
//文章名称：【golang】Go实现set类型
type Empty struct{}

var em Empty

// SetFactory 方法工厂
type SetFactory interface {
	InitSet()
	Add(val int)
	Remove(val int)
	Len() int
	Clear()
	Traverse()
}

//创建set类
type MySet struct {
	m map[int]Empty
}

func NewSetFactory() SetFactory {
	return &MySet{}
}

func (s *MySet) InitSet() {
	s.m = make(map[int]Empty)
}

//添加元素
func (s *MySet) Add(val int) {
	s.m[val] = em
}

//删除元素
func (s *MySet) Remove(val int) {
	delete(s.m, val)
}

//获取长度
func (s *MySet) Len() int {
	return len(s.m)
}

//清空set
func (s *MySet) Clear() {
	s.m = make(map[int]Empty)
}

//遍历set
func (s *MySet) Traverse() {
	for v := range s.m {
		fmt.Println(v)
	}
}
