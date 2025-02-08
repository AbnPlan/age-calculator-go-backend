package main

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/gorilla/schema"
)

type Date struct {
	Day int `schema:"day"`
	Month int `schema:"month"`
	Year int `schema:"year"`
}

type AgeCalcResponse struct {
	Age int
}

func HandleCalcAge(requestWriter http.ResponseWriter, request *http.Request) {
	requestWriter.Header().Set("Content-Type", "application/json")
	requestWriter.WriteHeader(http.StatusCreated)

	// Validate that the URL path is correct
	if request.URL.Path != "/age-calc" {
		http.Error(requestWriter, "404 not found.", http.StatusNotFound)
		return
	}

	// Populate request form
	if err := request.ParseForm(); err != nil {
		
    }

    // Create an empty Date
    birthDate := new(Date)
    ageResponse := new(AgeCalcResponse)

    // Decode Form into Date struct
    if err := schema.NewDecoder().Decode(birthDate, request.Form); err != nil {
        http.Error(requestWriter, "422 Invalid Query.", http.StatusUnprocessableEntity)
		return
    }

    birthDateObj := time.Date(birthDate.Year, time.Month(birthDate.Month), birthDate.Day, 0, 0, 0, 0, time.UTC)
    ageResponse.Age = calculateAge(birthDateObj)

    json.NewEncoder(requestWriter).Encode(ageResponse)
}

// Calculates age by comparing provided date with input age
func calculateAge(birthDate time.Time) int {
	var currentDate time.Time = time.Now().UTC()

	var age int = int(currentDate.Sub(birthDate).Hours()/24/365)

	if age < 0 || age > 292 {
		age = 0
	}

	return age
}

func main() {
	http.HandleFunc("/age-calc", HandleCalcAge)
	http.ListenAndServe(":3432", nil)
}
