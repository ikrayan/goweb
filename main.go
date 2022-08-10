package main

import (
	"goweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.HandlerIndex) //routing index
	
	mux.HandleFunc("/form", handler.FormHandler)
	mux.HandleFunc("/processform", handler.ProcessForm)

	mux.HandleFunc("/product", handler.ProductHandler)
	mux.HandleFunc("/formProduct", handler.FormProductHandler)
	mux.HandleFunc("/processFormProduct", handler.ProcessFormProductHandler)

	mux.HandleFunc("/member", handler.MemberHandler)
	mux.HandleFunc("/formMember", handler.FormMemberHandler)
	mux.HandleFunc("/processFormMember", handler.ProcessFormMemberHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting Port 8080")

	err := http.ListenAndServe(":8080", mux) //jalankan web server
	log.Fatal(err)
}