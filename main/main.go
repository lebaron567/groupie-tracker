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
	groupes := groupieTrackers.RecupInfo()
	fmt.Println(groupes)
	// fmt.Println(groupes)
	// for i := 0; i < len(groupes); i++ {
	// 	fmt.Print(groupes[i].Name)
	// }
	// fmt.Println(groupes.Name)
	// Load the first page of the game
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// level = r.FormValue("buttonLevel")
		// homePage.Execute(w, nil)
		homePage.Execute(w, groupes)
	})

	http.HandleFunc("/location", func(w http.ResponseWriter, r *http.Request) {
		locationPage.Execute(w, groupes)
	})

	http.HandleFunc("/artiste", func(w http.ResponseWriter, r *http.Request) {
		artistPage.Execute(w, groupes)
	})
	fmt.Println("Serveur start at : http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
