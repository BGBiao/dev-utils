package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)

	// 发送一个中断的信号
	// signal.Notify(c, os.Interrupt)

	// 接受全部的信号
	signal.Notify(c)

	// 接收到信号并进行处理
	s := <-c
	fmt.Println(s)
}
