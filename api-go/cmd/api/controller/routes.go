package controller

import "net/http"

func (controller Controller) InitializeRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{id}", controller.HandleResponse(controller.GetById))
	mux.HandleFunc("GET /users/{$}", controller.HandleResponse(controller.GetAll))
	mux.HandleFunc("PATCH /users/{id}", controller.HandleResponse(controller.UpdateById))
	mux.HandleFunc("POST /users/", controller.HandleResponse(controller.CreateOne))
	mux.HandleFunc("GET /users/email/{email}", controller.HandleResponse(controller.GetByEmail))
	mux.HandleFunc("PUT /users/{id}/password", controller.HandleResponse(controller.ResetPassword))
}
