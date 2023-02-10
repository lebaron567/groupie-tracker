package main

import (
	"fmt"
	"groupieTrackers"
	"log"
	"net/http"
)

func main() {
	homePage, artistPage, locationPage := groupieTrackers.LoadTemplates()

	// Load all assets :
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	infoArtist := groupieTrackers.RecupInfo()

	// Load the first page of the game
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// level = r.FormValue("buttonLevel")
		// homePage.Execute(w, nil)
		homePage.Execute(w, infoArtist)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		locationPage.Execute(w, infoArtist)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		artistPage.Execute(w, infoArtist)
	})

	fmt.Println("Serveur start at : http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
