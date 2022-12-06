package main

import (
	"go-crud/database"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load env from file
	godotenv.Load(".env")
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	config := database.Config{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		UserName: os.Getenv("MYSQL_USERNAME"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	connectionString := database.GetConnectionString(config)

	err := database.Connect(connectionString)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	database.Connector.AutoMigrate(&User{})

	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)

	userRouter := router.PathPrefix("/user").Subrouter()
	userRouter.Path("").Methods(http.MethodPost).HandlerFunc(CreateUser)
	userRouter.Path("").Methods(http.MethodGet).HandlerFunc(GetAllUsers)
	userRouter.Path("/{id}").Methods(http.MethodPut).HandlerFunc(EditUser)
	userRouter.Path("/{id}").Methods(http.MethodDelete).HandlerFunc(DeleteUser)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe(":8090", router))
}
