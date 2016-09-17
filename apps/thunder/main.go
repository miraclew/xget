package main

import (
    "os"
    "fmt"
    "flag"
    "strings"
    "encoding/base64"
    "net/url"
)

func main() {
    flag.Usage = func() {
        fmt.Fprintf(os.Stderr, "Usage: %s [option]... [Thunder URL]...\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "Options:\n");
        flag.PrintDefaults()
    }

    flag.Parse()

    if len(os.Args) < 2 {
        flag.Usage()
        return
    }

    thunderUrl := os.Args[len(os.Args)-1]
    thunderUrl = strings.TrimSpace(thunderUrl)
    protocol := "thunder://"

    if strings.Index(thunderUrl, protocol) != 0 {
        fmt.Errorf("URL is not a thunder URL\n")
        return
    }

    thunderUrl = strings.TrimPrefix(thunderUrl, protocol)
    buf, err := base64.StdEncoding.DecodeString(thunderUrl)
    if err != nil {
        fmt.Errorf("base64 decode error: %s\n", err)
        return
    }

    decodeURL := string(buf)
    decodeURL = strings.TrimPrefix(decodeURL, "AA")
    decodeURL = strings.TrimSuffix(decodeURL, "ZZ")

    decodeURL, err = url.QueryUnescape(decodeURL)
    if err != nil {
        fmt.Errorf("URL is not a thunder URL\n")
        return
    }
    fmt.Printf("%s\n", decodeURL)


}


