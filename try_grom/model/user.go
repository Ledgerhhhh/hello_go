package model

type User struct {
	Id       string `json:"id"  gorm:"column:id" gorm:"primaryKey"`
	UserName string `json:"username" gorm:"column:username"`
	Email    string `json:"email" gorm:"column:email"`
}

func (User) TableName() string {
	return "user" // 指定数据库表名为 "user"
}
