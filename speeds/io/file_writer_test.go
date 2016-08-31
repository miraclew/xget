package io

import (
    "os"
    "testing"
)

func TestWriteChunk(t *testing.T) {
    file, err := os.OpenFile("aaa.txt", os.O_RDWR | os.O_CREATE, 0666)
    if err != nil {
        t.Fatal(err.Error())
    }

    defer file.Close()

    _,err = file.Write([]byte{'x'})
    if err != nil {
        t.Fatal(err.Error())
    }

    err = writeChunk(file, 10, []byte{'a'})
    if err != nil {
        t.Fatal(err.Error())
    }

    err = writeChunk(file, 30, []byte{'b'})
    if err != nil {
        t.Fatal(err.Error())
    }
}

