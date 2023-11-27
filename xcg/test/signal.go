package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	// "time"
)

func main() {
	// 创建一个接收信号的通道
	signals := make(chan os.Signal, 1)

	// 使用 signal.Notify() 设置要捕获的信号
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("程序启动")

	// 同步等待信号
	sig := <-signals
	fmt.Printf("接收到信号：%v\n", sig)

	// 异步等待信号
	// go func() {
	// 	sig := <-signals
	// 	fmt.Printf("接收到信号：%v\n", sig)
	// }()

	// 执行一些清理工作或其他必要的操作
	fmt.Println("程序结束")

	// time.Sleep(10 * time.Second)
}
