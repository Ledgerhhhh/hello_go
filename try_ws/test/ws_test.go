package test

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_ws/service"
	"fmt"
	"testing"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("initconfig client error: %s\n", err)
	}
}
func TestWs(t *testing.T) {
	service.ListenService()
}
