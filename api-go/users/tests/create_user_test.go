package tests

import (
	"api/users/controller"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5"
)

func createUser(t *testing.T, testController *controller.Controller, user userField, expectCode int) {
	jsonData, err := json.Marshal(user)
	if err != nil {
		t.Error("Error Marshal :", err)
	}

	req := httptest.NewRequest("POST", "/users/", bytes.NewBuffer(jsonData))
	rr := httptest.NewRecorder()

	testController.HandleResponse(testController.CreateOne)(rr, req)

	handleBodyResponse(t, rr, expectCode)
}

func teardown(dbConn *pgx.Conn) {
	clearUsersFromDB(dbConn)
	dbConn.Close(context.Background())
}

func TestCreateUser(t *testing.T) {
	dbConnection, testController := setup()
	t.Cleanup(func() {
		teardown(dbConnection)
	})

	for _, values := range fuzzSliceCreateUser {
		name := values[0]
		expectCode := values[1].(int)
		dummyEntry := values[2].(userField)
		testName := fmt.Sprintf("On create user: %v ; expect code: %v ", name, expectCode)
		t.Run(testName, func(t *testing.T) {
			createUser(t, testController, dummyEntry, expectCode)
		})
		clearUsersFromDB(dbConnection)
	}
}

var fuzzSliceCreateUser = [][3]interface{}{
	{"Valid data", 201, userField{
		"email": "test@example.com",
		"name":  "John",
		"pass":  "password123",
	}},
	{"Invalid email format (not a string)", 400, userField{
		"email": 123,
		"name":  "Jane",
		"pass":  "password",
	}},
	{"Empty email", 400, userField{
		"email": "",
		"name":  "Alice",
		"pass":  "abcdef",
	}},
	{"Empty name (not a string)", 400, userField{
		"email": "another@example.com",
		"name":  nil,
		"pass":  "pass123",
	}},
	{"Empty password", 400, userField{
		"email": "john@example.com",
		"name":  "Mary",
		"pass":  "",
	}},
	{"Numeric name (not a string)", 400, userField{
		"email": "test@example.com",
		"name":  123,
		"pass":  "password",
	}},
	{"Valid integer password (string)", 201, userField{
		"email": "test@example.com",
		"name":  "John",
		"pass":  "123456",
	}},
	{"Invalid integer password (less than 6 digits)", 400, userField{
		"email": "test@example.com",
		"name":  "Alice",
		"pass":  12345,
	}},
	{"Invalid integer password (more than 50 digits)", 400, userField{
		"email": "test@example.com",
		"name":  "John",
		"pass":  "12345678901234567890123456789012345678901234567890",
	}},
	{"Boolean email (not a string)", 400, userField{
		"email": true,
		"name":  "John",
		"pass":  "password",
	}},
	{"Negative integer name (not a string)", 400, userField{
		"email": "test@example.com",
		"name":  -123,
		"pass":  "password",
	}},
	{"Slice password (not a string)", 400, userField{
		"email": "test@example.com",
		"name":  "John",
		"pass":  []int{1, 2, 3},
	}},
}
