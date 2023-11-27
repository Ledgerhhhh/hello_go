package main

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/mytask/initconfig"
	"com.ledger.goproject/mytask/nsqdutil"
	"fmt"
)

func init() {
	err := myconfig.InitGConfig()
	err = initconfig.SetupConfig()
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	err := nsqdutil.ProducerUtil(myconfig.GConfig.NsqdConfig.Topic, "hello")
	if err != nil {
		fmt.Println(err)
	}
}
