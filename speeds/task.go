package speeds

import "time"

const (
    CHUNK_SIZE = 4096

    TS_CREATE = 1
    TS_READY = 2
    TS_DOWNLOADING = 3
    TS_FIN_SUCCESS = 4
    TS_FIN_ERROR = 5
    TS_FIN_CANCEL = 6
)

func CreateTask(urlPath string) (error, *Task) {
    return nil, &Task{OrigUrl:urlPath, AcceptRange:false, Length:0, Status:TS_CREATE}
}

type Task struct {
    OrigUrl string
    FileHash string
    FileName string
    StartAt time.Time
    CompleteAt time.Time
    AcceptRange bool
    Length int64
    Status int
}

// Divide task to task items
type TaskItem struct {
    // index of this chunk
    Index int64
    Offset int64 // start position on whole file
    Size int64
    CompleteBytes int64 //
    Checksum []byte
}

