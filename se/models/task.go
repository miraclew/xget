package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model

	Name string
	OrigUrl string
	//Resource *Resource
	//Sources  []*Source

	CompleteAt time.Time
	Status     int
}
