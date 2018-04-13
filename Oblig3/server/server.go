package main

import (
	"io"
	"net/http"
	"log"
	//"Oblig3/api1"

	"io/ioutil"
	"fmt"
)

// handling homepage
func handler(w http.ResponseWriter, req *http.Request) {
io.WriteString(w, "Hi Client!\n")
}

//handling API 1
func api1handler(w http.ResponseWriter, req* http.Request) {
	url := "https://hotell.difi.no/api/json/stavanger/miljostasjoner"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseString := string(responseData)
	fmt.Fprint(w, responseString)
	io.WriteString(w, "API 1 details")
}

func main() {
http.HandleFunc("/", handler)
http.HandleFunc("/1", api1handler)
log.Fatal(http.ListenAndServe(":8080", nil))
}

