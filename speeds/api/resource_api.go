package api

import (
    "github.com/miraclew/xget/speeds"
)

type ResourceApi interface {
    FindByURL(url string) (*speeds.Resource, error)
    FindById(id string) (*speeds.Resource, error)

    GetChunks(resourceId string) ([]*speeds.Chunk, error)
    GetSources(resourceId string) ([]*speeds.Chunk, error)

    // QUESTION: do we trust client to add new resource?
}
