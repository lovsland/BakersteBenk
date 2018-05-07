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

	http.HandleFunc("/fotball", api1handler)
	http.HandleFunc("/ManUnited", api2handler)
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
		Status       FixtureStatus   `json:"status"`
		Matchday     int       `json:"matchday"`
		HomeTeamName string    `json:"homeTeamName"`
		AwayTeamName string    `json:"awayTeamName"`
		Competition  string     `json:"competition"`
		Result       struct {
			GoalsHomeTeam interface{} `json:"goalsHomeTeam"`
			GoalsAwayTeam interface{} `json:"goalsAwayTeam"`
		} `json:"result"`
		Odds interface{} `json:"odds"`
	} `json:"fixtures"`
}

type ManU struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Team struct {
			Href string `json:"href"`
		} `json:"team"`
	} `json:"_links"`
	Season   string `json:"season"`
	Count    int    `json:"count"`
	Fixtures []struct {
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
			GoalsHomeTeam int `json:"goalsHomeTeam"`
			GoalsAwayTeam int `json:"goalsAwayTeam"`
			HalfTime      struct {
				GoalsHomeTeam int `json:"goalsHomeTeam"`
				GoalsAwayTeam int `json:"goalsAwayTeam"`
			} `json:"halfTime"`
		} `json:"result"`
		Odds interface{} `json:"odds"`
	} `json:"fixtures"`
}


type FixtureScore struct {
	GoalsHomeTeam uint16
	GoalsAwayTeam uint16
}

type FixtureResult struct {
	FixtureScore

	HalfTime        *FixtureScore
	ExtraTime       *FixtureScore
	PenaltyShootout *FixtureScore
}

type FixtureStatus string
const (
	FixtureStatus_Scheduled FixtureStatus = "SCHEDULED"
	FixtureStatus_Timed     FixtureStatus = "TIMED"
	FixtureStatus_Postponed FixtureStatus = "POSTPONED"
	FixtureStatus_InPlay    FixtureStatus = "IN_PLAY"
	FixtureStatus_Canceled  FixtureStatus = "CANCELED"
	FixtureStatus_Finished  FixtureStatus = "FINISHED"
)


var a ManU
var fotball Fotball

func api1handler (w http.ResponseWriter, r *http.Request) {

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


func api2handler(w http.ResponseWriter, req* http.Request) {

	url := "http://api.football-data.org/v1/teams/66/fixtures"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &a)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("competitionstest.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, a)
	if err != nil {
		log.Fatal(err)
	}
}



