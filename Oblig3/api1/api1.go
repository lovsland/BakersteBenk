package api1

import (
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type stasjon struct {
	Entries []struct {
		Latitude    string `json:"latitude"`
		Navn        string `json:"navn"`
		Plast       string `json:"plast"`
		GlassMetall string `json:"glass_metall"`
		TekstilSko  string `json:"tekstil_sko,omitempty"`
		Longitude   string `json:"longitude"`
	} `json:"entries"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
	Posts int `json:"posts"`
}

func Api1(){
//get url from api
res,err :=http.Get("https://hotell.difi.no/api/json/stavanger/miljostasjoner")
if err != nil{
log.Fatal(err)
}
//read body
body,err := ioutil.ReadAll(res.Body)
res.Body.Close()
if err !=nil {
log.Fatal(err)
}

//decoding json
var m stasjon
err = json.Unmarshal(body, &m)
if err != nil {
log.Fatal(err)
}

fmt.Printf("Entries: %s", m.Entries )
}



