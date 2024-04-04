package main

import "fmt"

func fibonacciForSelect(c, quit chan int) {
	x, y := 0, 1
	// select 语句使得一个 goroutine 在多个通讯操作上等待
	// select 会阻塞，直到条件分支中的某个可以继续执行
	// 当多个都准备好的时候，会随机选择一个
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciForSelect(c, quit)
}
