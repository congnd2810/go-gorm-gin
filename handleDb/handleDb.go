package handleDb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"unique;not null"`
	Name      string `gorm:"uniqueIndex"`
	Age       uint8
	Email     string `gorm:"unique;not null"`
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Session struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Token  string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
	User    User `gorm:"foreignKey:UserID"`
}

type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}

func InitDb() *gorm.DB {
	dsn := "host=localhost user=congnd password=congnd dbname=firstPostgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	err = db.AutoMigrate(&User{}, &Post{}, &Session{})

	if err != nil {
		log.Fatal(err)
	}
	return db
}
