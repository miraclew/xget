package util

import (
    "testing"
    "fmt"
    "net/http"
)

const (
    url1 = "http://mirrors.hust.edu.cn/apache/kafka/0.10.0.0/kafka-0.10.0.0-src.tgz"
    url2 = "http://localhost:3000/public/js/ws.js?a=b&c=d"
)

func TestGetFileName(t *testing.T) {
    var name1, name2 string
    name1 = GetFileNameFromUrl(url1)
    name2 = GetFileNameFromUrl(url2)

    if name1 != "kafka-0.10.0.0-src.tgz" {
        t.Fail()
    }

    if name2 != "ws.js" {
        t.Fail()
    }

    fmt.Println(name1)
    fmt.Println(name2)
}

func TestHttpHeaders(t *testing.T) {
    res, err := http.DefaultClient.Head(url1)
    if err != nil {
        t.Fatal(err.Error())
    }

    if res.StatusCode != http.StatusOK {
        t.Fatal(res.Status)
    }

    //var contentLength, string
    //if contentLength = res.Header.Get("Content-Length"); len(contentLength) > 0 {
    //    fmt.Println("Length: "+ contentLength)
    //    if contentType := res.Header.Get("Content-Type"); len(contentType) > 0 {
    //        fmt.Println("Type: "+contentType)
    //    }
    //}
}
