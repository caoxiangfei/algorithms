package main

import (
	"fmt"
	"sync"
)

var m int = 100

var lock sync.Mutex

var wt sync.WaitGroup

func a() {
	defer wt.Done()
	lock.Lock()
	//m += 1
	fmt.Println("a")
	//time.Sleep(time.Millisecond * 10)
	lock.Unlock()
}

func b() {
	defer wt.Done()
	lock.Lock()
	//time.Sleep(time.Millisecond * 2)
	//m -= 1
	fmt.Println("b")
	lock.Unlock()
}

func main() {

	for i := 0; i < 100; i++ {
		go a()
		wt.Add(1)
		go b()
		wt.Add(1)
	}

	wt.Wait()
	//fmt.Printf("m: %v\n", m)
}
