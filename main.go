package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandlerIndex) //routing index
	mux.HandleFunc("/about", handler.HandlerAbout) //routing index
	mux.HandleFunc("/product", handler.ProductHandler) //routing index

	log.Println("Starting Port 8080")

	err := http.ListenAndServe(":8080", mux) //jalankan web server
	log.Fatal(err)
}