package models

import (
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
)

/*
	Get form error
*/
func IsValid(model interface{}) (err error) {
	valid := validation.Validation{}
	b, err := valid.Valid(model)
	if !b {
		for _, err := range valid.Errors {
			beego.Warning(err.Key, ":", err.Message)
			return errors.New(err.Message)
			// return errors.New(fmt.Sprintf("%s: %s", err.Key, err.Message))
		}
	}
	return nil
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@/todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
