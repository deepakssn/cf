package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// jsonresponse is the standard response structure used to send JSON response
type jsonSuccess struct {
	Result          string `json:"result"`
	SuccessCategory string `json:"successCategory"`
	SuccessMessage  string `json:"successMessage"`
}

type jsonError struct {
	ErrorCategory string `json:"errorCategory"`
	ErrorCode     string `json:"errorCode"`
	ErrorMessage  string `json:"errorMessage"`
	Result        string `json:"result"`
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

// GenerateFailureResponse generates the JSON response
func GenerateFailureResponse(ErrCat string, ErrCode string, ErrMsg string) []byte {
	response, err := (json.MarshalIndent(jsonError{Result: "fail", ErrorCategory: ErrCat, ErrorCode: ErrCode, ErrorMessage: ErrMsg}, "", " "))
	if err != nil {
		panic(err)
	}
	return response
}

// GenerateSuccessResponse generates the JSON response
func GenerateSuccessResponse(SuccessCat string, SuccessMsg string) []byte {
	response, err := (json.MarshalIndent(jsonSuccess{Result: "success", SuccessCategory: SuccessCat, SuccessMessage: SuccessMsg}, "", " "))
	if err != nil {
		panic(err)
	}
	return response
}
