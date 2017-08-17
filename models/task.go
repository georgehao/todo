package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"samples/todo/models"
	"time"
)

type Task struct {
	gorm.Model
	Content string `gorm:"size:255"`
	Status  int    `gorm:"size:11"`
	Degree  int    `gorm:"size:11"`
}

func (this *Task) All() []Task {
	var tasks []Task
	db.Order("id desc").Find(&tasks)
	return tasks
}

func (this *Task) Create() (uint, error) {
	this.CreatedAt = time.Now()
	err := db.Create(this).Error
	return this.ID, err
}

func (this *Task) Done() (tasks []models.Task) {
	db.Where("status = ?", 1).Find(&tasks)
	return
}

func (this *Task) Page(per int, offset int) ([]Task, error) {
	var tasks []Task
	err := db.Order("id desc").Offset(offset).Limit(per).Find(&tasks).Error
	return tasks, err
}

// MoreImportant get the last 15 import task which isn't complete
func (this *Task) MoreImportant() ([]Task, error) {
	var degree1 []Task
	err := db.Where("status = ? and degree = ?", 0, 1).Find(&degree1).Error
	if err != nil {
		return nil, err
	}

	if len(degree1) >= MOST_IMPORTANT {
		return degree1[:MOST_IMPORTANT], nil
	}

	var last int = MOST_IMPORTANT - len(degree1)
	var degree2 []Task
	err = db.Where("status = ? and degree = ?", 0, 2).Limit(last).Find(&degree2).Error
	if err != nil {
		return nil, err
	}

	last = last - len(degree2)
	if last <= 0 {
		return append(degree1, degree2...), nil
	}

	var degree3 []Task
	err = db.Where("status = ? and degree = ?", 0, 3).Limit(last).Find(&degree3).Error
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