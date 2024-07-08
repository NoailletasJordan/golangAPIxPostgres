package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"auth/cmd/api/controller"
	"auth/cmd/api/model"
	"auth/cmd/api/service"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

const port = "4020"

func main() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	dbConnection, err := pgx.Connect(context.Background(),
		fmt.Sprintf("postgres://%v:%v@%v:%v/%v",
			os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_EXTERNAL_HOST"), os.Getenv("POSTGRES_EXTERNAL_PORT"), os.Getenv("POSTGRES_DB")))

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbConnection.Close(context.Background())

	mux := http.NewServeMux()
	newModel := model.Model{DBConnection: dbConnection}
	newService := service.Service{Model: &newModel}
	newController := controller.Controller{Service: &newService}
	newController.InitializeRoute(mux)

	fmt.Println("listenin on port", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("Error :", err)
	}
}
