package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行参数
	methodPtr := flag.String("method", "defaultMethod", "指定要触发的方法")
	// 添加其他命令行参数...

	// 解析命令行参数
	flag.Parse()
	
	// 调用相应的方法
	switch *methodPtr {
	case "method1":
		method1()
	case "method2":
		method2()
	default:
		fmt.Println("未知的方法:", *methodPtr)
	}
}

func method1() {
	fmt.Println("执行方法1")
	// 在这里添加方法1的逻辑
}

func method2() {
	fmt.Println("执行方法2")
	// 在这里添加方法2的逻辑
}
