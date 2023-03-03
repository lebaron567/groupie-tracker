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
		sortingChoices := r.FormValue("sortingChoices")
		searchUser := r.FormValue("userSearch")
		if sortingChoices != "" {
			if sortingChoices == "AscendingAlphabeticalSorting" {
				listGroups = groupieTrackers.AscendingAlphabeticalSorting(listGroups)
			} else if sortingChoices == "DescendingAlphabeticalSorting" {
				listGroups = groupieTrackers.DescendingAlphabeticalSorting(listGroups)
			} else if sortingChoices == "SortingAscendingCreationDate" {
				listGroups = groupieTrackers.SortingCreationDate(listGroups, true)
			} else if sortingChoices == "SortingDescendingCreationDate" {
				listGroups = groupieTrackers.SortingCreationDate(listGroups, false)
			} else if sortingChoices == "BubbleSortByNumberMemberAscending" {
				listGroups = groupieTrackers.BubbleSortByNumberMemberAscending(listGroups)
			} else if sortingChoices == "BubbleSortByNumberMemberDescending" {
				listGroups = groupieTrackers.BubbleSortByNumberMemberDescending(listGroups)
			}
		}
		if searchUser != "" {
			listGroups = groupieTrackers.SearchGroupe(searchUser, listGroups)
		} else {
			listGroups[0].IsSearch = true
		}
		artistPage.Execute(w, listGroups)
	})

	http.HandleFunc("/concert", func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("info")
		idNum, _ := strconv.Atoi(id)
		fmt.Println(idNum)
		concertPage.Execute(w, listGroups[idNum-1])
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
