package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"habit/helpers/infra"
	"habit/internal/core/usecases"
	"habit/internal/handler"
	"habit/internal/repositories"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World!")
	port := ":8080"
	router := mux.NewRouter()

	db := infra.ConnectMongo()
	todoRepo := repositories.NewTodoMongoRepo(db)
	todoUsecase := usecases.NewTodoUseCase(todoRepo)

	handler.NewTodoHandler(todoUsecase, router)

	log.Println("Server is starting at " + port)
	log.Fatal(http.ListenAndServe(port, router))

}
