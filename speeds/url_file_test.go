package speeds

import (
    "testing"
    "fmt"
    "errors"
    "net/http"
)

const (
    url1 = "http://mirrors.hust.edu.cn/apache/kafka/0.10.0.0/kafka-0.10.0.0-src.tgz"
    url2 = "http://localhost:3000/public/js/ws.js"
)

func TestGetFileName(t *testing.T) {
    var err1, err2 error
    var name1, name2 string
    name1, err1 = GetFileName(url1)
    name2, err2 = GetFileName(url2)

    if err1 != nil || err2 != nil {
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
