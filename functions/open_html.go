package functions

import (
	"html/template"
	"net/http"
)

func OpenHtml(fileName string, response http.ResponseWriter, data any) {
	temp, err := template.ParseFiles("templates/" + fileName)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
		return
	}
	err = temp.Execute(response, data)
	if err != nil {
		ErrorHandler(response, http.StatusInternalServerError)
		return
	}
}
