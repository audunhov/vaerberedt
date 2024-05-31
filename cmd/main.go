package main

import (
	"audunhov/vaerberedt/cmd/forecast"
	"audunhov/vaerberedt/cmd/list"
	"audunhov/vaerberedt/cmd/views"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "POST" {
			w.WriteHeader(400)
			w.Write([]byte("Invalid request, expected method POST"))
			return
		}

		newList := list.PackingList{
			Name: r.FormValue("name"),
		}

		views.List(newList).Render(r.Context(), w)
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {

		resp, err := forecast.Load(forecast.CompactRequest{Lat: 50, Lon: 50})

		if err != nil {
			w.WriteHeader(404)
			w.Write([]byte(fmt.Sprintf("Something went wrong: %s", err)))
			return
		}

		views.Test(*resp).Render(r.Context(), w)
	})

	mux.Handle("/list/:id", templ.Handler(views.HomePage()))
	mux.Handle("/", templ.Handler(views.HomePage()))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", mux)
}
