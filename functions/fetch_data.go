package functions

import (
	"encoding/json"
	"net/http"
	"sync"
)

func FetchArtists(response http.ResponseWriter) []Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
	}
	return artists
}

func fetchAdditionalData(artist *Artist, response http.ResponseWriter) {
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		locations := fetchData[Location](artist.LocationsURL, response)
		if locations != nil {
			artist.Locations = locations.Locations
		}
	}()

	go func() {
		defer wg.Done()
		dates := fetchData[Date](artist.ConcertDatesURL, response)
		if dates != nil {
			artist.ConcertDates = dates.Dates
		}
	}()

	go func() {
		defer wg.Done()
		relations := fetchData[RelationData](artist.RelationsURL, response)
		if relations != nil {
			artist.Relations = relations.DatesLocations
		}
	}()

	wg.Wait()
}

func fetchData[T Location | RelationData | Date](url string, response http.ResponseWriter) *T {
	resp, err := http.Get(url)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	var data T
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
	}
	return &data
}
