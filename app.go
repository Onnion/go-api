package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
    "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "asdasd")
}

func main() {

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})



	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/movies", home).Methods("GET")
	r.HandleFunc("/movies", home).Methods("POST")
	r.HandleFunc("/movies", home).Methods("PUT")
	r.HandleFunc("/movies", home).Methods("DELETE")

	// the regex validation inside {id} expecifc what route will access
	r.HandleFunc("/movies/{id:[0-9]+}", home).Methods("GET")
	r.HandleFunc("/movies/{id:[a-z]+}", home).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
