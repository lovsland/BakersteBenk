package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/fotball", response)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

type Fotball struct {
	TimeFrameStart string `json:"timeFrameStart"`
	TimeFrameEnd   string `json:"timeFrameEnd"`
	Count          int    `json:"count"`
	Fixtures       []struct {
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Competition struct {
				Href string `json:"href"`
			} `json:"competition"`
			HomeTeam struct {
				Href string `json:"href"`
			} `json:"homeTeam"`
			AwayTeam struct {
				Href string `json:"href"`
			} `json:"awayTeam"`
		} `json:"_links"`
		Date         time.Time `json:"date"`
		Status       string    `json:"status"`
		Matchday     int       `json:"matchday"`
		HomeTeamName string    `json:"homeTeamName"`
		AwayTeamName string    `json:"awayTeamName"`
		Result       struct {
			GoalsHomeTeam interface{} `json:"goalsHomeTeam"`
			GoalsAwayTeam interface{} `json:"goalsAwayTeam"`
		} `json:"result"`
		Odds interface{} `json:"odds"`
	} `json:"fixtures"`
}

var fotball Fotball

func response(w http.ResponseWriter, r *http.Request) {

	url := "http://api.football-data.org/v1/fixtures/"
	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &fotball)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("fotballtest.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, fotball)
	if err != nil {
		log.Fatal(err)
	}
}


