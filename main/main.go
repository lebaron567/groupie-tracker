package main

import (
	"fmt"
	"groupieTrackers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	homePage, artistPage, locationPage, concertPage := groupieTrackers.LoadTemplates()
	fmt.Println("Serveur start at : http://localhost:8080/")
	// Load all assets :
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	//récuperer les info de l'"api"
	listGroups := groupieTrackers.RecupInfo()

	// Load the first page of the game
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homePage.Execute(w, listGroups)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		locationPage.Execute(w, listGroups)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		searchUser := r.FormValue("userSearch")
		// id := r.FormValue("id")
		// fmt.Println(id)
		if searchUser != "" {
			listGroups = groupieTrackers.SearchGroupe(searchUser, listGroups)
		}

		artistPage.Execute(w, listGroups)
	})

	http.HandleFunc("/concert", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("info")
		idNum ,_ :=strconv.Atoi(id)
		fmt.Println(listGroups[idNum])
		concertPage.Execute(w, listGroups[idNum])
	})


	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
