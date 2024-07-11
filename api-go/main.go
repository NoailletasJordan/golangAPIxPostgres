package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	userC "api/users/controller"
	userM "api/users/model"
	userR "api/users/routes"
	userS "api/users/service"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

const port = "4020"

func main() {
	err := godotenv.Load("../.env")
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
	userModel := userM.Model{DBConnection: dbConnection}
	userService := userS.Service{Model: &userModel}
	userController := userC.Controller{Service: &userService}
	userR.InitializeRoute(mux, userController)

	mux.HandleFunc("GET /swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4020/swagger/doc.json"),
	))

	fmt.Println("listenin on port", port)
	err = http.ListenAndServe(":"+port, mux)
	if err != nil {
		fmt.Println("Error :", err)
	}
}
