package impl

import (
    "testing"
    "github.com/miraclew/xget/apps/xget/web/model"
    "github.com/jinzhu/gorm"
    "log"
)

func TestTaskRepositoryImpl_Add(t *testing.T) {
    db, err := gorm.Open("sqlite3", "./xget.db")
    if err != nil {
        t.Fatal(err)
    }
    db.AutoMigrate(&model.Task{})

    defer db.Close()

    r := NewTaskRepositoryImpl(db)
    task := &model.Task{
        Name:"TestName",
    }

    err = r.Add(task)
    if err != nil {
        t.Fatal(err)
    }
}

func TestTaskRepositoryImpl_Delete(t *testing.T) {
    db, err := gorm.Open("sqlite3", "./xget.db")
    if err != nil {
        t.Fatal(err)
    }
    err = db.AutoMigrate(&model.Task{}).Error
    if err != nil {
        t.Fatal(err)
    }

    defer db.Close()

    r := NewTaskRepositoryImpl(db)
    task := &model.Task{
        Model: gorm.Model{
            ID: 123,
        },
        Name:"TestName",
    }

    err = r.Delete(task)
    if err != nil {
        t.Fatal(err)
    }
}

func TestTaskRepositoryImpl_FindAll(t *testing.T) {
    db, err := gorm.Open("sqlite3", "./xget.db")
    if err != nil {
        t.Fatal(err)
    }
    err = db.AutoMigrate(&model.Task{}).Error
    if err != nil {
        t.Fatal(err)
    }

    defer db.Close()

    r := NewTaskRepositoryImpl(db)


    tasks, err := r.FindAll()
    if err != nil {
        t.Fatal(err)
    }

    if len(tasks) <= 0 {
        t.Fatal()
    }

    for _, task := range tasks  {
        log.Printf("%d", task.ID)
    }

}