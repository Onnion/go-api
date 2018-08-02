package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func newCard(name string) string {
	return name
}

func card(w http.ResponseWriter, r *http.Request) {
	cards := newDeck()

	hand, ramainingCards := cards.deal(5)

	hand.print()
	fmt.Println("///////////////////////")
	ramainingCards.print()
}

func home(w http.ResponseWriter, r *http.Request) {
	url := "https://jsonplaceholder.typicode.com/posts"
	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, "err %s", err)
	} else {
		body, _ := ioutil.ReadAll(res.Body)
		text := string(body)
		fmt.Fprintf(w, "%s", text)
	}

}

func main() {
	route := mux.NewRouter()
	route.HandleFunc("/", card).Methods("GET")

	// the regex validation inside {id} expecifc what route will access
	route.HandleFunc("/movies/{id:[0-9]+}", home).Methods("GET")

	// defineRoutes(route)

	srv := &http.Server{
		Handler: route,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
