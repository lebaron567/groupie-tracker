package groupieTrackers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Dates        []string
	location     []string
	IsSearch     bool
}

// {"artists":"https://groupietrackers.herokuapp.com/api/artists","locations":"https://groupietrackers.herokuapp.com/api/locations","dates":"https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation"}
func RecupInfo() []groupe {
	var listGroups []groupe
	var groups groupe
	url := "https://groupietrackers.herokuapp.com/api/artists" // adresse url artist
	infoArtist := RecupInfoArtists(url)
	infoDate := RecupDates(infoArtist)
	infoLocation := RecupLocation(infoArtist)
	for i := 0; i < len(g); i++ {
		groups.Image = infoArtist[i].Image
		groups.Name = infoArtist[i].Name
		groups.Members = infoArtist[i].Members
		groups.CreationDate = infoArtist[i].CreationDate
		groups.FirstAlbum = infoArtist[i].FirstAlbum
		groups.Dates = infoDate[i].Dates
		groups.location = infoLocation[i].Locations
		groups.IsSearch = false
		listGroups = append(listGroups, groups)
	}
	listGroups[0].IsSearch = true
	return listGroups
}

func RecupInfoArtists(url string) []artist {
	req, _ := http.NewRequest("GET", url, nil)
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
