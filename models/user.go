package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	//"crypto/sha256"
	//"io"
	"fmt"
)

type User struct {
	gorm.Model
	Email         string `gorm:"size:255"`
	Password      string `gorm:"size:20"`
	LastLoginTime time.Time
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}

func (user *User) Create() error {
	//h := sha256.New()
	//io.WriteString(h, user.Password)
	//user.Password = string(h.Sum(nil))
	return db.Create(user).Error
}

func (user *User) Users() (users []User, err error) {
	err = db.Find(users).Error
	return
}

func (user *User) Where(key string, value string) (u User, err error) {
	err = db.Where(key + " = ?", value).Find(&u).Error
	return
}

func (user *User) Find() (User, error) {
	var u User
	err := db.Find(&u, user.ID).Error
	fmt.Println(err)
	return u, err
}

func (user *User) Save() error {
	return db.Save(user).Error
}
