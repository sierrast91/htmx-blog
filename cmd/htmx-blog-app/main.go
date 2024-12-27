package main

import (
	"html/template"
	"log"
	"net/http"

	"htmx-blog-app/internal/route"
)

func main() {

	// template init
	template := template.Must(template.ParseGlob("templates/*/*.html"))
	// router init
	router := route.NewRouter(nil, template)
	// server init
	listenAddr := ":8080"
	log.Println("server stattin on port ", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		log.Fatal(err)
	}
}
