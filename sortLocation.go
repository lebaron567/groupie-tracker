package groupieTrackers

import (
	"strings"
)

// type concertByLocation struct {
// 	Pays   string
// 	Ville  string
// 	Groupe []string
// 	Date   [][]string
// }

// func SortLieux(data []groupe) []concertByLocation {
// 	lieux := []concertByLocation{}
// 	var lieu concertByLocation
// 	var newLieu bool =true
// 	for i := 0; i < len(data); i++ {
// 		for x := 0; x < len(data[i].Location); x++ {
// 			for y := 0; y < len(lieux); y++ {
// 				if strings.Split(data[i].Location[x], "-")[1] == string(lieux[y].Pays) && strings.Split(data[i].Location[x], "-")[0] == string(lieux[y].Ville) {
// 					lieux[y].Groupe = append(lieux[y].Groupe,data[i].Name )
// 					lieux[y].Date = append(lieux[y].Date, data[i].Dates[x])
// 					newLieu = false
// 				}
// 			}
// 			if newLieu {
// 				lieu.Groupe = []string{data[i].Name}
// 				lieu.Date = append(lieu.Date, data[i].Dates[x])
// 				lieu.Pays = strings.Split(data[i].Location[x], "-")[1]
// 				lieu.Ville = strings.Split(data[i].Location[x], "-")[0]
// 				lieux = append(lieux, lieu)
// 			}else{
// 				newLieu = true
// 			}
// 			lieu = concertByLocation{}
// 		}
// 	}
// 	return lieux
// }

type concertByLocation struct {
	Pays   string
	Villes []ville
}

type ville struct {
	Name   string
	Groupe []string
	Date   [][]string
}

func SortLieux(data []groupe) []concertByLocation {
	lieux := []concertByLocation{}
	var lieu concertByLocation
	var newPays bool = true
	for i := 0; i < len(data); i++ {
		for x := 0; x < len(data[i].Location); x++ {
			for y := 0; y < len(lieux); y++ {
				if strings.Split(data[i].Location[x], "-")[1] == string(lieux[y].Pays) {
					newVille, jj := verifAddVille(lieux[y].Villes, data[i].Location[x])
					newPays = false
					if !newVille {
						var tempo ville
						tempo.Name = strings.Split(data[i].Location[x], "-")[0]
						tempo.Groupe = append(tempo.Groupe, data[i].Name)
						tempo.Date = append(tempo.Date, data[i].Dates[x])
						lieux[y].Villes = append(lieux[y].Villes, tempo)
					} else {
						lieux[y].Villes[jj].Groupe = append(lieux[y].Villes[jj].Groupe, data[i].Name)
						lieux[y].Villes[jj].Date = append(lieux[y].Villes[jj].Date, data[i].Dates[x])
					}
				}
			}
			if newPays {
				lieu.Pays = strings.Split(data[i].Location[x], "-")[1]
				var tempo ville
				tempo.Name = strings.Split(data[i].Location[x], "-")[0]
				tempo.Groupe = append(tempo.Groupe, data[i].Name)
				tempo.Date = append(tempo.Date, data[i].Dates[x])
				lieu.Villes = append(lieu.Villes, tempo)
				lieux = append(lieux, lieu)
			} else {
				newPays = true
			}
			lieu = concertByLocation{}
		}
	}
	return lieux
}

func verifAddVille(data []ville, data2 string) (bool, int) {
	for index, Ville := range data {
		if strings.Split(data2, "-")[0] == Ville.Name {
			return true, index
		}
	}
	return false, 0
}
