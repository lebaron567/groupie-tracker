package groupieTrackers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type groupe struct {
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

// {"artists":"https://groupietrackers.herokuapp.com/api/artists","locations":"https://groupietrackers.herokuapp.com/api/locations","dates":"https://groupietrackers.herokuapp.com/api/dates","relation":"https://groupietrackers.herokuapp.com/api/relation"}
func RecupArtists() []groupe {
	var g []groupe
	url := "https://groupietrackers.herokuapp.com/api/artists" // adresse url
	req, _ := http.NewRequest("GET", url, nil)
	res, erre := http.DefaultClient.Do(req)
	if erre != nil {
		fmt.Println("Error", erre)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	//fmt.Println(string(body))
	err := json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("Person:", g[1])
	}
	return g
}

func RecupDates(g []groupe) []date {
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
		//fmt.Println(string(body))
		err := json.Unmarshal([]byte(body), &d)
		if err != nil {
			fmt.Println("Error", err)
		} else {
			listDate = append(listDate, d)
		}
	}
	fmt.Print(listDate)
	return listDate
}

func RecupLocation(g []groupe) []location {
	var listDate []location
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
			listDate = append(listDate, l)
		}
	}
	return listDate
}
