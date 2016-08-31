package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model

	//Resource *Resource
	//Sources  []*Source
	Name string

	CompleteAt time.Time
	Status     int
}
