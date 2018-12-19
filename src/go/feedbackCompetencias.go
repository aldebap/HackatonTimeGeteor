////////////////////////////////////////////////////////////////////////////////
//	feedbackCompetencias.go  -  Dec/15/2018  -  aldebaran perseke
//
//	Web server for feedbackCompetencias app
////////////////////////////////////////////////////////////////////////////////

package main

import (
	//	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gorilla/mux"
)

////////////////////////////////////////////////////////////////////////////////
//	home of the competency selection
////////////////////////////////////////////////////////////////////////////////

func competencySelection(_httpResponse http.ResponseWriter, _httpRequest *http.Request) {

	if "/" == _httpRequest.URL.Path {
		_httpRequest.URL.Path = "/selecao.html"
	}
	selecaoPath := path.Dir("./html/selecao.html")
	_httpResponse.Header().Set("Content-type", "text/html")
	http.ServeFile(_httpResponse, _httpRequest, selecaoPath)
}

////////////////////////////////////////////////////////////////////////////////
//	RESTFul service to confirm the competency selection
////////////////////////////////////////////////////////////////////////////////

func confirmSelectionRequest(_httpResponse http.ResponseWriter, _httpRequest *http.Request) {

	//	_ = json.NewDecoder(_httpRequest.Body).Decode(&request)

	_httpResponse.WriteHeader(http.StatusCreated)
}

////////////////////////////////////////////////////////////////////////////////
//	Start the RESTFul web server
////////////////////////////////////////////////////////////////////////////////

func main() {

	var verbose bool

	//	parse command line arguments
	flag.BoolVar(&verbose, "verbose", false, "print a detailed trace execution")

	flag.Parse()

	if 0 < len(flag.Args()) {
		fmt.Fprintf(os.Stderr, "invalid argument: %s", flag.Args())
		panic(-1)
	}

	//	splash screen
	fmt.Printf(">>>>> Feedback Competency web server\n\n")

	//	start the Web Server
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("./html"))

	router.PathPrefix("/").Handler(competencySelection).Methods("GET")
	//	router.HandleFunc("/img", competencySelection)
	//	router.HandleFunc("/js", competencySelection)
	router.HandleFunc("/confirm-selection", confirmSelectionRequest).Methods("POST")

	log.Panic(http.ListenAndServe(":8080", router))
}
