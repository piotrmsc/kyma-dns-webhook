package internal

import (
	"fmt"
	"net/http"
)

func PresentHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("present handler has been called")
	})
}

func CleanupHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("cleanup handler has been called")
	})
}
