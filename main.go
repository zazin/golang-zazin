package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang-zazin/internal/delivery/http/user_handler"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	addr := fmt.Sprintf(":%s", port)

	r := mux.NewRouter()
	initHandler(r)
	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println(fmt.Sprintf("server running at : %s", addr))
	log.Fatal(srv.ListenAndServe())
}

func initHandler(router *mux.Router) {
	user_handler.UserHandler(router)
}
