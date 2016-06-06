package main

import (
    "flag"
    "fmt"
    "os"
    "net/http"
    "io"
    "github.com/miraclew/xget/speeds"
    "errors"
    "time"
    "github.com/gosuri/uiprogress"
    "strconv"
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

    uiprogress.Start()
    for i := 0; i < flag.NArg(); i++ {
        err := dl(flag.Arg(i))
        if err != nil {
            fmt.Println(err.Error())
        }
    }
}

func dl(u string) error {
    fmt.Println("Start dl: ", u)

    res, err := http.DefaultClient.Get(u)
    if err != nil {
        return err
    }

    if res.StatusCode != http.StatusOK {
        return errors.New("dl failed: "+res.Status)
    }

    var contentLength string
    if contentLength = res.Header.Get("Content-Length"); len(contentLength) > 0 {
        fmt.Println("Length: "+ contentLength)
        if contentType := res.Header.Get("Content-Type"); len(contentType) > 0 {
            fmt.Println("Type: "+contentType)
        }
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

    t1 := time.Now();
    totalBytes := 0

    length, err := strconv.ParseInt(contentLength, 10, 0)
    bar := uiprogress.AddBar(int(length))
    bar.AppendCompleted()
    bar.PrependElapsed()

    for ; ;  {
        readBytes, err := res.Body.Read(buf)
        if readBytes > 0 {
            totalBytes += readBytes
            bar.Set(totalBytes)
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

    t2 := time.Now()
    cost := t2.Sub(t1)
    var speed int64
    speed = int64(totalBytes) * time.Second.Nanoseconds()/ cost.Nanoseconds()
    fmt.Printf("Done(%s): Time=%s Speed=%d(B/s)\n", t2.String(), cost.String(), speed)

    return nil
}