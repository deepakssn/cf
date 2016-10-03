package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// GetDomain will send the list of allowed domains and blocked domains
func GetDomain(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Server", "Blip Gopher!")
	w.WriteHeader(200)
	response, err := GetDomainList()
	if err != nil {
		http.Error(w, "Unable to process data at this time"+err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(response))
}

// GetDomainList will return all the domains
func GetDomainList() ([]byte, error) {
	type domains struct {
		AllowedDomains []string
		BlockedDomains []string
	}

	var dbDomainList domains
	db, err := sql.Open("mysql", "sql:deepu@/devtest")
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	db.Ping()
	var (
		domain string
		allow  int
	)
	rows, err := db.Query("SELECT DOMAIN, ALLOW FROM DOMAIN")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&domain, &allow)
		if err != nil {
			log.Println(err)
		}
		if allow == 1 {
			dbDomainList.AllowedDomains = append(dbDomainList.AllowedDomains, domain)
		} else {
			dbDomainList.BlockedDomains = append(dbDomainList.BlockedDomains, domain)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
	return (json.MarshalIndent(dbDomainList, "", " "))
}
