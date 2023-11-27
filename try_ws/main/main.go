package main

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_ws/service"
	"fmt"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("initconfig client error: %s\n", err)
	}
}
func main() {
	service.ListenService()
}
