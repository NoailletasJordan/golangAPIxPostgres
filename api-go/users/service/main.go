package service

import (
	"api/users/model"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/iancoleman/strcase"
)

type Service struct {
	Model *model.Model
}

func (service *Service) GetById(id string) (*model.User, error) {
	query := fmt.Sprintf("select * from users where id = '%v';", id)
	out, err := service.Model.GetOne(query)
	return out, err
}

func (service Service) GetAll() ([]*model.User, error) {
	query := "select id, email, name, pass, permission_level, created_at, updated_at from users;"

	out, err := service.Model.GetAll(query)
	return out, err
}

func (service Service) UpdateById(id string, updateFields map[string]any) (*model.User, error) {
	userOld, err := service.GetById(id)
	if err != nil {
		return nil, err
	}

	userNew := userOld
	var querySetters []string // [`"key" = 'value'`,]
	for key, value := range updateFields {
		field := reflect.ValueOf(userNew).Elem().FieldByName(getUpperFirst(strcase.ToCamel(key)))
		if !field.IsValid() {
			return nil, errors.New("field does not exist: " + key)
		}
		if !field.CanSet() {
			return nil, errors.New("unable to set field " + key)
		}
		field.Set(reflect.ValueOf(value))
		querySetters = append(querySetters, fmt.Sprintf("\"%v\" = '%v'", key, value))
	}

	query := fmt.Sprintf(`UPDATE "public"."users" SET %v where id = '%v';`, strings.Join(querySetters, ", "), id)
	err = service.Model.Execute(query)
	if err != nil {
		return nil, err
	}

	return service.GetById(id)
}

func (service Service) CreateOne(fields map[string]any) (*model.User, error) {
	type format struct {
		labels []string
		values []string
	}
	var formattedInput format
	for key, value := range fields {
		stringValue := fmt.Sprintf("'%v'", value)
		formattedInput.labels = append(formattedInput.labels, "\""+key+"\"")
		formattedInput.values = append(formattedInput.values, stringValue)
	}

	uuidString := uuid.New().String()
	query := fmt.Sprintf(`INSERT INTO "public"."users" ( "id", %v ) 
	VALUES ( '%v', %v );`, strings.Join(formattedInput.labels, ", "), uuidString, strings.Join(formattedInput.values, ", "))

	err := service.Model.Execute(query)
	if err != nil {
		return nil, err
	}

	return service.GetById(uuidString)
}

func (service Service) GetByEmail(email string) (*model.User, error) {
	query := fmt.Sprintf("select * from users where email = '%v';", email)
	out, err := service.Model.GetOne(query)
	return out, err
}

func getUpperFirst(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
