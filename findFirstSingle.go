//用一个map记录重复的字符串
//遍历字符串找到第一个不重复的
package main

import (
	"fmt"
)

type Emptyr struct{}

func main() {
	var str string
	fmt.Scan(&str)

	//1. 建立一个map当set使用
	em := Emptyr{}
	myset := make(map[string]Emptyr)

	//2. 遍历字符串 如果字符不在map里, 就添加,然后遍历字符看看是否重复
	//如果在map里 就取出下一个字符向后遍历
	counter := 0
	flag2 := true

	for counter < len(str) {
		flag := true
		temp := str[counter : counter+1]
		_, ok := myset[temp]
		if !ok {
			myset[temp] = em

			for i := counter + 1; i < len(str); i++ {
				if str[i:i+1] == temp {
					flag = false
					flag2 = flag
					break
				}
			}
			if flag {
				fmt.Println(temp)
				break
			}

		}
		counter++
	}
	if !flag2 {
		fmt.Println(-1)
	}
}
