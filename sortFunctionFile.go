package groupieTrackers

// ---------------------- In this file there are all the data sorting functions ---------------------- //

// Ascending alphabetical sorting (a --> z)
func AscendingAlphabeticalSorting(artistGroup []groupe) []groupe {
	if len(artistGroup) <= 1 {
		return artistGroup
	} else {
		lower := []groupe{}
		equals := []groupe{}
		higher := []groupe{}
		pivot := artistGroup[len(artistGroup)/2]
		for _, element := range artistGroup {
			if element.Name < pivot.Name {
				lower = append(lower, element)
			} else if element.Name > pivot.Name {
				higher = append(higher, element)
			} else if element.Name == pivot.Name {
				equals = append(equals, element)
			}
		}
		returned := AscendingAlphabeticalSorting(lower)
		returned = append(returned, equals...)
		returned = append(returned, AscendingAlphabeticalSorting(higher)...)
		return returned
	}
}

// Descending alphabetical sorting (z --> a)
func DescendingAlphabeticalSorting(artistGroup []groupe) []groupe {
	if len(artistGroup) <= 1 {
		return artistGroup
	} else {
		lower := []groupe{}
		equals := []groupe{}
		higher := []groupe{}
		pivot := artistGroup[len(artistGroup)/2]
		for _, element := range artistGroup {
			if element.Name < pivot.Name {
				lower = append(lower, element)
			} else if element.Name > pivot.Name {
				higher = append(higher, element)
			} else if element.Name == pivot.Name {
				equals = append(equals, element)
			}
		}
		returned := DescendingAlphabeticalSorting(higher)
		returned = append(returned, equals...)
		returned = append(returned, DescendingAlphabeticalSorting(lower)...)
		return returned
	}
}

// Sorting function for creation dates, if inAscendingOrder is true then list sorts in ascending order if it is false then in descending order
// Method used quick sorting
func SortingCreationDate(artistGroup []groupe, inAscendingOrder bool) []groupe {
	QuickSortingCreationDate(artistGroup, 0, len(artistGroup)-1, inAscendingOrder)
	return artistGroup
}

func QuickSortingCreationDate(artistGroup []groupe, indexStart int, indexEnd int, inAscendingOrder bool) {
	var pivotIndex int
	if indexStart < indexEnd {
		if inAscendingOrder {
			pivotIndex = SortingAscendingCreationDate(artistGroup, indexStart, indexEnd)
		} else if !inAscendingOrder {
			pivotIndex = SortingDescendingCreationDate(artistGroup, indexStart, indexEnd)
		}
		QuickSortingCreationDate(artistGroup, indexStart, pivotIndex-1, inAscendingOrder)
		QuickSortingCreationDate(artistGroup, pivotIndex+1, indexEnd, inAscendingOrder)
	}
}

func SortingAscendingCreationDate(artistGroup []groupe, indexStart int, indexEnd int) int {
	pivot := artistGroup[indexEnd]
	indexPivot := indexStart
	for i := indexStart; i <= indexEnd; i++ {
		if artistGroup[i].CreationDate < pivot.CreationDate {
			elementSave := artistGroup[i]
			artistGroup[i] = artistGroup[indexPivot]
			artistGroup[indexPivot] = elementSave
			indexPivot++
		}
	}
	elementSave := artistGroup[indexPivot]
	artistGroup[indexPivot] = artistGroup[indexEnd]
	artistGroup[indexEnd] = elementSave
	return indexPivot
}

func SortingDescendingCreationDate(artistGroup []groupe, indexStart int, indexEnd int) int {
	pivot := artistGroup[indexEnd]
	indexPivot := indexStart
	for i := indexStart; i <= indexEnd; i++ {
		if artistGroup[i].CreationDate > pivot.CreationDate {
			elementSave := artistGroup[i]
			artistGroup[i] = artistGroup[indexPivot]
			artistGroup[indexPivot] = elementSave
			indexPivot++
		}
	}
	elementSave := artistGroup[indexPivot]
	artistGroup[indexPivot] = artistGroup[indexEnd]
	artistGroup[indexEnd] = elementSave
	return indexPivot
}

// Function sorted by number of group members, method used bubble sort
func BubbleSortByNumberMemberAscending(artistGroup []groupe) []groupe {
	lenArtistGroup := len(artistGroup)
	hasChanged := true
	for hasChanged || lenArtistGroup > 1 {
		hasChanged = false
		for i := 0; i < len(artistGroup)-1; i++ {
			if len(artistGroup[i].Members) > len(artistGroup[i+1].Members) {
				elementSave := artistGroup[i]
				artistGroup[i] = artistGroup[i+1]
				artistGroup[i+1] = elementSave
				hasChanged = true
			}
		}
		lenArtistGroup--
	}
	return artistGroup
}

func BubbleSortByNumberMemberDescending(artistGroup []groupe) []groupe {
	lenArtistGroup := len(artistGroup)
	hasChanged := true
	for hasChanged || lenArtistGroup > 1 {
		hasChanged = false
		for i := 0; i < len(artistGroup)-1; i++ {
			if len(artistGroup[i].Members) < len(artistGroup[i+1].Members) {
				elementSave := artistGroup[i]
				artistGroup[i] = artistGroup[i+1]
				artistGroup[i+1] = elementSave
				hasChanged = true
			}
		}
		lenArtistGroup--
	}
	return artistGroup
}
