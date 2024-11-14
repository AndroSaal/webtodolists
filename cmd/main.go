package main

import (
	todo "ToDoApp"
	"log"

	"../pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", hendlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while run HTTP server: %s", err.Error)
	}
}
