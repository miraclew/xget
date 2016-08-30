package api

import "github.com/miraclew/xget/speeds/model"

type ResourceApi interface {
    FindByURL(url string) (*model.Resource, error)
    FindById(id string) (*model.Resource, error)

    GetChunks(resourceId string) ([]*model.Chunk, error)
    GetSources(resourceId string) ([]*model.Chunk, error)

    // QUESTION: do we trust client to add new resource?
}
