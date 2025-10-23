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

	http.HandleFunc("/game/init/traitement", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/game/init", http.StatusSeeOther)
		}

		name1 := r.FormValue("pseudoJoueur1")
		name2 := r.FormValue("pseudoJoueur2")

		game.TraitementPlayerInit(name1)
		game.TraitementPlayerInit(name2)
	})

	path, _ := os.Getwd()
	fileServer := http.FileServer(http.Dir(path + "/assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.ListenAndServe("localhost:8000", nil)
}
