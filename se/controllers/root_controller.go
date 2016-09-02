package controllers

import "github.com/ant0ine/go-json-rest/rest"

type RootController struct {

}

func (sc *RootController) Get(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(&JsonResponse{
		Code: 0,
		Message: "OK",
		Data: "Hello, this is api server. support routes:\n/status",
	})
}

