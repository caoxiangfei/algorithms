package dataStructures

import (
	"fmt"
	"math"
)

//heap 存放在数组中 从根节点开始 按照广度搜索的顺序
//这里我根据coursera上面的课程data structures  and algorithms 的ppt完成算法编写
//课程地址：https://www.coursera.org/specializations/data-structures-algorithms
//看了他们的pseudo code 设计的非常好
//好像他们还在这里留了一点小缺陷

type PriorityQueue struct {
	PQ []int
}

type PQFactory interface {
	CreatePQ([]int)
	SiftUp(int)
	SiftDown(int)
	Insert(int2 int)
	ExtractMax() int
	Remove(int)
	ChangePriority(int, int)
	PrintHeap()
	HeapSort(hs []int) (re []int)
	BuildHeap(bh []int)
	HeapSort2(hs []int)
	PartialSorting(ps []int, k int) (re []int)
}

func NewPQFactory() PQFactory {
	return &PriorityQueue{}
}

func (pq *PriorityQueue) CreatePQ(p []int) {
	pq.PQ = p
}

func (pq *PriorityQueue) SiftUp(i int) {

	for i > 0 && pq.PQ[(i-1)/2] < pq.PQ[i] {
		pq.PQ[(i-1)/2], pq.PQ[i] = pq.PQ[i], pq.PQ[(i-1)/2]
		i = (i - 1) / 2
	}
}

func (pq *PriorityQueue) SiftDown(i int) {
	maxIndex := i
	left := 2*i + 1
	if left < len(pq.PQ) && pq.PQ[left] > pq.PQ[maxIndex] {
		maxIndex = left
	}
	right := 2*i + 2
	if right < len(pq.PQ) && pq.PQ[right] > pq.PQ[maxIndex] {
		maxIndex = right
	}
	if i != maxIndex {
		pq.PQ[i], pq.PQ[maxIndex] = pq.PQ[maxIndex], pq.PQ[i]
		pq.SiftDown(maxIndex)
	}
}

func (pq *PriorityQueue) Insert(value int) {
	size := len(pq.PQ)
	pq.PQ = append(pq.PQ, value)
	pq.SiftUp(size)
}

func (pq *PriorityQueue) ExtractMax() int {
	result := pq.PQ[0]

	pq.PQ[0] = pq.PQ[len(pq.PQ)-1]
	pq.PQ = pq.PQ[:len(pq.PQ)-1]

	pq.SiftDown(0)
	return result
}

func (pq *PriorityQueue) Remove(index int) {
	pq.PQ[index] = int(math.Inf(1))
	pq.SiftUp(index)
	pq.ExtractMax()
}

func (pq *PriorityQueue) ChangePriority(index int, value int) {
	oldValue := pq.PQ[index]
	pq.PQ[index] = value
	if value > oldValue {
		pq.SiftUp(index)
	} else {
		pq.SiftDown(index)
	}
}

func (pq *PriorityQueue) PrintHeap() {
	fmt.Print(pq.PQ)
}

func (pq *PriorityQueue) HeapSort(hs []int) (re []int) {
	pq.CreatePQ([]int{})
	for i := 0; i < len(hs); i++ {
		pq.Insert(hs[i])
	}
	for j := len(hs) - 1; j >= 0; j-- {
		hs[j] = pq.ExtractMax()
	}

	return hs
}

func (pq *PriorityQueue) BuildHeap(bh []int) {
	pq.CreatePQ(bh)
	size := len(bh) - 1
	for i := size / 2; i >= 0; i-- {
		pq.SiftDown(i)
	}
}

func (pq *PriorityQueue) HeapSort2(hs []int) {
	pq.BuildHeap(hs)
	for size := len(hs) - 1; size > 0; size-- {
		pq.PQ[0], pq.PQ[size] = pq.PQ[size], pq.PQ[0]
		pq.SiftDown(0)
	}
}

func (pq *PriorityQueue) PartialSorting(ps []int, k int) (re []int) {
	pq.BuildHeap(ps)
	for i := 0; i < k; i++ {
		re = append(re, pq.PQ[0])
		pq.ExtractMax()
	}
	return re
}
