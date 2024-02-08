package main

import (
	"BookApiGo/handler"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	router.Post("/api/v1/books", handler.CreateNewBook)
	router.Get("/api/v1/books/{id}", handler.GetBook)
	router.Get("/api/v1/books", handler.ListAllBooks)
	router.Put("/api/v1/books/{id}", handler.UpdateBook)
	router.Delete("/api/v1/books/{id}", handler.DeleteBook)

	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Errorf("%w", err)
	}
}
