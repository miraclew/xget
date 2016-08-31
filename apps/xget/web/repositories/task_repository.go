package repositories

import (
    "github.com/miraclew/xget/speeds"
)

type TaskRepository interface {
    FindAll() ([]*speeds.Task, error)
}