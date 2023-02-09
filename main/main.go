package main

import (
	"groupieTrackers"
	"net/http"
)

type display struct {
	test string
}

func main() {
	homePage := groupieTrackers.LoadTemplates()
	var displaye display
	// displaye.test = "éjhgrnejigndrjg"

	// Load all assets :
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	groupieTrackers.RecupInfo()

	// Load the first page of the game
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// level = r.FormValue("buttonLevel")
		// homePage.Execute(w, nil)
		homePage.Execute(w, displaye)

	})

}
