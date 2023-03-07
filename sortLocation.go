package groupieTrackers

import (
	"strings"
)

type concertByLocation struct {
	Pays   string
	Ville  string
	Groupe []string
	Date   [][]string
}

func SortLieux(data []groupe) []concertByLocation {
	lieux := []concertByLocation{}
	var lieu concertByLocation
	var newLieu bool =true
	for i := 0; i < len(data); i++ {
		for x := 0; x < len(data[i].Location); x++ {
			for y := 0; y < len(lieux); y++ {
				if strings.Split(data[i].Location[x], "-")[1] == string(lieux[y].Pays) && strings.Split(data[i].Location[x], "-")[0] == string(lieux[y].Ville) {
					lieux[y].Groupe = append(lieux[y].Groupe,data[i].Name ) 
					lieux[y].Date = append(lieux[y].Date, data[i].Dates[x])
					newLieu = false	
				}
			}
			if newLieu {
				lieu.Groupe = []string{data[i].Name}
				lieu.Date = append(lieu.Date, data[i].Dates[x])
				lieu.Pays = strings.Split(data[i].Location[x], "-")[1]
				lieu.Ville = strings.Split(data[i].Location[x], "-")[0]
				lieux = append(lieux, lieu)
			}else{
				newLieu = true
			}
			lieu = concertByLocation{}
		}
	}
	return lieux
}
