package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// jsonresponse is the standard response structure used to send JSON response
type jsonSuccess struct {
	SuccessCat string `json:"successCategory"`
	SuccessMsg string `json:"successMessage"`
}

type jsonError struct {
	ErrCat  string `json:"errorCategory"`
	ErrCode string `json:"errorCode"`
	ErrMsg  string `json:"errorMessage"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/domains", GetDomain)
	router.HandleFunc("/email", ValidateUser)
	router.HandleFunc("/otp", ValidateOTP)
	log.Fatal(http.ListenAndServe(":8081", router))
}

// Index is a standard index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
