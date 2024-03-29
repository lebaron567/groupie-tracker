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
	infoPrinted := groupieTrackers.Init()
	listLocation := groupieTrackers.SortLieux(infoPrinted.ArtistList)

	// Load the first page of the game
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		homePage.Execute(w, groupieTrackers.RomdomArtist())
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		infoPrinted.ArtistList = groupieTrackers.RecupAllRelation(infoPrinted.ArtistList)
		infoPrinted.PaginatedArtistList = groupieTrackers.DiviserEnListeDeXelement(infoPrinted.ArtistList, len(infoPrinted.PaginatedArtistList[0]))
		listLocation = groupieTrackers.SortLieux(infoPrinted.ArtistList)
		locationPage.Execute(w, listLocation)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		numberOfItemsOnPage := r.FormValue("numberPage")
		pageChoice := r.FormValue("page")
		searchUser := r.FormValue("userSearch")
		sortingChoices := r.FormValue("sortingChoices")
		reset := r.FormValue("Reinitialiser")
		if reset != "" {
			infoPrinted = groupieTrackers.Init()
			listLocation = groupieTrackers.SortLieux(infoPrinted.ArtistList)
		}
		if pageChoice != "" {
			if pageChoice == "precedent" && infoPrinted.IndexCurrentPage > 0 {
				infoPrinted.IndexCurrentPage--
			} else if pageChoice == "suivant" && infoPrinted.IndexCurrentPage < len(infoPrinted.PaginatedArtistList)-1 {
				infoPrinted.IndexCurrentPage++
			}
		} else {
			infoPrinted.IndexCurrentPage = 0
			if searchUser != "" {
				infoPrinted = groupieTrackers.SearchGroupe(searchUser, infoPrinted)
			}
			if sortingChoices != "" {
				infoPrinted.PaginatedArtistList = groupieTrackers.SortElement(sortingChoices, len(infoPrinted.PaginatedArtistList[0]))
			}
			if numberOfItemsOnPage != "" {
				number, _ := strconv.Atoi(numberOfItemsOnPage)
				infoPrinted.PaginatedArtistList = groupieTrackers.ReconstituerEtDiviserEnListeDeXelement(infoPrinted.PaginatedArtistList, number)
			} else {
				infoPrinted.PaginatedArtistList = groupieTrackers.ReconstituerEtDiviserEnListeDeXelement(infoPrinted.PaginatedArtistList, 52)
			}
		}
		artistPage.Execute(w, infoPrinted)
	})

	http.HandleFunc("/concert", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("info")
		idNum, _ := strconv.Atoi(id)
		if len(infoPrinted.ArtistList[idNum].Location) == 0 {
			infoPrinted.ArtistList = groupieTrackers.RecupRelation(infoPrinted.ArtistList, idNum)
			infoPrinted.PaginatedArtistList = groupieTrackers.DiviserEnListeDeXelement(infoPrinted.ArtistList, len(infoPrinted.PaginatedArtistList[0]))
		}
		concertPage.Execute(w, infoPrinted.ArtistList[idNum])
	})

	http.HandleFunc("/pays", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("info")
		idNum, _ := strconv.Atoi(id)
		paysPage.Execute(w, listLocation[idNum])
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
