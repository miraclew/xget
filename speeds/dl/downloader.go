package speeds

type Downloader struct {
    task *Task
}

func NewDownloader(t *Task) *Downloader {
    return &Downloader{task:t}
}

func (d *Downloader) Start() error {
    return nil
}

func (d *Downloader) Stop() error {
    return nil
}

func (d *Downloader) Cancel () error {
    return nil
}



