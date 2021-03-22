package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"restapi/internal/handler"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/posts", handler.AddPost).Methods("POST")

	r.HandleFunc("/posts", handler.GetAllPosts).Methods("GET")

	r.HandleFunc("/post/{id}", handler.GetPost).Methods("GET")

	r.HandleFunc("/posts/{id}", handler.UpdatePost).Methods("PUT")

	r.HandleFunc("/posts/{id}", handler.PatchPost).Methods("PATCH")

	r.HandleFunc("/posts/{id}", handler.DeletePost).Methods("DELETE")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()
		err := srv.Shutdown(ctx)
		done <- err
	}()

	log.Printf("starting server at %s", srv.Addr)
	_ = srv.ListenAndServe()

	err := <-done
	log.Printf("shutting server down with %v", err)
}
