package main

import (
	"app/router"
	"log"
	"net/http"
	"time"
)

func main() {
	server := http.Server{
		Addr:         ":8080",
		Handler:      router.Router(),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 30,
	}

	log.Fatalln(server.ListenAndServe())
}
