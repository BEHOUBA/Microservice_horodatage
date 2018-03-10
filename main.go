package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

// Time struct to create the api json data
type Time struct {
	Unix    int    `json:"unix"`
	Natural string `json:"natural"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexFunc)
	router.HandleFunc("/{time}", getTime)

	http.ListenAndServe(":8080", router)
}

// getTime function give produce the json data to be returned depending on user query

func getTime(w http.ResponseWriter, r *http.Request) {
	var jsonResponse string
	req := mux.Vars(r)
	timeString := req["time"]
	reqToInt, err := strconv.Atoi(timeString)
	if err != nil {
		reqToUnix, err := humanToUnix(timeString)
		if err != nil {
			jsonResponse = jsonMaker(0, "null")
			fmt.Fprint(w, jsonResponse)
			return
		}
		jsonResponse = jsonMaker(reqToUnix, timeString)
		fmt.Fprint(w, jsonResponse)
		return
	}

	jsonResponse = jsonMaker(reqToInt, unixToHuman(reqToInt))

	fmt.Fprint(w, jsonResponse)
}

// jsonMaker
func jsonMaker(unix int, human string) string {
	response := Time{unix, human}
	jsonData, _ := json.Marshal(response)
	return string(jsonData)
}

// unixToHuman function take unix date format in return the corresponding human readable format
func unixToHuman(value int) string {
	time := time.Unix(int64(value), 0)
	year, month, day := time.Date()
	return month.String() + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year)
}

// humanToUnix function take a human readable date format and return the correspondin unix date format
func humanToUnix(value string) (int, error) {
	var dateFormated []string

	dateLayout := "Jan-2-2006"
	dateSlice := strings.Split(strings.Replace(value, ",", "", -1), " ")

	for _, str := range dateSlice {
		checkValue, err := strconv.Atoi(str)
		if err != nil {
			dateFormated = append(dateFormated, monthFormatter(str))
		} else {
			if checkValue <= 31 {
				dateFormated = append(dateFormated, str)
			} else {
				dateFormated = append(dateFormated, str)
			}
		}
	}
	date, err := time.Parse(dateLayout, strings.Join(dateFormated, "-"))
	if err != nil {
		return 0, err
	}
	return int(date.Unix()), nil
}

// the indexFunc generate the home page of the projet
func indexFunc(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("index.html"))
	template.Execute(w, nil)
}

// the monthFormatter function format de month by taking only the first 3 letter and making the first one uppercase
func monthFormatter(month string) string {
	if len(month) < 3 {
		return ""
	}
	firstThreeLetter := month[:3]
	letterSlice := strings.Split(firstThreeLetter, "")
	letterSlice[0] = strings.ToUpper(letterSlice[0])
	result := strings.Join(letterSlice, "")
	return result
}
