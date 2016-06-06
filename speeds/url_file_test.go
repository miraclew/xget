package speeds

import (
    "testing"
    "fmt"
)

func TestGetFileName(t *testing.T) {

    url1 := "http://mirrors.hust.edu.cn/apache/kafka/0.10.0.0/kafka-0.10.0.0-src.tgz"
    url2 := "http://localhost:3000/public/js/ws.js"

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
