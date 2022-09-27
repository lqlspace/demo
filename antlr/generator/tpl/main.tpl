package main

import (
	"log"
	"net/http"

	"lqlspace/demo/antlr/generator/server/handler"
)

func main() {
	src := &http.Server{
		Addr:         ":8888",
		Handler:      handler.InitRouter(),
	}

	log.Println("server is running")
	log.Fatal(src.ListenAndServe())
}