package handler

import (
	"fmt"
	"goweb/entity"
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

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Tampilin view error happen, Keep calm", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{} {
	// 	"title": "Learning Golang",
	// 	"content": "Learning Golang with Agung Setiawan",
	// }
	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}
	
	// message := "Hello World"
	// w.Write([]byte("Home"))

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Execute data error happen, Keep calm", http.StatusInternalServerError)
		return
	}	
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Query().Get("id")
	// idNumb, err := strconv.Atoi(id)

	// if err != nil || idNumb < 1 {
	// 	http.NotFound(w, r)
	// 	return
	// }

	

	// fmt.Fprintf(w, "Product page : %d", idNumb)
	// data := map[string]interface{} {
	// 	"content": idNumb,
	// }
	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 3}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Parse data error happen, Keep calm", http.StatusInternalServerError)
		return
	}

	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 12, Photo: "/static/img/mobilio.jpg"},
		{ID: 2, Name: "Xpander", Price: 250000000, Stock: 6, Photo: "/static/img/xpander.jpg"},
		{ID: 3, Name: "Stargazer", Price: 280000000, Stock: 1, Photo: "/static/img/stargazer.jpg"},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Execute data error happen, Keep calm", http.StatusInternalServerError)
		return
	}	
}

func FormProductHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "formProduct.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}	

		return
	}

	http.Error(w, "Get error, Take it easy boy", http.StatusBadRequest)
}

func ProcessFormProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Parsing form error, Keep calm", http.StatusInternalServerError)
			return
		}

		idProduct := r.Form.Get("id")
		name := r.Form.Get("name")
		priceProduct := r.Form.Get("price")
		stockProduct := r.Form.Get("stock")
		photo := r.Form.Get("photo")

		id, err := strconv.Atoi(idProduct)
		if err != nil {
			fmt.Println("Post ID error")
			return
			}

		price, err := strconv.Atoi(priceProduct)
		if err != nil {
			fmt.Println("Post ID error")
			return
			}
		
		stock, err := strconv.Atoi(stockProduct)
		if err != nil {
			fmt.Println("Post ID error")
			return
			}

		tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Parse view error happen, Keep calm", http.StatusInternalServerError)
			return
		}

		data := []entity.Product{
			{ID: 1, Name: "Mobilio", Price: 220000000, Stock: 12, Photo: "/static/img/mobilio.jpg"},
			{ID: 2, Name: "Xpander", Price: 250000000, Stock: 6, Photo: "/static/img/xpander.jpg"},
			{ID: 3, Name: "Stargazer", Price: 280000000, Stock: 1, Photo: "/static/img/stargazer.jpg"},
			{ID: id, Name: name, Price: price, Stock: stock, Photo: photo},
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Execute data error happen, Keep calm", http.StatusInternalServerError)
			return
		}	
	}
}

//BATAS HANDLER PRODUCT

func MemberHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles(path.Join("views", "member.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Parse view error happen, Keep calm", http.StatusInternalServerError)
		return
	}

	data := []entity.Member{
		{ID: 1, Name: "M. Ikra Yan Hidayat", Level: "Newbie", Status: "Active", Photo: "https://picsum.photos/id/111/200"},
		{ID: 2, Name: "Deno Norsanto", Level: "Master", Status: "Active", Photo: "https://picsum.photos/id/222/200"},
		{ID: 3, Name: "Yuhazmi Dartius", Level: "Syifu", Status: "Not Active", Photo: "https://picsum.photos/id/99/200"},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Execute data error happen, Keep calm", http.StatusInternalServerError)
		return
	}	
}

func FormMemberHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "formMember.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}	

		return
	}

	http.Error(w, "Get error, Take it easy boy", http.StatusBadRequest)
}

func ProcessFormMemberHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Parsing form error, Keep calm", http.StatusInternalServerError)
			return
		}

		idMember := r.Form.Get("id")
		name := r.Form.Get("name")
		level := r.Form.Get("level")
		status := r.Form.Get("status")
		photo := r.Form.Get("photo")

		id, err := strconv.Atoi(idMember)
		if err != nil {
			fmt.Println("Post ID error")
			return
			}

		tmpl, err := template.ParseFiles(path.Join("views", "member.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Parse view error happen, Keep calm", http.StatusInternalServerError)
			return
		}

		data := []entity.Member{
			{ID: 1, Name: "M. Ikra Yan Hidayat", Level: "Newbie", Status: "Active", Photo: "https://picsum.photos/id/111/200"},
			{ID: 2, Name: "Deno Norsanto", Level: "Master", Status: "Active", Photo: "https://picsum.photos/id/222/200"},
			{ID: 3, Name: "Yuhazmi Dartius", Level: "Syifu", Status: "Not Active", Photo: "https://picsum.photos/id/99/200"},
			{ID: id, Name: name, Level: level, Status: status, Photo: photo},
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Execute data error happen, Keep calm", http.StatusInternalServerError)
			return
		}	
	}
}

func FormHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}	

		return
	}

	http.Error(w, "Get error, Take it easy boy", http.StatusBadRequest)
}

func ProcessForm(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Parsing form error, Keep calm", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		pesan := r.Form.Get("pesan")

		data := map[string]interface{}{
			"name": name,
			"pesan": pesan,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "resultform.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Send result form error, Keep calm", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Shit happen, Keep calm", http.StatusInternalServerError)
			return
		}	

		return
	}

	http.Error(w, "Post error, Take it easy boy", http.StatusBadRequest)
}