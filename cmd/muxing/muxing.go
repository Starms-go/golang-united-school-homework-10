package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {

	// type head map[string]string

	router := mux.NewRouter()
	router.HandleFunc("/name/{param}", getParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBad).Methods(http.MethodGet)
	router.HandleFunc("/data", getData).Methods(http.MethodPost)
	router.HandleFunc("/headers", getHeader).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func getParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["param"]
	response := fmt.Sprintf("Hello, %s!", param)
	fmt.Fprint(w, response)
}

func getBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(500)
}

func getData(w http.ResponseWriter, r *http.Request) {
	// requestBody := []byte{}
	// r.Body.Read(requestBody)

	defer r.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(r.Body)

	// Error checking of the ioutil.ReadAll() request
	// if err != nil {
	// 	log.Fatal(err)
	// }

	bodyString := string(bodyBytes)

	response := fmt.Sprintf("I got message:\n%s", bodyString)
	fmt.Fprint(w, response)
}

func getHeader(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	aInt, _ := strconv.Atoi(a)

	b := r.Header.Get("b")
	bInt, _ := strconv.Atoi(b)

	headResult := strconv.Itoa(aInt + bInt)
	w.Header().Set("a+b", headResult)

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
