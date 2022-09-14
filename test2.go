package main

import "fmt"

//set 的实现方式参考了网上的众多资料，其中一位博客的文章我感觉写的非常简单易懂
//就利用这篇文章的内容来实现set了
//csdn：https://blog.csdn.net/HaoDaWang?type=blog
//文章名称：【golang】Go实现set类型

type Empty struct{}

//创建set类
type MySet struct {
	M map[int]Empty
}

func main() {

	var n int
	fmt.Scanln(&n)

	mess := make([]int, n)
	x := make([]int, n)

	//读取mess
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		mess[i] = m
	}
	//读取x
	for i := 0; i < n; i++ {
		var temp int
		fmt.Scan(&temp)
		x[i] = temp
	}

	//测试输入情况
	//fmt.Println(n)
	//fmt.Println(mess)
	//fmt.Println(x)

	//将所有元素装入切片
	elements := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < x[i]; j++ {
			elements = append(elements, mess[i])
		}
	}

	//fmt.Println(elements)

	//创建set
	ms := map[int]Empty{}
	ms[0] = Empty{}

	for _, v := range elements {
		preverse := map[int]Empty{}
		//深度copy
		for k := range ms {
			preverse[k] = Empty{}
		}

		l := len(preverse)

		//fmt.Println(v)

		counter := 0
		for k := range preverse {

			if counter < l {
				val := k + v
				ms[val] = Empty{}
				//fmt.Println(val)
			}
			counter++
		}
		//fmt.Println("###")
	}

	fmt.Print(len(ms))
	fmt.Print(ms)

}
