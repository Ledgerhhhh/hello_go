package main

import (
	"bufio"
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_lua/conn"
	"context"
	"fmt"
	"os"
)

func init() {
	_ = myconfig.InitGConfig()
	_ = conn.InitRedisConfig()
}

func main() {
	client := conn.RClient
	lua := getLua("hash.lua")
	result, err := client.Eval(context.Background(), lua, []string{"ledger", "ledger1"}, 10, "ledger").Result()
	if err != nil {
		fmt.Println("Error executing Lua script:", err)
		return
	}
	fmt.Println("result", result)
}

func redisApi() {
	set := conn.RClient.Set(context.Background(), "ledger", 5, 0)
	fmt.Println(set.Val())
}

func getLua(luaName string) string {
	file, err := os.OpenFile("try_lua/lua_script/"+luaName, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	reader := bufio.NewReader(file)
	var bytes = make([]byte, 1024)
	length, err := reader.Read(bytes)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	lua := string(bytes[0:length])
	return lua
}
