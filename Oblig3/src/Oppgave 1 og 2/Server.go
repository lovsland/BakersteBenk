package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"io"
)

func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/1", api1handler)
	http.HandleFunc("/2", api2handler)
	http.HandleFunc("/3", api3handler)
	http.HandleFunc("/4", api4handler)
	http.HandleFunc("/5", api5handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hi Client!\n")

}

// lager structs for aa holde paa JSON data

type Badeplasser struct {
	Entries []struct {
		Name string `json:"name"`
		Longitude string `json:"longitude"`
		Latitude string `json:"latitude"`
	}	`json:"entries"`
}

type Skoleruter struct {
	Entries []struct {
		Sfodag string `json:"sfodag"`
		Skole string `json:"skole"`
		Kommentar string `json:"kommentar"`
		Elevdag string `json:"elevdag"`
		Dato string `json:"dato"`
		Laererdag string `json:"laererdag"`
	}	`json:"entries"`
}

type Offentligtoalett struct {
	Entries []struct {
		Plassering string `json:"plassering"`
		Adresse string `json:"adresse"`
		AapningstiderSoendag string `json:"aapningstider_soendag"`
		AapningstiderHverdag string `json:"aapningstider_hverdag"`
		AapningstiderLoerdag string `json:"aapningstider_loerdag"`
		Pris string `json:"pris"`
		Stellerom string `json:"stellerom"`
		Rullestol string `json:"rullestol"`
		Longitude string `json:"longitude"`
		Latitude string `json:"latitude"`
	}	`json:"entries"`
}


type Kommune struct {
	Entries []struct {
		Navn string`json:"navn"`
		Kommunenr string `json:"kommune"`
		Fylke string `json:"fylke"`
	} `json:"entries"`
}

type Stasjon struct {
	Entries []struct {
		Latitude  string `json:"latitude"`
		Navn      string `json:"navn"`
		Plast   string `json:"plast"`
		GlassMetal string `json:"glass_metall"`
		TekstilSko      string `json:"tekstil_sko,omitempty"`
		Longitude string `json:"longitude"`
	} `json:"entries"`
}

var m Stasjon
var kommune Kommune
var badeplasser Badeplasser
var offentligtoalett Offentligtoalett
var skoleruter Skoleruter


func api1handler(w http.ResponseWriter, req* http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &m)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	tmpl, err := template.ParseFiles("miljostasjoner.html")
	if err != nil {
		log.Print(err)
	}

	err = tmpl.Execute(w, m)
	if err != nil {
		log.Fatal(err)
	}
}



func api2handler(w http.ResponseWriter, req* http.Request) {

	url := "https://hotell.difi.no/api/json/difi/geo/kommune"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &kommune)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("kommuner.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, kommune)
	if err != nil {
		log.Fatal(err)
	}
}

func api3handler (w http.ResponseWriter, req* http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/badeplasser"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &badeplasser)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("badeplasser.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, badeplasser)
	if err != nil {
		log.Fatal(err)
	}

}

func api4handler (w http.ResponseWriter, req* http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/offentligetoalett"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &offentligtoalett)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("offentligetoalett.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, offentligtoalett)
	if err != nil {
		log.Fatal(err)
	}
}

func api5handler (w http.ResponseWriter, req* http.Request) {

	url := "https://hotell.difi.no/api/json/stavanger/skoleruter"

	result, _ := http.Get(url)

	body, _ := ioutil.ReadAll(result.Body)

	jsonErr := json.Unmarshal(body, &skoleruter)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	t, err := template.ParseFiles("skoleruter.html")
	if err != nil {
		log.Print(err)
	}

	err = t.Execute(w, skoleruter)
	if err != nil {
		log.Fatal(err)
	}
}
