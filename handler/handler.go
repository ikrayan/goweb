package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	message := "Hello World"
	w.Write([]byte(message))
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