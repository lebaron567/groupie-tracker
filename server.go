package groupieTrackers

import (
	"text/template"
)

// Function that load all templates for all pages
func LoadTemplates() (*template.Template, *template.Template, *template.Template, *template.Template,*template.Template) {
	homePage := template.Must(template.ParseFiles("./front/index.html"))
	artistPage := template.Must(template.ParseFiles("./front/artiste.html"))
	locationPage := template.Must(template.ParseFiles("./front/location.html"))
	concertPage := template.Must(template.ParseFiles("./front/concert.html"))
	paysPage := template.Must(template.ParseFiles("./front/pays.html"))

	return homePage, artistPage, locationPage, concertPage, paysPage
}

func SearchGroupe(nameSearch string, g []groupe) []groupe {
	g2 := []groupe{}
	for index, element := range g {
		if element.Name == nameSearch {
			g2 = append(g2, element)
			g2[0].IsSearch = true
		} else {
			g[index].IsSearch = false
		}
	}
	for _, element := range g {
		if element.Name != nameSearch {
			g2 = append(g2, element)
		}
	}
	return g2
}

func DiviserEnDeux(g []groupe) ([]groupe, []groupe) {
	page1 := []groupe{}
	page2 := []groupe{}
	for index, element := range g {
		if index <= (len(g)-1)/2 {
			page1 = append(page1, element)
		} else {
			page2 = append(page2, element)
		}
	}
	return page1, page2
}
