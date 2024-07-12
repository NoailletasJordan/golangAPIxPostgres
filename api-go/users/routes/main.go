package routes

import (
	_ "api/docs"
	c "api/users/controller"
	"net/http"
)

type SucessResponse struct {
	Data  any    `json:"data"`
	Error string `json:"error" example:"null"`
	Code  int    `json:"code" enums:"200"`
}

type FailureResponse struct {
	Data  string `json:"data" example:"null"`
	Error string `json:"error" example:"Error Message"`
	Code  int    `json:"code" enums:"400,404,500"`
}

func InitializeRoute(mux *http.ServeMux, controller c.Controller) {
	GetAll(mux, controller)
	GetById(mux, controller)
	CreateOne(mux, controller)
	UpdateById(mux, controller)
	GetByEmail(mux, controller)
	ResetPassword(mux, controller)
}

type ResetPasswordBody = struct {
	Pass string `json:"pass"`
}

// @title Showcase
// @BasePath /

// @Summary		Reset Password
// @Tags			users
// @Produce		json
// @Param   id	path	string	true "User Id"
// @Param   PartialUser	body	ResetPasswordBody	true "body"
// @Success 201 {object} SucessResponse{data=model.User}
// @Failure 400 {object} FailureResponse
// @Failure 404 {object} FailureResponse
// @Failure 500 {object} FailureResponse
// @Router			/users/{id}/password [put]
func ResetPassword(mux *http.ServeMux, controller c.Controller) {
	mux.HandleFunc("PUT /users/{id}/password", controller.HandleResponse(controller.ResetPassword))
}

// @Summary		Get User by ID
// @Tags			users
// @Produce		json
// @Param   email	path	string	true	"User email"
// @Success 200 {object} SucessResponse{data=model.User}
// @Failure 400 {object} FailureResponse
// @Failure 404 {object} FailureResponse
// @Failure 500 {object} FailureResponse
// @Router			/users/email/{email} [get]
func GetByEmail(mux *http.ServeMux, controller c.Controller) {
	mux.HandleFunc("GET /users/email/{email}", controller.HandleResponse(controller.GetByEmail))
}

// temp
type UpdateByIdBody = struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// @Summary		Update by Id
// @Tags			users
// @Produce		json
// @Param   id	path	string	true "User Id"
// @Param   PartialUser	body	UpdateByIdBody	true "body"
// @Success 201 {object} SucessResponse{data=model.User}
// @Failure 400 {object} FailureResponse
// @Failure 404 {object} FailureResponse
// @Failure 500 {object} FailureResponse
// @Router			/users/{id} [patch]
func UpdateById(mux *http.ServeMux, controller c.Controller) {
	mux.HandleFunc("PATCH /users/{id}", controller.HandleResponse(controller.UpdateById))
}

// @Summary		Get All Users
// @Tags			users
// @Produce		json
// @Success 200 {object} SucessResponse{data=[]model.User}
// @Failure 500 {object} FailureResponse
// @Router			/users/ [get]
func GetAll(mux *http.ServeMux, controller c.Controller) {
	mux.HandleFunc("GET /users/{$}", controller.HandleResponse(controller.GetAll))
}

// @Summary		Get One User
// @Tags			users
// @Produce		json
// @Param   id	path	string	true	"User UUID"
// @Success 200 {object} SucessResponse{data=model.User}
// @Failure 400 {object} FailureResponse
// @Failure 404 {object} FailureResponse
// @Failure 500 {object} FailureResponse
// @Router			/users/{id} [get]
func GetById(mux *http.ServeMux, controller c.Controller) {
	mux.HandleFunc("GET /users/{id}", controller.HandleResponse(controller.GetById))
}

type NewUser = struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
}

// @Summary		Create One User
// @Tags			users
// @Produce		json
// @Param   PartialUser	body	NewUser	true "body"
// @Success 201 {object} SucessResponse{data=model.User}
// @Failure 400 {object} FailureResponse
// @Failure 404 {object} FailureResponse
// @Failure 500 {object} FailureResponse
// @Router			/users/ [post]
func CreateOne(mux *http.ServeMux, controller c.Controller) {
	mux.HandleFunc("POST /users/", controller.HandleResponse(controller.CreateOne))
}
