package tests

import (
	"api/users/controller"
	"context"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

var validId = uuid.New().String()

func TestGetById(t *testing.T) {
	dbConnexion, testController := setup()
	t.Cleanup(func() {
		teardown(dbConnexion)
	})

	// ADD a user
	query := fmt.Sprintf(`INSERT INTO "public"."users" ( "id", "email", "name", "pass") 
	VALUES ('%v', 'dummy@test.com', 'travis', 'some123pass' );`, validId)
	_, err := dbConnexion.Exec(context.Background(), query)

	if err != nil {
		t.Fatal("Failed to add the valid user")
	}

	for _, values := range fuzzSliceGetById {
		name := values[0]
		expectCode := values[1].(int)
		dummyEntry := values[2].(string)
		testName := fmt.Sprintf("On create user: %v ; expect code: %v ", name, expectCode)
		t.Run(testName, func(t *testing.T) {
			getById(t, testController, dummyEntry, expectCode)
		})
		clearUsersFromDB(dbConnexion)
	}
}

func getById(t *testing.T, testController *controller.Controller, id string, expectCode int) {
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/users/"+id, nil)
	req.SetPathValue("id", id) // because {id} wildcard is not registered
	testController.HandleResponse(testController.GetById)(rr, req)

	handleBodyResponse(t, rr, expectCode)
}

var fuzzSliceGetById = [][3]interface{}{
	{"Valid ID (integer)", 200, validId},
	{"Invalid ID (string)", 400, "abc"},
	{"Negative integer ID", 400, "-5"},
	{"Float ID (not an integer)", 400, "3.14"},
	{"ID exceeding integer range", 400, "999999999999999"},
	{"Empty ID", 400, ""},
	{"Zero ID", 400, "0"},
	{"Boolean ID (not an integer)", 400, "true"},
	{"List of IDs", 400, "1,2,3"},
	{"Nonexistent ID", 404, uuid.New().String()},
	{"Special characters in ID", 400, "@#"},
}
