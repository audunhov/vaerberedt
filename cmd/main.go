package main

import (
	"audunhov/vaerberedt/cmd/forecast"
	"audunhov/vaerberedt/cmd/list"
	"audunhov/vaerberedt/views"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /create", func(w http.ResponseWriter, r *http.Request) {
		newList := list.PackingList{
			Name: r.FormValue("name"),
		}

		views.List(newList).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /test", func(w http.ResponseWriter, r *http.Request) {

		resp, err := forecast.Load(forecast.CompactRequest{Lat: 50, Lon: 50})

		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Something went wrong: %s", err)))
			return
		}

		views.Test(*resp).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		views.EditItem(id).Render(r.Context(), w)
	})

	mux.HandleFunc("POST /edit/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(r.FormValue("name"))
		views.Item(id).Render(r.Context(), w)
	})

	mux.Handle("GET /list/:id", templ.Handler(views.HomePage()))
	mux.Handle("GET /", templ.Handler(views.HomePage()))
	mux.Handle("GET /new", templ.Handler(views.CreatePage()))
	mux.Handle("GET /items", templ.Handler(views.ItemsPage()))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
