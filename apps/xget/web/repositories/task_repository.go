package repositories

import (
    "github.com/miraclew/xget/apps/xget/web/model"
)

type TaskRepository interface {
    FindAll() ([]model.Task, error)
    Add(t *model.Task) error
    Delete(id uint) error
}