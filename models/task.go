package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Content string `gorm:"size:255"`
	Status  int    `gorm:"size:11"`
	Degree  int    `gorm:"size:11"`
	UserId  uint
}

func (this *Task) All() []Task {
	var tasks []Task
	user := User{Model:gorm.Model{ID: this.UserId}}
	db.Model(&user).Order("id desc").Related(&tasks)
	return tasks
}

func (this *Task) Create() (uint, error) {
	this.CreatedAt = time.Now()
	err := db.Create(this).Error
	return this.ID, err
}

func (this *Task) Done() (tasks []Task) {
	user := User{Model:gorm.Model{ID: this.UserId}}
	db.Model(&user).Where("status = ?", 1).Related(&tasks)
	return
}

func (this *Task) Page(per int, offset int) ([]Task, error) {
	var tasks []Task
	user := User{Model:gorm.Model{ID: this.UserId}}
	err := db.Model(&user).Order("id desc").Offset(offset).Limit(per).Related(&tasks).Error
	return tasks, err
}

// MoreImportant get the last 15 import task which isn't complete
func (this *Task) MoreImportant() ([]Task, error) {
	user := User{Model:gorm.Model{ID: this.UserId}}
	var degree1 []Task
	err := db.Model(&user).Where("status = ? and degree = ?", 0, 1).Related(&degree1).Error
	if err != nil {
		return nil, err
	}

	if len(degree1) >= MOST_IMPORTANT {
		return degree1[:MOST_IMPORTANT], nil
	}

	var last int = MOST_IMPORTANT - len(degree1)
	var degree2 []Task
	err = db.Model(&user).Where("status = ? and degree = ?", 0, 2).Limit(last).Related(&degree2).Error
	if err != nil {
		return nil, err
	}

	last = last - len(degree2)
	if last <= 0 {
		return append(degree1, degree2...), nil
	}

	var degree3 []Task
	err = db.Model(&user).Where("status = ? and degree = ?", 0, 3).Limit(last).Related(&degree3).Error
	if err != nil {
		return nil, err
	}

	return append(append(degree1, degree2...), degree3...), nil
}

func (this *Task) Modify() {
	db.First(this, this.ID)
	if this.Status == 0 {
		db.Model(this).Update("status", 1)
	} else {
		db.Model(this).Update("status", 0)
	}
}
