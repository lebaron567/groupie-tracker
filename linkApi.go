package groupieTrackers

// ---------------------- In this file there are all the in this file there are all the relations with the api ---------------------- //

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// api link
// {"artists":"https://groupietrackers.herokuapp.com/api/artists","locations":"https://groupietrackers.herokuapp.com/api/locations","dates":"https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation"}

// Structure for retrieving artist data
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

// Creation of a global variable that will contain all the artists of the api
var artistList []artist

// Variable that will link the info of the api artist and the info of the api relation
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

// Variable that will be sent to the page to be displayed
type printedInfo struct {
	ArtistList          []groupe
	PaginatedArtistList [][]groupe
	IsNotFind           bool
	IndexCurrentPage    int
}

// Function that initializes the relationship with the api needed to launch the site
func Init() printedInfo {
	listGroups := RecupInfoArtist()
	var infoPrinted printedInfo
	infoPrinted.PaginatedArtistList = DiviserEnListeDeXelement(listGroups, len(listGroups))
	infoPrinted.ArtistList = listGroups
	infoPrinted.IsNotFind = false
	return infoPrinted
}

// Add in the list of groups/artists of the artists info
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

// Artist info retrieval
func RecupInfoArtists() []artist {
	url := "https://groupietrackers.herokuapp.com/api/artists" // address url artist
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

// Variable that will contain all the dates of the artists of the api
type date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Recovering info api dates
func RecupDates(g []artist) []date {
	var listDate []date
	for i := 0; i < len(g); i++ {
		url := g[i].ConcertDates // url address
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

// Variable that will contain all the artist rentals of the api
type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

//Recovering info api location
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

// Initialization of the variables that will contain the relationship information
func initialisationRelation(listGroups []groupe) []groupe {
	for index := range listGroups {
		listGroups[index].Location = []string{}
		listGroups[index].Dates = [][]string{}
	}
	return listGroups
}

// Creation of variable necessary to retrieve data from the relation api
type indexage struct {
	Index []relation `json:"index"`
}
type relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Implementation of all variables with relationship info
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

// Implementation of a variable with the necessary info from relations
func RecupRelation(listGroups []groupe, indexGroupImplemented int) []groupe {
	id := strconv.Itoa(listGroups[indexGroupImplemented].Id + 1)
	url := "https://groupietrackers.herokuapp.com/api/relation" + "/" + id
	req, _ := http.NewRequest("GET", url, nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	var r relation
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &r)
	if err != nil {
		fmt.Println("Error", err)
	}
	for location := range r.DatesLocations {
		listGroups[indexGroupImplemented].Location = append(listGroups[indexGroupImplemented].Location, location)
		listGroups[indexGroupImplemented].Dates = append(listGroups[indexGroupImplemented].Dates, r.DatesLocations[location])
	}
	return listGroups
}
