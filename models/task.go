package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
	"samples/todo/models"
)

type Task struct {
	gorm.Model
	Content string `gorm:"size:255"`
	Status  int    `gorm:"size:11"`
	Degree  int    `gorm:"size:11"`
}

func (this *Task) All() []Task {
	var tasks []Task
	db.Find(&tasks)
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
