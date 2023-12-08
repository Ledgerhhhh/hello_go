package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

func main() {
	// JSON 数据包含一个用户数组
	jsonData := `
	{
		"users": [
			{"name": "Alice", "age": 25},
			{"name": "Bob", "age": 30},
			{"name": "Charlie", "age": 22}
		],
		"status": "active"
	}
	`

	// 使用 gjson 迭代数组
	result := gjson.Get(jsonData, "users.#.name")

	// 检查是否存在
	if result.Exists() {
		// 获取数组中所有 "name" 的值
		names := result.Array()
		fmt.Printf("User Names:%v", names)

	} else {
		fmt.Println("Path not found in JSON.")
	}
}
