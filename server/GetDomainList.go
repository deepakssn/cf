package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetDomain will send the list of allowed domains and blocked domains
func GetDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// For Testing Purpse ONLY
	if origin := r.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Origin", origin)
	}
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

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
