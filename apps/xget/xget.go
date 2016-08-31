package main

import (
    "flag"
    "fmt"
    "os"
    "net/http"
    "github.com/ant0ine/go-json-rest/rest"
    log "github.com/Sirupsen/logrus"
    "github.com/miraclew/xget/apps/xget/web/controllers"
)

var (
    showVersion      = flag.Bool("version", false, "print version stringx")
    wsIp             = flag.String("ws-ip", "0.0.0.0", "<ip> to listen on for WebSocket clients")
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [option]... [URL]...\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Options:\n");
        flag.PrintDefaults()
    }

    flag.Parse()
}

func runRestApi() {
    root := &controllers.RootController{}

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