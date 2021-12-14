package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"golang-zazin/internal/delivery/http/user_handler"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	initHandler(r)
	addr := "127.0.0.1:8000"
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
