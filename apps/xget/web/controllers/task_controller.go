package controllers

import "github.com/ant0ine/go-json-rest/rest"

type TaskController struct {

}

func (tc *TaskController) GetList(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: "Hello, this is api server. support routes:\n/status",
    })
}

func (tc *TaskController) Add(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: "Hello, this is api server. support routes:\n/status",
    })
}

func (tc *TaskController) Delete(w rest.ResponseWriter, r *rest.Request) {
    w.WriteJson(&JsonResponse{
        Code: 0,
        Message: "OK",
        Data: "Hello, this is api server. support routes:\n/status",
    })
}
