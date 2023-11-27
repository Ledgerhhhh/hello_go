package main

import (
	"com.ledger.goproject/try_grom/config"
	"com.ledger.goproject/try_grom/model"
)

func init() {
	config.SetUpDB()
}
func main() {
	u := &model.User{}
	err := config.DbL.AutoMigrate(u)
	if err != nil {
		return
	}
}
