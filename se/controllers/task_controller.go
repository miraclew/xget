package controllers

import (
    "github.com/ant0ine/go-json-rest/rest"
    "net/http"
    "github.com/miraclew/xget/speeds/util"
    "strconv"
    "github.com/miraclew/xget/se/repositories"
    "github.com/miraclew/xget/se/models"
)

type AddTaskReq struct {
    URL string
    BtFile string
}

type TaskController struct {
    Repo repositories.TaskRepository
}

func (tc *TaskController) GetAll(w rest.ResponseWriter, r *rest.Request) {
    tasks, err := tc.Repo.GetAll()
    if err != nil {
        rest.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: tasks,
    })
}

func (tc *TaskController) Add(w rest.ResponseWriter, r *rest.Request) {
    req := &AddTaskReq{}
    r.DecodeJsonPayload(req)
    if req.URL == "" {
        rest.Error(w, "url is required", http.StatusBadRequest)
        return
    }

    // TODO: check if there's any task with same orignUrl exists

    fileName := util.GetFileNameFromUrl(req.URL)
    task := &models.Task{
        Name:fileName,
        OrigUrl:req.URL,
    }

    tc.Repo.Add(task)
    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: task,
    })
}

func (tc *TaskController) Delete(w rest.ResponseWriter, r *rest.Request) {
    id, err := strconv.ParseUint(r.PathParam("id"), 10, 64)
    if err != nil {
        rest.Error(w, "invalid request", http.StatusBadRequest)
        return
    }

    tc.Repo.Delete(uint(id))

    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: "",
    })
}
