package main

import (
	"fmt"
	"groupieTrackers"
	"log"
	"net/http"
	"strconv"
)

func main() {
	homePage, artistPage, locationPage, concertPage, paysPage := groupieTrackers.LoadTemplates()
	fmt.Println("Serveur start at : http://localhost:8080/home")
	// Load all assets :
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	//récuperer les info de l'"api"
	listGroups := groupieTrackers.RecupInfo()
	newlistGroups := groupieTrackers.DiviserEnListeDeXelement(listGroups, len(listGroups))
	listLocatin := groupieTrackers.SortLieux(listGroups)
	numberPageChoice := 0

	// Load the first page of the game
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		homePage.Execute(w, groupieTrackers.RomdomArtist())
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		locationPage.Execute(w, listLocatin)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		numberOfItemsOnPage := r.FormValue("numberPage")
		pageChoice := r.FormValue("page")
		searchUser := r.FormValue("userSearch")
		sortingChoices := r.FormValue("sortingChoices")

		if pageChoice != "" {
			if pageChoice == "precedent" && numberPageChoice > 0 {
				numberPageChoice--
			} else if pageChoice == "suivant" && numberPageChoice < len(newlistGroups)-1 {
				numberPageChoice++
			}
		} else {
			numberPageChoice = 0
			if searchUser != "" {
				newlistGroups = groupieTrackers.SearchGroupe(searchUser, newlistGroups)
			} else if sortingChoices != "" {
				newlistGroups = groupieTrackers.SortElement(sortingChoices, len(newlistGroups[0]))
			} else {
				if numberOfItemsOnPage != "" {
					number, _ := strconv.Atoi(numberOfItemsOnPage)
					newlistGroups = groupieTrackers.ReconstituerEtDiviserEnListeDeXelement(newlistGroups, number)
				} else {
					newlistGroups = groupieTrackers.ReconstituerEtDiviserEnListeDeXelement(newlistGroups, 52)
				}
			}
		}
		listGroups = newlistGroups[numberPageChoice]
		if searchUser == "" {
			listGroups[0].IsSearch = true
		}
		artistPage.Execute(w, listGroups)
	})

	http.HandleFunc("/concert", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("info")
		idNum, _ := strconv.Atoi(id)
		concertPage.Execute(w, listGroups[idNum])
	})

	http.HandleFunc("/pays", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("info")
		idNum, _ := strconv.Atoi(id)
		paysPage.Execute(w, listLocatin[idNum])
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func RecupInfo() {
	panic("unimplemented")
}
