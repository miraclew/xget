package speeds

import "time"

const (
	TS_CREATE      = 1
	TS_READY       = 2
	TS_DOWNLOADING = 3
	TS_FIN_SUCCESS = 4
	TS_FIN_ERROR   = 5
	TS_FIN_CANCEL  = 6
)

type Task struct {
	Resource *Resource
	Sources  []*Source

	CreatedAt  time.Time
	CompleteAt time.Time
	Status     int
}

// Divide task to task items
type TaskItem struct {
    Chunk *Chunk
    Source *Source
    CompleteBytes int64 //
}
