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

func SortElement(sortingChoices string, lenNewlistGroups int) [][]groupe {
	listGroups := RecupInfo()
	if sortingChoices == "AscendingAlphabeticalSorting" {
		listGroups = AscendingAlphabeticalSorting(listGroups)
	} else if sortingChoices == "DescendingAlphabeticalSorting" {
		listGroups = DescendingAlphabeticalSorting(listGroups)
	} else if sortingChoices == "SortingAscendingCreationDate" {
		listGroups = SortingCreationDate(listGroups, true)
	} else if sortingChoices == "SortingDescendingCreationDate" {
		listGroups = SortingCreationDate(listGroups, false)
	} else if sortingChoices == "BubbleSortByNumberMemberAscending" {
		listGroups = BubbleSortByNumberMemberAscending(listGroups)
	} else if sortingChoices == "BubbleSortByNumberMemberDescending" {
		listGroups = BubbleSortByNumberMemberDescending(listGroups)
	}
	nlistGroups := DiviserEnListeDeXelement(listGroups, lenNewlistGroups)
	return nlistGroups
}

func SearchGroupe(nameSearch string, artistGroup [][]groupe) [][]groupe {
	newArtistGroup := ReconstituerList(artistGroup)
	g2 := []groupe{}
	for index, element := range newArtistGroup {
		if element.Name == nameSearch {
			g2 = append(g2, element)
			g2[0].IsSearch = true
		} else {
			newArtistGroup[index].IsSearch = false
		}
	}
	for _, element := range newArtistGroup {
		if element.Name != nameSearch {
			g2 = append(g2, element)
		}
	}
	newArtistGroup2 := DiviserEnListeDeXelement(g2, len(artistGroup[0]))
	return newArtistGroup2
}

func DiviserEnListeDeXelement(artistGroup []groupe, x int) [][]groupe {
	newArtistGroup := [][]groupe{}
	page := []groupe{}
	for index := range artistGroup {
		page = append(page, artistGroup[index])
		if len(page) == x {
			newArtistGroup = append(newArtistGroup, page)
			page = []groupe{}
		}
	}
	return newArtistGroup
}

func ReconstituerList(artistGroup [][]groupe) []groupe {
	newGroupe := []groupe{}
	for index := range artistGroup {
		for index2 := range artistGroup[index] {
			newGroupe = append(newGroupe, artistGroup[index][index2])
		}
	}
	return newGroupe
}

func ReconstituerEtDiviserEnListeDeXelement(artistGroup [][]groupe, x int) [][]groupe {
	newGroupe := ReconstituerList(artistGroup)
	newArtistGroup := DiviserEnListeDeXelement(newGroupe, x)
	return newArtistGroup
}
