package model

const (
    CHUNK_SIZE = 4096
)

// Chunk represent a part of resource
type Chunk struct {
    Index int64
    Offset int64 // start position on whole file
    Size int64
    Checksum []byte
}
