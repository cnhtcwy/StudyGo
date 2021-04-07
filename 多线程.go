package main

import (
	"fmt"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}
func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	//time.Sleep(1*time.Second)
	go Run(&wg)
	wg.Wait()
}
