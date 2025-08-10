package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Hello World! 当前时间: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("这是一个 GitHub Actions 定时任务脚本")
}
