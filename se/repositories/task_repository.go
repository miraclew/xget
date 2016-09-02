package repositories

import (
    "github.com/miraclew/xget/se/models"
)

type TaskRepository interface {
    GetAll() ([]models.Task, error)
    Add(t *models.Task) error
    Delete(id uint) error
}