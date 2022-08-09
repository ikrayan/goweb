package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandlerIndex) //routing index
	mux.HandleFunc("/product", handler.ProductHandler) //routing index
	mux.HandleFunc("/form", handler.FormHandler)
	mux.HandleFunc("/member", handler.MemberHandler) 
	mux.HandleFunc("/processform", handler.ProcessForm)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting Port 8080")

	err := http.ListenAndServe(":8080", mux) //jalankan web server
	log.Fatal(err)
}