package groupieTrackers

import "text/template"

// Function that load all templates for all pages
func LoadTemplates() (*template.Template, *template.Template, *template.Template) {
	homePage := template.Must(template.ParseFiles("./front/index.html"))
	artistPage := template.Must(template.ParseFiles("./front/artiste.html"))
	locationPage := template.Must(template.ParseFiles("./front/location.html"))

	return homePage, artistPage, locationPage
}
