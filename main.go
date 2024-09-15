package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server","DevBox",)
	w.Header().Add("Go","23.1")

	w.Write([]byte("Hello, World!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w,"Display a specific snippet with ID %d", id)
	
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Specific Create"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Specific Create Post"))
}

func main () {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	

	mux.HandleFunc("GET /snippet/view/{id}", snippetView)


	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)


	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}