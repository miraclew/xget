package controllers

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/miraclew/xget/apps/xget/web/repositories"
    "github.com/miraclew/xget/apps/xget/web/model"
    "net/http"
)

type TaskController struct {
    repo repositories.TaskRepository
}

func (tc *TaskController) GetList(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: tc.repo.FindAll(),
    })
}

func (tc *TaskController) Add(w rest.ResponseWriter, r *rest.Request) {
    task := &model.Task{}
    r.DecodeJsonPayload(task)
    // TODO: validate

    tc.repo.Add(task)
    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: "",
    })
}

func (tc *TaskController) Delete(w rest.ResponseWriter, r *rest.Request) {
    id := r.PathParam("id")
    if len(id) <= 0 {
        rest.Error(w, "invalid request", http.StatusBadRequest)
    }

    tc.repo.Delete(id)

    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: "",
    })
}
