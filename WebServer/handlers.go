package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//HandleRoot index Controller
func HandleRoot(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writter, "<b>Hola Mundo</b>")
}

//HandleProfile - handler for /profile
func HandleProfile(writter http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writter, "<b>This is the user profile</b><a href='http://google.com'>this is a link</a>")
}

//GenericPostRequest handles JSON transformation
func GenericPostRequest(writter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var metadata MetaData

	err := decoder.Decode(&metadata)

	if err != nil {
		fmt.Fprintf(writter, "error %v", err)
		return
	}

	fmt.Fprintf(writter, "Payload %v\n", metadata)
}

//UserPostRequest handles JSON transformation
func UserPostRequest(writter http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var user User

	err := decoder.Decode(&user)

	if err != nil {
		fmt.Fprintf(writter, "error %v", err)
		return
	}
	response, error := user.ToJSON()
	if error != nil {
		writter.WriteHeader(http.StatusInternalServerError)
		return
	}
	writter.Header().Set("Content-Type", "application/json")
	writter.Write(response)
}
