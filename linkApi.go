package groupieTrackers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
)

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

type location struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var g []artist

type groupe struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Dates        [][]string
	Location     []string
	IsSearch     bool
}

// {"artists":"https://groupietrackers.herokuapp.com/api/artists","locations":"https://groupietrackers.herokuapp.com/api/locations","dates":"https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation"}
func RecupInfo() []groupe {
	var listGroups []groupe
	var groups groupe
	url := "https://groupietrackers.herokuapp.com/api/artists" // adresse url artist
	infoArtist := RecupInfoArtists(url)
	for i := 0; i < len(g); i++ {
		groups.Image = infoArtist[i].Image
		groups.Name = infoArtist[i].Name
		groups.Members = infoArtist[i].Members
		groups.CreationDate = infoArtist[i].CreationDate
		groups.FirstAlbum = infoArtist[i].FirstAlbum
		groups.IsSearch = false
		listGroups = append(listGroups, groups)
	}
	listGroups[0].IsSearch = true
	listGroups = RecupRealtion(listGroups)
	return listGroups
}

func RecupInfoArtists(url string) []artist {
	req, _ := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/artists", nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error", err)
	}
	return g
}

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

type indexage struct {
	Index []relation `json:"index"`
}
type relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func RecupRealtion(g []groupe) []groupe {
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
			g[index].Location = append(g[index].Location, location) // rajouter un s à locations
			g[index].Dates = append(g[index].Dates, i.Index[index].DatesLocations[location])
		}
	}
	return g
}

func RomdomArtist() groupe {
	var artistRamdom artist = artist{}
	randomInt := rand.Intn(51)
	for randomInt==0{
		randomInt = rand.Intn(51)
	}
	url := "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(randomInt)
	req, _ := http.NewRequest("GET", url, nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	err := json.Unmarshal([]byte(body), &artistRamdom)
	if err != nil {
		fmt.Println("Error", err)
	}
	var groups groupe
	groups.Image = artistRamdom.Image
	groups.Name = artistRamdom.Name
	groups.Members = artistRamdom.Members
	groups.CreationDate = artistRamdom.CreationDate
	groups.FirstAlbum = artistRamdom.FirstAlbum
	return groups
}
