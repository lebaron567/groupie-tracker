package groupieTrackers

import "text/template"

// Function that load all templates for all pages
func LoadTemplates() *template.Template {
	homePage := template.Must(template.ParseFiles("./templates/index.html"))
	// gamePage := template.Must(template.ParseFiles("./templates/hangman.html"))
	// winPage := template.Must(template.ParseFiles("./templates/win.html"))
	// failPage := template.Must(template.ParseFiles("./templates/fail.html"))

	return homePage
}
