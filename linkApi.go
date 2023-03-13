package groupieTrackers

// ---------------------- In this file there are all the in this file there are all the relations with the api ---------------------- //

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Liens des api
// {"artists":"https://groupietrackers.herokuapp.com/api/artists","locations":"https://groupietrackers.herokuapp.com/api/locations","dates":"https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation"}

// Structure pour récupérer les données des artistes
type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Création varriable globale qui va contenir tous les artistes de l'api
var artistList []artist

// Varriable qui va lier les info de l'api artist et les infos de l'api relation
type groupe struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Dates        [][]string
	Location     []string
}

// Varriable qui serra envoyer sur la page afin d'être afficher
type printedInfo struct {
	ArtistList          []groupe
	PaginatedArtistList [][]groupe
	IsNotFind           bool
	IndexCurrentPage    int
}

func Init() printedInfo {
	listGroups := RecupInfoArtist()
	var infoPrinted printedInfo
	infoPrinted.PaginatedArtistList = DiviserEnListeDeXelement(listGroups, len(listGroups))
	infoPrinted.ArtistList = listGroups
	infoPrinted.IsNotFind = false
	return infoPrinted
}

// Ajout dans la liste de groupes/artistes des info artistes
func RecupInfoArtist() []groupe {
	var listGroups []groupe
	var groups groupe
	infoArtist := RecupInfoArtists()
	for i := 0; i < len(infoArtist); i++ {
		groups.Image = infoArtist[i].Image
		groups.Name = infoArtist[i].Name
		groups.Members = infoArtist[i].Members
		groups.CreationDate = infoArtist[i].CreationDate
		groups.FirstAlbum = infoArtist[i].FirstAlbum
		listGroups = append(listGroups, groups)
	}
	listGroups = initialisationRelation(listGroups)
	return listGroups
}

// Récupération info api artiste
func RecupInfoArtists() []artist {
	url := "https://groupietrackers.herokuapp.com/api/artists" // adresse url artist
	req, _ := http.NewRequest("GET", url, nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &artistList)
	if err != nil {
		fmt.Println("Error", err)
	}
	return artistList
}

// Varriable qui va contenir tous les dates des artistes de l'api
type date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Récupération info api dates
func RecupDates(g []artist) []date {
	var listDate []date
	for i := 0; i < len(g); i++ {
		url := g[i].ConcertDates // adresse url
		req, _ := http.NewRequest("GET", url, nil)
		res, erre := http.DefaultClient.Do(req)
		if erre != nil {
			fmt.Println("Error", erre)
		}
		var d date
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		err := json.Unmarshal([]byte(body), &d)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			listDate = append(listDate, d)
		}
	}
	return listDate
}

// Varriable qui va contenir tous les locations des artistes de l'api
type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Récupération info api location
func RecupLocation(g []artist) []location {
	var lisrRelation []location
	for i := 0; i < len(g); i++ {
		url := g[i].Locations // adresse url
		req, _ := http.NewRequest("GET", url, nil)
		res, erre := http.DefaultClient.Do(req)
		if erre != nil {
			fmt.Println("Error", erre)
		}
		var l location
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		err := json.Unmarshal([]byte(body), &l)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			lisrRelation = append(lisrRelation, l)
		}
	}
	return lisrRelation
}

// Initialisation des varriables qui contiendrons les infos de relations
func initialisationRelation(listGroups []groupe) []groupe {
	for index := range listGroups {
		listGroups[index].Location = []string{}
		listGroups[index].Dates = [][]string{}
	}
	return listGroups
}

// Création varriable nécessaire à la récupération de données de l'api relation
type indexage struct {
	Index []relation `json:"index"`
}
type relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Implémentation de toutes les varriables avec les infos de relations
func RecupAllRelation(g []groupe) []groupe {
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/relation", nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	var i indexage
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &i)
	if err != nil {
		fmt.Println("Error", err)
	}
	for index := 0; index < len(g); index++ {
		for location := range i.Index[index].DatesLocations {
			g[index].Location = append(g[index].Location, location)
			g[index].Dates = append(g[index].Dates, i.Index[index].DatesLocations[location])
		}
	}
	return g
}

// Implémentation d'une varriables avec les infos nécessaire de relations
func RecupRelation(listGroups []groupe, indexGroupImplemented int) []groupe {
	id := strconv.Itoa(listGroups[indexGroupImplemented].Id)
	url := "https://groupietrackers.herokuapp.com/api/relation" + "/" + id
	req, _ := http.NewRequest("GET", url, nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	var i indexage
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &i)
	if err != nil {
		fmt.Println("Error", err)
	}
	for location := range i.Index[0].DatesLocations {
		listGroups[indexGroupImplemented].Location = append(listGroups[indexGroupImplemented].Location, location)
		listGroups[indexGroupImplemented].Dates = append(listGroups[indexGroupImplemented].Dates, i.Index[0].DatesLocations[location])
	}
	return listGroups
}
