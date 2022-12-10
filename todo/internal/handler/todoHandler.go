package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"habit/internal/core/ports"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	todoUseCase ports.TodoUseCase
}

func NewTodoHandler(todoUseCase ports.TodoUseCase, router *mux.Router) *TodoHandler {
	handler := &TodoHandler{todoUseCase: todoUseCase}
	//TODO maybe add todo handler and call below lines from there with additional configurations
	router.HandleFunc("/health", handler.HealthCheckHandler).Methods("GET")
	router.HandleFunc("/todos", handler.GetAll).Methods("GET")
	router.HandleFunc("/todos/{id}", handler.Get).Methods("GET")
	//router.HandleFunc("/todos", handler.Create).Methods("POST")
	return handler
}

func (t *TodoHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Hi there! I'm OK")
}
func (t *TodoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := t.todoUseCase.GetAll()
	fmt.Println("in here")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	jsonResponse, jsonError := json.Marshal(todos)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (t *TodoHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo, err := t.todoUseCase.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	jsonResponse, jsonError := json.Marshal(todo)

	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
