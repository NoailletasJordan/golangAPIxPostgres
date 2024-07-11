package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

type userField = map[string]any

func clearUsersFromDB(DBconn *pgx.Conn) {
	_, err := DBconn.Exec(context.Background(), "TRUNCATE TABLE users")
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(-1)
	}
}

type Result struct {
	Data  any `json:"data"`
	Error any `json:"error"`
	Code  int `json:"code"`
}

func GetJsonBodyTest(t *testing.T, jsonBytes []byte) *Result {
	var body Result
	err := json.Unmarshal(jsonBytes, &body)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}
	return &body
}

func handleBodyResponse(t *testing.T, rr *httptest.ResponseRecorder, expectCode int) {
	body := GetJsonBodyTest(t, rr.Body.Bytes())
	if rr.Code != expectCode || body.Code != expectCode {
		t.Errorf(`expected statuscode & body.code %v got :
				"rr.Code" -> %v 
				"body.Code" -> %v 
				"body.Error" -> %v`, expectCode, rr.Code, body.Code, body.Error)
	}
}
