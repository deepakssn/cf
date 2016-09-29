package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	dbconnect()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/domains", GetDomain)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index is a standard index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// GetDomain will send the list of allowed domains and blocked domains
func GetDomain(w http.ResponseWriter, r *http.Request) {
	// son := []byte(`{ "allowedDomains": ["infosys.com","skidata.com","ansrsource.com"],"blockedDomains": ["test.com","quikr.com", "olx.in"]}`)
	w.Header().Set("Content-Type", "application/json")
	response, err := GetDomainList()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

// GetDomainList will return all the domains
func GetDomainList() ([]byte, error) {
	type domains struct {
		AllowedDomains []string
		BlockedDomains []string
	}
	var AllowedDomainList = []string{"infosys.com", "skidata.com"}
	var BlockedDomainList = []string{"test.com", "gmail.com"}
	var myList = domains{AllowedDomains: AllowedDomainList, BlockedDomains: BlockedDomainList}
	return (json.MarshalIndent(myList, "", " "))
}
