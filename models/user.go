package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"github.com/astaxie/beego/validation"
)

type User struct {
	gorm.Model
	Email         string `gorm:"size:255" form:"Email" valid:"Required;Email"`
	Password      string `gorm:"size:20" form:"Password" valid:"Required;MinSize(6)"`
	Repassword    string `gorm:"-" form:"Repassword" valid:"Required"` // let gorm ignore this field
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

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "Does not matched password, repassword")
	}
}

func (user *User) Create() (err error) {
	err = db.Create(user).Error
	return err
}

func (user *User) Where(key string, value string) (u User, err error) {
	err = db.Where(key+" = ?", value).Find(&u).Error
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

func Users() (users []User, err error) {
	err = db.Find(users).Error
	return
}

func (user *User) UserExist() (bool, error) {
	var count int = 0
	err := db.Model(&User{}).Where("email = ?", user.Email).Count(&count).Error
	fmt.Println(err, count)
	if err != nil || count != 1 {
		return false, err
	}

	return true, nil
}