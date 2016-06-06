package main

import (
    "flag"
    "fmt"
    "os"
    "net/http"
    "io"
    "github.com/miraclew/xget/speeds"
    "errors"
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

    for i := 0; i < flag.NArg(); i++ {
        err := dl(flag.Arg(i))
        if err != nil {
            fmt.Println(err.Error())
        }
    }
}

func dl(u string) error {
    fmt.Println("start dl: ", u)

    res, err := http.DefaultClient.Get(u)
    if err != nil {
        return err
    }

    if res.StatusCode != http.StatusOK {
        return errors.New("dl failed: "+res.Status)
    }

    buf := make([]byte, 1024)
    filePath, err := speeds.GetFileName(u);
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    defer file.Close()

    for ; ;  {
        readBytes, err := res.Body.Read(buf)
        if readBytes > 0 {
            file.Write(buf[:readBytes])
        }

        if err != nil {
            if err != io.EOF {
                return err
            } else {
                //fmt.Println(err.Error())
                break
            }
        }
    }

    return nil
}