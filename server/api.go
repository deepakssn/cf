package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/domains", GetDomain)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Index is a standard index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// GetDomain will send the list of valid domains
func GetDomain(w http.ResponseWriter, r *http.Request) {
	// son := []byte(`{ "allowedDomains": ["infosys.com","skidata.com","ansrsource.com"],"blockedDomains": ["test.com","quikr.com", "olx.in"]}`)

	fmt.Fprintln(w, "Domains")
}
