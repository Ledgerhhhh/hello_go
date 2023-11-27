package test

import (
	"com.ledger.goproject/myconfig"
	"com.ledger.goproject/try_es/client"
	"com.ledger.goproject/try_es/util"
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestEsClient(t *testing.T) {

}

func TestJsonMarshal(t *testing.T) {
	u := user{
		Name: "ledger",
	}
	marshal, err := json.Marshal(u)
	if err != nil {
		return
	}
	t.Logf("%s", marshal)
}

type user struct {
	Name string `json:"name"`
}

func TestJsonUnmarshal(t *testing.T) {
	var u user
	s := "{\"Name\":\"ledger\"}"
	err := json.Unmarshal([]byte(s), &u)
	if err != nil {
		return
	}
	t.Logf("%s", u)
}

func TestJsonEncode(t *testing.T) {
	file, err := os.Open("output.json")
	if err != nil {
		err = fmt.Errorf("open file error: %s\n", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	var data map[string]interface{}
	_ = json.NewDecoder(file).Decode(&data)
	t.Log(data)
}

type ledger struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestJsonDecode(t *testing.T) {
	file, err := os.Create("output.json")
	if err != nil {
		_ = fmt.Errorf("open file error: %s\n", err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	l := ledger{
		Name: "ledger",
		Age:  18,
	}
	err = json.NewEncoder(file).Encode(&l)
	if err != nil {
		return
	}
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func init() {
	err := myconfig.InitGConfig()
	if err != nil {
		return
	}
	err = client.InitEsClient()
	if err != nil {
		return
	}
}

func TestCreatIndex(t *testing.T) {
	_ = util.CreateIndex("test")
}

func TestInsertDoc(t *testing.T) {
	u := User{
		Id:    "123",
		Name:  "ledger",
		Age:   18,
		Email: "123@qq.com",
	}
	err := util.InsertDocument("test", u)
	if err != nil {
		return
	}
}
