package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{} {
		"title": "Learning Golang",
		"content": "Learning Golang with Agung Setiawan",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
		return
	}

	// message := "Hello World"
	// w.Write([]byte("Home"))
}

func HandlerAbout(w http.ResponseWriter, r *http.Request) {
	message := "About Page"
	w.Write([]byte(message))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	// message := "Product Page"
	// w.Write([]byte(message))

	fmt.Fprintf(w, "Product Page : %d", idNumb)
}