package tests

import (
	"api/users/controller"
	"api/users/model"
	"api/users/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func setup() (*pgx.Conn, *controller.Controller) {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConnection, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_EXTERNAL_HOST"), os.Getenv("POSTGRES_EXTERNAL_PORT"), os.Getenv("POSTGRES_DB_TEST")))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	clearUsersFromDB(dbConnection)

	newModel := model.Model{DBConnection: dbConnection}
	newService := service.Service{Model: &newModel}
	newController := controller.Controller{Service: &newService}

	return dbConnection, &newController
}

func TestGetAll(t *testing.T) {
	dbConnection, testController := setup()
	t.Cleanup(func() {
		teardown(dbConnection)
	})

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/users/", nil)

	testController.HandleResponse(testController.GetAll)(rr, req)

	handleBodyResponse(t, rr, http.StatusOK)
}
