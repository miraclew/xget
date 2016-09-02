package impl

import (
    "testing"
    "github.com/jinzhu/gorm"
    "log"
    "github.com/miraclew/xget/se/models"
)

func TestTaskRepositoryImpl_Add(t *testing.T) {
    db, err := gorm.Open("sqlite3", "./xget.db")
    if err != nil {
        t.Fatal(err)
    }
    db.AutoMigrate(&models.Task{})

    defer db.Close()

    r := NewTaskRepositoryImpl(db)
    task := &models.Task{
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
    err = db.AutoMigrate(&models.Task{}).Error
    if err != nil {
        t.Fatal(err)
    }

    defer db.Close()

    r := NewTaskRepositoryImpl(db)
    task := &models.Task{
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
    err = db.AutoMigrate(&models.Task{}).Error
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