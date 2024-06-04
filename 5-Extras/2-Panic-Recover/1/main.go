package main

import (
	"fmt"
	"log"
	"net/http"
)

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("started handler")
		defer func() {
			if r := recover(); r != nil {
				log.Printf("recovered panic: %v\n", r)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})

	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("panic")
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Couldn't listen on %s: %v\n", ":8080", err)
	}
}
