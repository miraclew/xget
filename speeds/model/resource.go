package model

// Resource represents an file on internet.
type Resource struct {
    Id string
    Name string
    Size int64

    Checksum []byte
}
