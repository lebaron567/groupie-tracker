package groupieTrackers

import "text/template"

// Function that load all templates for all pages
func LoadTemplates() (*template.Template, *template.Template, *template.Template) {
	homePage := template.Must(template.ParseFiles("./front/index.html"))
	artistPage := template.Must(template.ParseFiles("./front/artiste.html"))
	locationPage := template.Must(template.ParseFiles("./front/location.html"))
	return homePage, artistPage, locationPage
}

func SearchGroupe(nameSearch string, g []groupe) []groupe {
	for index, element := range g {
		if element.Name == nameSearch {
			save := g[index]
			g[index] = g[0]
			g[index].IsSearch = false
			g[0] = save
			g[0].IsSearch = true
		} else {
			g[index].IsSearch = false
		}
	}
	return g
}
