package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Time struct to create the api json data
type Time struct {
	Unix    int64  `json:"unix"`
	Natural string `json:"natural"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexFunc)
	router.HandleFunc("/{time}", getTime)

	http.ListenAndServe(":8080", router)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	req := mux.Vars(r)
	reqToInt, err := strconv.ParseInt(req["time"], 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reqToHuman := unixToHuman(reqToInt)
	response := Time{reqToInt, reqToHuman}

	jsonResp, _ := json.Marshal(response)

	fmt.Fprint(w, string(jsonResp))
}

func unixToHuman(value int64) string {
	time := time.Unix(value, 0)
	year, month, day := time.Date()
	return month.String() + " " + strconv.Itoa(day) + ", " + strconv.Itoa(year)
}

func indexFunc(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("index.html"))
	template.Execute(w, nil)
}
