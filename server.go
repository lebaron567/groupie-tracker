package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type groupe struct {
	Id           int      `json"id"`
	Image        string   `json"image"`
	Name         string   `json"name"`
	Members      []string `json"members"`
	CreationDate int      `json"creationDate"`
	FirstAlbum   string   `json"firstAlbum"`
	Locations    string   `json"locations"`
	ConcertDates string   `json"concertDates"`
	Relations    string   `json"relations"`
}

var g []groupe

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists" // adresse url
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body[1]))
	err := json.Unmarshal([]byte(body), &g)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("Person:", g[0])
}
