package test

import (
	"com.ledger.goproject/myconfig"
	"testing"
)

func TestInitConfig(t *testing.T) {
	err := myconfig.InitGConfig()
	if err != nil {
		panic(err)
	}
}
