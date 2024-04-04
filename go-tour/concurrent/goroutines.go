package main

import (
	"fmt"
	"time"
)

/*
*

	goroutine 是由 Go 运行时环境管理的轻量级线程
*/
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 开启一个新的 goroutine 执行
	go say("world")
	say("hello")
}
