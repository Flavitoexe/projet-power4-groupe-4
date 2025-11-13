package main

import (
	"fmt"
	"game/game"
	"html/template"
	"net/http"
	"os"
)

func main() {

	listTemplates, errTemplate := template.ParseGlob("./templates/*.html")
	if errTemplate != nil {
		fmt.Println(errTemplate.Error())
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "menu", nil)
	})

	http.HandleFunc("/game/init", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "game-init", nil)
	})

	http.HandleFunc("/game/play", func(w http.ResponseWriter, r *http.Request) {
		grille := game.InitGrille()
		listTemplates.ExecuteTemplate(w, "game-play", grille)
	})

	path, _ := os.Getwd()
	fileServer := http.FileServer(http.Dir(path + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8000", nil)
}
