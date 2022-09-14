//遍历所有数字
//如果数字被7整除 或者含有7 就加1
//将这两个判断单独做成方法
package main

import "fmt"
import "strings"
import "strconv"

func isTimesOfSeven(num int) bool {
	if num%7 == 0 {
		return true
	} else {
		return false
	}
}

func isContainSeven(num int) bool {

	numStr := strconv.Itoa(num)
	return strings.Contains(numStr, "7")
}

func main() {
	var n int
	fmt.Scan(&n)

	counter := 0
	for i := 1; i < n+1; i++ {
		if isTimesOfSeven(i) || isContainSeven(i) {
			counter++
		}
	}

	fmt.Println(counter)
}
