package main

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_es/client"
	"com.ledger.goproject/try_es/util"
	"fmt"
)

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		_ = fmt.Errorf("service error: %s\n", err)
	}
	err = client.InitEsClient()
	if err != nil {
		_ = fmt.Errorf("service error: %s\n", err)
	}
}

func main() {
	//TestCreateIndex()
	//TestInsertDoc()
	//TestSearchDocuments()
	//TestUpdateDocument()
	//TestDeleteDocument()
	TestBulkAddDocuments()
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func TestInsertDoc() {
	u := User{
		Id:    "adssamkndaokmdsapk",
		Name:  "ledger",
		Age:   18,
		Email: "123@qq.com",
	}
	err := util.InsertDocument("test", u)
	if err != nil {
		fmt.Printf("service error: %s\n", err)
	} else {
		fmt.Println("success")
	}
}
func TestCreateIndex() {
	err := util.CreateIndex("test")
	if err != nil {
		return
	}
}
func TestSearchDocuments() {
	query := "{\n  \"query\": {\n    \"match_all\": {}\n  }\n}\n"
	//var u User
	res, err := util.SearchDocuments("test", query)
	if err != nil {
		fmt.Printf("service error: %s\n", err)
		return
	}
	fmt.Println(res, "==========")
}
func TestUpdateDocument() {
	user := User{
		Id:    "999999",
		Name:  "0000",
		Age:   180,
		Email: "1299993@qq.com",
	}
	err := util.UpdateDocument("test", "0b5bvosBs52i_6ADphrD", user)
	if err != nil {
		fmt.Printf("service error: %s\n", err)
	} else {
		fmt.Println("success")
	}
}

func TestDeleteDocument() {
	err := util.DeleteDocument("test", "SL5SvosBs52i_6ADIRh5")
	if err != nil {
		fmt.Printf("service error: %s\n", err)
	} else {
		fmt.Println("success")
	}

}
func TestBulkAddDocuments() {
	var users []User
	i := append(users, User{
		Id:    "9",
		Name:  "ledger",
		Age:   18,
		Email: "123@qq.com",
	})
	i2 := append(i, User{
		Id:    "90",
		Name:  "ledger",
		Age:   18,
		Email: "123@qq.com",
	})
	var interfaceSlice []interface{}
	for _, u := range i2 {
		interfaceSlice = append(interfaceSlice, u)
	}
	err := util.BulkAddDocument("test", interfaceSlice)
	if err != nil {
		fmt.Printf("service error: %s\n\n", err)
	} else {
		fmt.Println("success")
	}
}
