package speeds

import "time"

const (
    CHUNK_SIZE = 4096
)

type DLTask struct {
    OrigUrl string
    FileHash string
    FileName string
    StartAt time.Time
    CompleteAt time.Time
    AcceptRange bool
    Length int64
}

type DLChunk struct {
    Offset int64
    CompleteBytes int64
}

