package main

import (
	"fmt"
	"groupieTrackers"
	"log"
	"net/http"
)

func main() {
	homePage, artistPage, locationPage := groupieTrackers.LoadTemplates()
	fmt.Println("Serveur start at : http://localhost:8080/")
	// Load all assets :
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	listGroups := groupieTrackers.RecupInfo()

	// Load the first page of the game
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// level = r.FormValue("buttonLevel")
		// homePage.Execute(w, nil)
		homePage.Execute(w, listGroups)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		locationPage.Execute(w, listGroups)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		artistPage.Execute(w, listGroups)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
