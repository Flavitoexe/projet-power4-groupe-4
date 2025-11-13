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
		err := r.URL.Query().Get("error")
		listTemplates.ExecuteTemplate(w, "game-init", map[string]string{
			"Error": err,
		})
	})

	http.HandleFunc("/game/init/traitement", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/game/init", http.StatusSeeOther)
		}

		name1 := r.FormValue("pseudoJoueur1")
		name2 := r.FormValue("pseudoJoueur2")

		color1 := r.FormValue("couleurJoueur1")
		color2 := r.FormValue("couleurJoueur2")

		err1 := game.TraitementPlayerInit(name1)
		err2 := game.TraitementPlayerInit(name2)

		if err1 != nil || err2 != nil || name1 == name2 {
			http.Redirect(w, r, "/game/init?error=1", http.StatusSeeOther)
		}

		if color1 == color2 {
			http.Redirect(w, r, "/game/init?error=2", http.StatusSeeOther)
		}

		http.Redirect(w, r, "/game/play", http.StatusSeeOther)
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
