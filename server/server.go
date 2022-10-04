package server

import (
	"L0/cache"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()
	router.HandleFunc("/orders", orderHandler)
	router.HandleFunc("/index", indexHandler)
	http.Handle("/", router)

	http.ListenAndServe(":8181", nil)
}

func orderHandler(writer http.ResponseWriter, request *http.Request) {

	id := request.URL.Query().Get("id")
	order, ok := cache.GetByID(id)
	if !ok {
		http.NotFound(writer, request)
		return
	}

	t, err := template.ParseFiles("./static/orders.html")
	if err != nil {
		http.Error(writer, "Internal Server Error!", 500)
		return
	}
	err = t.Execute(writer, order)
	if err != nil {
		http.Error(writer, "Internal Server Error", 500)
		return
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./static/index.html")
	if err != nil {
		http.Error(writer, "Internal Server Error!", 500)
		return
	}
	err = t.Execute(writer, "")
	if err != nil {
		http.Error(writer, "Internal Server Error", 500)
		return
	}
}
