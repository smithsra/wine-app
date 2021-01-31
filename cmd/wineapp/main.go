package main

import (
	"log"
	"net/http"

	"github.com/smithsra/wine-app/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Handler)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/matchwine", handlers.MatchWine)
	http.HandleFunc("/matchwine/process", handlers.MatchWineProcess)
	http.HandleFunc("/wine.jpg", handlers.WinePic)
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
