package main

import "fmt"

func main() {
	//c1 := make(chan int,1)
	//c1 <- 1
	//fmt.Println(<-c1)
	c2 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			c2 <- i
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(<-c2)
	}

}
