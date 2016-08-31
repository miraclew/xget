package controllers

import "github.com/ant0ine/go-json-rest/rest"

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type RootController struct {

}

func (sc *RootController) Get(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(&JsonResponse{
		Code: 0,
		Message: "OK",
		Data: "Hello, this is api server. support routes:\n/status",
	})
}

