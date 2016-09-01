package main

import (
    "flag"
    "fmt"
    "os"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"
    log "github.com/Sirupsen/logrus"
    "github.com/miraclew/xget/apps/xget/web/controllers"
    "github.com/jinzhu/gorm"
    "github.com/miraclew/xget/apps/xget/web/model"
    "github.com/miraclew/xget/apps/xget/web/repositories"
    "github.com/miraclew/xget/apps/xget/web/repositories/impl"
)

var (
    showVersion      = flag.Bool("version", false, "print version stringx")

    db *gorm.DB
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [option]... [URL]...\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Options:\n");
        flag.PrintDefaults()
    }

    flag.Parse()
    var err error
    db, err = gorm.Open("sqlite3", "./xget.db")
    if err != nil {
        log.Fatal(err)
    }

    db.AutoMigrate(&model.Task{})
    defer db.Close()

    runRestApi();
}

func runRestApi() {
    root := &controllers.RootController{}
    var repo repositories.TaskRepository
    repo = impl.NewTaskRepositoryImpl(db)
    task := &controllers.TaskController{
        Repo: repo,
    }

    middlewares := []rest.Middleware{
        &rest.TimerMiddleware{},
        &rest.RecorderMiddleware{},
        &rest.PoweredByMiddleware{},
        &rest.RecoverMiddleware{
            EnableResponseStackTrace: true,
        },
        &rest.JsonIndentMiddleware{},
        &rest.ContentTypeCheckerMiddleware{},
    }

    api := rest.NewApi()
    api.Use(middlewares...)
    router, err := rest.MakeRouter(
        rest.Get("/", root.Get),
        rest.Get("/tasks", task.GetAll),
        rest.Post("/tasks", task.Add),
        rest.Delete("/tasks/:id", task.Delete),
    )

    if err != nil {
        log.Fatal(err)
    }

    api.SetApp(router)

    http.Handle("/api/", http.StripPrefix("/api", api.MakeHandler()))
    http.Handle("/web/", http.StripPrefix("/web", http.FileServer(http.Dir("./web/public"))))

    log.Info("ListenAndServe :8080")
    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
        os.Exit(-1)
    }
}