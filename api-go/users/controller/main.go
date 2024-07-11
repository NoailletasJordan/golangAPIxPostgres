package controller

import (
	"api/users/service"
	"net/http"
)

type Controller struct {
	Service *service.Service
}

type rules map[string]string

func (controller *Controller) GetById(w http.ResponseWriter, r *http.Request) Response {
	idString := r.PathValue("id")
	validate := getValidator()
	err := validate.Var(idString, "required,isUUID")

	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	out, err := controller.Service.GetById(idString)

	return Response{Data: out, Error: err}
}

func (controller *Controller) GetAll(w http.ResponseWriter, r *http.Request) Response {
	out, err := controller.Service.GetAll()
	return Response{Data: out, Error: err}
}

type UpdateByIdBody = struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (controller *Controller) UpdateById(w http.ResponseWriter, r *http.Request) Response {
	w.Header().Set("Content-Type", "application/json")
	idString := r.PathValue("id")
	body, err := GetJSONBody(r.Body)
	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}
	var ruleMap = rules{
		"email": `omitnil,required,type-string,email`,
		"name":  `omitnil,required,type-string,max=40`,
	}

	validate := getValidator()
	err = validateMapCustom(validate, body, ruleMap)

	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	out, err := controller.Service.UpdateById(idString, body)
	return Response{Data: out, Error: err}
}

func (controller *Controller) CreateOne(w http.ResponseWriter, r *http.Request) Response {
	body, err := GetJSONBody(r.Body)
	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	var ruleMap = rules{
		"email": `required,type-string,email`,
		"name":  `required,type-string,max=40`,
		"pass":  `required,type-string,min=6,max=40`,
	}
	validate := getValidator()
	err = validateMapCustom(validate, body, ruleMap)
	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	user, err := controller.Service.CreateOne(body)
	if err != nil {
		return Response{Error: err, Code: http.StatusInternalServerError}
	}

	return Response{Data: user, Code: http.StatusCreated}
}

func (controller *Controller) GetByEmail(w http.ResponseWriter, r *http.Request) Response {
	validation := getValidator()
	email := r.PathValue("email")

	err := validation.Var(email, "required,email")
	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	user, err := controller.Service.GetByEmail(email)
	if err != nil {
		return Response{Error: err, Code: http.StatusInternalServerError}
	}

	return Response{Data: user}
}

func (controller *Controller) ResetPassword(w http.ResponseWriter, r *http.Request) Response {
	idString := r.PathValue("id")
	body, err := GetJSONBody(r.Body)
	if err != nil {
		return Response{Error: err, Code: http.StatusInternalServerError}
	}
	validate := getValidator()
	err = validate.Var(idString, "required,isUUID")
	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	ruleMap := rules{
		"pass": "required,min=6,max=60",
	}

	err = validateMapCustom(validate, body, ruleMap)
	if err != nil {
		return Response{Error: err, Code: http.StatusBadRequest}
	}

	user, err := controller.Service.UpdateById(idString, body)
	if err != nil {
		return Response{Error: err, Code: http.StatusInternalServerError}
	}

	return Response{Data: user, Error: err}
}
