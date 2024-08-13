package functions

import (
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Handler(response http.ResponseWriter, request *http.Request) {
	artistsData := FetchArtists(response)
	if request.Method == http.MethodGet {
		if request.URL.Path == "/" {
			OpenHtml("index.html", response, artistsData)
		} else if strings.Contains(request.URL.Path, "/artist/") {
			ArtistHandler(response, request, artistsData)
		} else {
			ErrorHandler(response, http.StatusNotFound)
		}
	} else {
		ErrorHandler(response, http.StatusMethodNotAllowed)
	}
}

func ArtistHandler(response http.ResponseWriter, request *http.Request, artistsData []Artist) {
	idStr := request.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(response, http.StatusNotFound)
		return
	}
	var artist Artist
	for _, a := range artistsData {
		if a.ID == id {
			artist = a
			break
		}
	}
	if artist.ID == 0 {
		ErrorHandler(response, http.StatusNotFound)
		return
	}
	fetchAdditionalData(&artist, response)
	OpenHtml("artist.html", response, artist)
}

func SetupStaticFilesHandlers(response http.ResponseWriter, request *http.Request) {
	fileinfo, err := os.Stat("." + request.URL.Path)
	if !os.IsNotExist(err) && !fileinfo.IsDir() {
		http.FileServer(http.Dir("./")).ServeHTTP(response, request)
	} else {
		ErrorHandler(response, http.StatusNotFound)
		return
	}
}

func ErrorHandler(response http.ResponseWriter, status int) {
	errorMap := map[int][]string{
		404: {"Page not found", "Sorry, the page you are looking for does not exist.", "You may have mistyped the address or the page may have moved."},
		405: {"Method Not Allowed", "Sorry, the method you are using is not allowed for this page.", "Please check the URL and try again with a different request method."},
		500: {"Internal Server Error", "Sorry, something went wrong on our end.", "Please try again later."},
	}
	var error Error
	response.WriteHeader(status)
	for key, value := range errorMap {
		if key == status {
			error.Err = key
			error.Name = value[0]
			error.Messages = value[1:]
			OpenHtml("error.html", response, error)
		}
	}
	errorMap = nil
}
