package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type handlerMayErr func(w http.ResponseWriter, r *http.Request) Response

type Response struct {
	Data  any   `json:"data"`
	Error error `json:"error"`
	Code  int   `json:"code"`
}

func (controller *Controller) HandleResponse(handlerMayError handlerMayErr) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				statusCode := http.StatusNotExtended
				value, _ := json.Marshal(map[string]any{
					"error": "Internal server panic",
					"code":  statusCode,
					"data":  nil,
				})
				w.WriteHeader(statusCode)
				w.Write(value)
			}
		}()

		response := handlerMayError(w, r)
		w.Header().Set("Content-Type", "application/json")

		statusCode := response.Code
		if response.Code == 0 {
			statusCode = http.StatusOK
		}

		// Error returned by the controller's Method
		if response.Error != nil {

			// Entity not found (404)
			if response.Error == pgx.ErrNoRows {
				statusCode = http.StatusNotFound
				value, _ := json.Marshal(map[string]any{
					"error": "Entity not found",
					"code":  statusCode,
					"data":  nil,
				})
				w.WriteHeader(statusCode)
				w.Write(value)
				return
			}

			// Other errors
			if statusCode < 300 {
				statusCode = 500
			}
			value, _ := json.Marshal(map[string]any{
				"error": response.Error.Error(),
				"code":  statusCode,
				"data":  nil,
			})
			w.WriteHeader(statusCode)
			w.Write(value)
			return
		}

		responseJson, marshalError := json.Marshal(map[string]any{
			"error": nil,
			"code":  statusCode,
			"data":  response.Data,
		})

		// Error marshalling the response
		if marshalError != nil {
			fmt.Println("Error :  ", marshalError.Error())
			statusCode = http.StatusInternalServerError
			errorMarshal, _ := json.Marshal(map[string]any{
				"error": marshalError.Error(),
				"code":  statusCode,
				"data":  nil,
			})
			w.WriteHeader(statusCode)
			w.Write(errorMarshal)
			return
		}

		// Send OK
		w.WriteHeader(statusCode)
		w.Write(responseJson)
	}
}

func GetJSONBody(streamBody io.ReadCloser) (map[string]any, error) {
	defer streamBody.Close()
	var body map[string]any
	err := json.NewDecoder(streamBody).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
