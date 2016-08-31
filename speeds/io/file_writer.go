package io

import (
    "os"
    //"fmt"
    "fmt"
)

func writeChunk(file *os.File, offset int64, b []byte) error {
    ret, err := file.Seek(offset, 0)
    if err != nil {
        return err
    }

    fmt.Printf("offset: %d\n" ,ret)
    fmt.Println("write ", b)
    for i := 0; i < 10; i++ {
        _,err := file.Write(b)
        if err != nil {
            return err
        }
    }

    return nil
}