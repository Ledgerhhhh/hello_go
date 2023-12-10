package main

import (
	"com.ledger.goproject/try_grom/config"
	"com.ledger.goproject/try_grom/models"
	"fmt"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func init() {
	config.SetUpDB()
	db = config.DbL

}
func main() {
	//GetFirst()
	//GetUsers()
	//userWhere()
	//useRaw()
	//useSelect()
	//useOrder()
	//useCount()
	//The()
	//create()
	//save()
	//update()
	delete()
}

func delete() {
	user := models.User{
		Id: "11",
	}
	db.Delete(&user)
}
func update() {
	db.Model(&models.User{}).Where("id=?", 11).Update("username", "9999")
}
func save() {
	user := models.User{
		Id:        "11",
		Username:  "222",
		Email:     "222",
		Password:  "222",
		FullName:  "222",
		Bio:       "222",
		AvatarURL: "222",
		Posts:     nil,
		//CreatedAt: time.Now(),
	}
	db.Save(&user)
}
func create() {
	user := models.User{
		Username:  "111",
		Email:     "111",
		Password:  "111",
		FullName:  "111",
		Bio:       "111",
		AvatarURL: "111",
		Posts:     nil,
		CreatedAt: time.Now(),
	}
	db.Create(&user)

}
func The() {
	var users []models.User
	db.Raw("SELECT *\nFROM users u\nLEFT JOIN posts p\nON u.id=p.user_id\n WHERE u.id=1\n").
		Preload("Posts").
		Scan(&users)
	for _, v := range users {
		fmt.Printf("%+v", v)
	}
}
func useCount() {
	var count int64
	db.Model(&models.User{}).Count(&count)
	fmt.Println(count)
}
func useOrder() {
	var users []models.User
	db.Select("username, email").
		Order("created_at desc").
		Offset(1).
		Limit(1).
		Find(&users)
	for _, v := range users {
		fmt.Printf("%+v", v)
		fmt.Println()
	}
}
func useSelect() {
	var users []models.User
	db.Select("username, email").Limit(1).Find(&users)
	for _, v := range users {
		fmt.Printf("%+v", v)
		fmt.Println()
	}
}
func useRaw() {
	var users []models.User
	sqlStr := `SELECT *
			FROM users
			WHERE username=?
			AND email=?`
	result := db.
		Raw(sqlStr, "john_doe", "john.doe@example.com").
		Scan(&users)
	if result.Error != nil {
		fmt.Println("查询出错:", result.Error)
		return
	}
	for _, v := range users {
		fmt.Printf("%+v", v)
	}
}

func userWhere() {
	var user models.User
	result := db.Where("username = ? AND email = ?", "john_doe", "john.doe@example.com").First(&user)
	if result.Error != nil {
		fmt.Println("查询出错:", result.Error)
		return
	}
	fmt.Printf("%+v", user)
}

func GetUsers() {
	var users []models.User
	result := config.DbL.Find(&users)
	if result.Error != nil {
		fmt.Println("查询出错:", result.Error)
		return
	}
	for _, v := range users {
		fmt.Printf("%+v", v)
		fmt.Println()
	}
}

func GetFirst() {
	var user models.User
	result := config.DbL.First(&user, 1) // 查询ID为1的用户
	if result.Error != nil {
		fmt.Println("查询出错:", result.Error)
		return
	}
	fmt.Printf("%+v\n", user)
}
