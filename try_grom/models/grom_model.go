package models

import (
	"gorm.io/gorm"
	"time"
)

// User 表示系统中的用户
type User struct {
	Id        string    `gorm:"unique;not null;primaryKey'" json:"Id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	FullName  string    `json:"full_name"`
	Bio       string    `json:"bio"`
	AvatarURL string    `json:"avatar_url"`
	Posts     []Post    `gorm:"foreignKey:UserID" json:"posts"`
	CreatedAt time.Time `json:"created_at"`
}

func (u User) TableName() string {
	return "users"
}

// Post 表示博客文章
type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `json:"-"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
	Tags    []Tag  `gorm:"many2many:post_tags;" json:"tags"`
}

func (p Post) TableName() string {
	return "post"
}

// Tag 表示文章标签
type Tag struct {
	gorm.Model
	Name  string `gorm:"unique;not null" json:"name"`
	Posts []Post `gorm:"many2many:post_tags;" json:"posts"`
}

func (t Tag) TableName() string {
	return "tag"
}
