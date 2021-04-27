package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	// 使用带有context 的信号通知
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)

	defer stop()

	// os.Getpid()
	p, err := os.FindProcess(os.Getpid())

	// 给进程p 发个中断信号
	if err := p.Signal(os.Interrupt); err != nil {
		log.Fatal(err)
	}

	select {
	case <-time.After(time.Second):
		fmt.Println("missed signal")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
		stop()
	}
}
