package impl

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "github.com/miraclew/xget/se/models"
)

type TaskRepositoryImpl struct {
    db *gorm.DB
}

func NewTaskRepositoryImpl(db *gorm.DB) (*TaskRepositoryImpl) {
    return &TaskRepositoryImpl{
        db:db,
    }
}

func (r *TaskRepositoryImpl) GetAll() ([]models.Task, error) {
    var tasks []models.Task

    err := r.db.Find(&tasks).Error
    if err != nil {
        return nil, err
    }

    return tasks, nil
}

func (r *TaskRepositoryImpl) Add(task *models.Task) (error) {
    return r.db.Create(task).Error
}

func (r *TaskRepositoryImpl) Delete(id uint) (error) {
    return r.db.Where("id=?", id).Delete(&models.Task{}).Error
}